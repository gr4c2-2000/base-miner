package yamlparse

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

func ParseYamlAndCreateStuct[T any](Path string) (*T, error) {
	str := new(T)
	err := ParseYaml(Path, str)
	return str, err
}
func ParseYaml(dfPath string, destination interface{}) error {
	source, err := ioutil.ReadFile(dfPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(source, destination)
	if err != nil {
		return err
	}
	return nil
}
