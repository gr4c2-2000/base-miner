package kernal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/gr4c2-2000/base-miner/pkg/esconnection"
	"github.com/gr4c2-2000/base-miner/pkg/mysql"
)

var kernal Kernal

func init() {
	kernal = Kernal{}
	kernal.DataSourceConfig = config.InitDataSource()
	kernal.MysqlGatway = mysql.InitMySql(kernal.DataSourceConfig)
	kernal.ElasticGateway = esconnection.InitES(kernal.DataSourceConfig)
}

type Kernal struct {
	DataSourceConfig *config.DataSource
	MysqlGatway      *mysql.MySqlGateway
	ElasticGateway   esconnection.EsMap
}

func Get() *Kernal {
	return &kernal
}
