package connector

import (
	"github.com/gr4c2-2000/base-miner/pkg/data"
	"github.com/gr4c2-2000/base-miner/pkg/mysql"
)

type Mysql struct {
	*mysql.MySqlConnector
}

func (mr *Mysql) QueryFromFile(query data.QueryInterface) error {
	err := mr.MySqlConnector.Query(query.GetDataSource(), query.GetQuery(), query.Recive(), query.GetArgs()...)
	if err != nil {
		return err
	}
	return nil
}
