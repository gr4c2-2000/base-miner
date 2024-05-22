package repo

import (
	. "github.com/gr4c2-2000/base-miner/pkg/data"
)

type ExampleTable struct {
	Id      int     `db:"id"`
	Field1  string  `db:"field_1"`
	Number1 float64 `db:"number_1"`
	Number2 int     `db:"number_2"`
}

func NewBankMidQuery() (*Query[ExampleTable], error) {
	Struct := Query[ExampleTable]{}
	Struct.SetDataSource("one")
	_, err := Struct.SetFile("ExampleTable.sql")
	if err != nil {
		return nil, err
	}
	return &Struct, nil
}
