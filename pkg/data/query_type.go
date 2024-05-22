package data

import (
	"io/ioutil"
)

type QueryInterface interface {
	GetQuery() string
	GetDataSource() string
	GetArgs() []interface{}
	Recive() interface{}
}
type Query[T any] struct {
	Result     []T
	query      string
	dataSource string
	args       []interface{}
	file       string
}

func (q *Query[T]) InitResult() {
	q.Result = make([]T, 0)
}
func (q *Query[T]) Recive() interface{} {
	return &q.Result
}
func (q *Query[T]) GetQuery() string {
	return q.query
}

func (q *Query[T]) GetDataSource() string {
	return q.dataSource
}

func (q *Query[T]) GetArgs() []interface{} {
	return q.args
}

func (q *Query[T]) SetArgs(args ...interface{}) *Query[T] {
	q.args = args
	return q
}
func (q *Query[T]) SetDataSource(dbName string) *Query[T] {
	q.dataSource = dbName
	return q
}
func (q *Query[T]) SetFile(filename string) (*Query[T], error) {

	str, err := ioutil.ReadFile("../../query-templates/" + filename)
	if err != nil {
		return nil, err
	}
	q.file = filename
	q.query = string(str)
	return q, nil
}

func (q *Query[T]) SetQuery(query string) *Query[T] {
	q.query = query
	return q
}
