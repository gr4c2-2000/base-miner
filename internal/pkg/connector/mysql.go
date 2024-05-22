package connector

import (
	"github.com/gr4c2-2000/base-miner/pkg/data"
	"github.com/gr4c2-2000/base-miner/pkg/mysql"
)

type MysqlRepository struct {
	mysql.MySqlGateway
}

func (mr *MysqlRepository) QueryFromFile(query data.QueryInterface) error {
	err := mr.MySqlGateway.Query(query.GetDataSource(), query.GetQuery(), query.Recive(), query.GetArgs()...)
	if err != nil {
		return err
	}
	return nil
}
