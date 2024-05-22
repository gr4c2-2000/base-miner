package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gr4c2-2000/base-miner/pkg/common"
	"github.com/gr4c2-2000/base-miner/pkg/error_handlers"
	"github.com/gr4c2-2000/base-miner/pkg/logger"
	yamlparse "github.com/gr4c2-2000/base-miner/pkg/yaml-parse"
)

type DatabaseConfig struct {
	DatabaseName string `yaml:"databaseConnectionName" validate:"required"`
	Host         string `yaml:"host" validate:"required"`
	Port         string `yaml:"port" validate:"required"`
	User         string `yaml:"user" validate:"required"`
	Password     string `yaml:"password" validate:"required"`
	Database     string `yaml:"database" validate:"required"`
	Charset      string `yaml:"charset"`
}

type ElasticSearchConfig struct {
	DatabaseName string `yaml:"databaseConnectionName" validate:"required"`
	Host         string `yaml:"host" validate:"required"`
	IgnoreSSL    bool   `yaml:"ignoreSSL"`
	Version      int    `yaml:"ver" validate:"oneof=5 8"`
	Default      bool   `yaml:"isDefault" `
	SocksProxy   string `yaml:"SocksProxy"`
}

type DataSource struct {
	Database      []DatabaseConfig      `yaml:"database" validate:"required,dive,required"`
	ElasticSearch []ElasticSearchConfig `yaml:"elasticsearch" validate:"required,dive,required"`
}

func InitDataSource() *DataSource {
	source := ParseConfig("../../config/database.yaml")
	return source
}

func ParseConfig(dfPath string) *DataSource {
	validate := common.GetValidator()
	config, err := yamlparse.ParseYamlAndCreateStuct[DataSource](dfPath)
	if err != nil {
		logger.CriticalErrorLog(err)
	}
	err = validate.Struct(config)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		logger.CriticalErrorLog(validationErrors)
	}

	return config
}

func (config *DataSource) FindConnectionByName(name string) (*DatabaseConfig, error) {

	for _, value := range config.Database {
		if value.DatabaseName == name {
			return &value, nil
		}
	}
	return nil, error_handlers.CONNECTION_NOT_EXISTS_ERROR
}

func (config *DataSource) FindElasticConnectionByName(name string) (*ElasticSearchConfig, error) {

	for _, value := range config.ElasticSearch {
		if value.DatabaseName == name {
			return &value, nil
		}
	}
	return nil, error_handlers.CONNECTION_NOT_EXISTS_ERROR
}

func (dbConfig *DatabaseConfig) GetConnectionString() string {
	ConnectionString := dbConfig.User + ":" + dbConfig.Password + "@" + common.TCP + "(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database
	if dbConfig.Charset != "" {
		ConnectionString = ConnectionString + "?charset=" + dbConfig.Charset
	}
	return ConnectionString
}

func (config *DataSource) GetDefaultES() (*ElasticSearchConfig, error) {
	for _, value := range config.ElasticSearch {
		if value.Default == true {
			return &value, nil
		}
	}
	return nil, error_handlers.CONNECTION_NOT_EXISTS_ERROR
}
