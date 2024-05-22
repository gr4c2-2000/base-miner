package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gr4c2-2000/base-miner/pkg/common"
	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/rotisserie/eris"
)

type MySqlGateway struct {
	config           *config.DataSource
	mysqlConnections map[string]*sql.DB
}

func InitMySql(Config *config.DataSource) *MySqlGateway {
	msq := MySqlGateway{}
	msq.config = Config
	msq.mysqlConnections = make(map[string]*sql.DB)
	return &msq
}
func (m *MySqlGateway) Query(ConnectionName string, query string, reciver interface{}, args ...interface{}) error {

	stmt, err := m.prepare(ConnectionName, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return eris.Wrap(err, "Error in Query")
	}
	err = RowsScan(reciver, rows, false)
	if err != nil {
		return eris.Wrap(err, "Error in Scan")
	}
	return nil
}

func (m *MySqlGateway) QueryToMapWithArgs(ConnectionName string, query string, args ...interface{}) ([]map[string]interface{}, error) {

	stmt, err := m.prepare(ConnectionName, query)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	defer stmt.Close()

	sqlResult, err := stmt.Query(args...)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	Result := make([]map[string]interface{}, 0)
	colTypes, err := sqlResult.ColumnTypes()
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}

	for sqlResult.Next() {
		scanArgs := PrepareScanArgs(colTypes)
		err := sqlResult.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		row := PrepareRow(colTypes, scanArgs)
		Result = append(Result, row)
	}
	if err != nil {
		return nil, eris.Wrapf(err, "")

	}
	return Result, nil
}

func (m *MySqlGateway) Insert(ConnectionName string, query string, args ...interface{}) (int64, error) {
	db, err := m.connection(ConnectionName)

	if err != nil {
		return 0, eris.Wrapf(err, "")
	}
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, eris.Wrapf(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, eris.Wrapf(err, "")
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, eris.Wrapf(err, "")
	}

	return lastID, nil
}
func (m *MySqlGateway) prepare(ConnectionName string, query string) (*sql.Stmt, error) {
	db, err := m.connection(ConnectionName)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return stmt, nil
}

func (m *MySqlGateway) connection(ConnectionName string) (*sql.DB, error) {
	val, ok := m.mysqlConnections[ConnectionName]
	if ok && val.Ping() == nil {
		return m.mysqlConnections[ConnectionName], nil
	}

	err := m.setConnection(ConnectionName)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return m.mysqlConnections[ConnectionName], nil
}

func (m *MySqlGateway) getConnection(ConnectionName string) (*sql.DB, error) {
	dbConfig, err := m.config.FindConnectionByName(ConnectionName)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	db, err := sql.Open(common.MYSQL, dbConfig.GetConnectionString())
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(50)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return db, nil
}

func (m *MySqlGateway) setConnection(ConnectionName string) error {
	con, err := m.getConnection(ConnectionName)
	if err != nil {
		return eris.Wrapf(err, "")
	}
	m.mysqlConnections[ConnectionName] = con
	return nil
}
