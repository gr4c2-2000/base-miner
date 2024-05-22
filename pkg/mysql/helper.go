package mysql

import (
	"database/sql"
	"unicode/utf8"

	"github.com/djimenez/iconv-go"
)

const mysqlDateFormat = "2006-01-02"

func PrepareScanArgs(colTypes []*sql.ColumnType) []interface{} {
	scanArgs := make([]interface{}, len(colTypes))
	for i, v := range colTypes {

		switch v.DatabaseTypeName() {
		case "VARCHAR", "TEXT", "TIMESTAMP", "DATETIME", "JSON":
			scanArgs[i] = new(sql.NullString)
		case "BOOL":
			scanArgs[i] = new(sql.NullBool)
		case "TINYINT", "SMALLINT", "INT", "BIGINT":
			scanArgs[i] = new(sql.NullInt64)
		default:
			scanArgs[i] = new(sql.NullString)
		}
	}
	return scanArgs
}

func PrepareRow(colTypes []*sql.ColumnType, scanArgs []interface{}) map[string]interface{} {
	row := map[string]interface{}{}

	for i, v := range colTypes {
		if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
			row[v.Name()] = z.Bool
			continue
		}
		if z, ok := (scanArgs[i]).(*sql.NullString); ok {
			row[v.Name()] = convertToUtf8(z.String)
			continue
		}
		if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
			row[v.Name()] = z.Int64
			continue
		}
		if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
			row[v.Name()] = z.Float64
			continue
		}
		if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
			row[v.Name()] = z.Int32
			continue
		}

		row[v.Name()] = scanArgs[i]
	}

	return row
}
func convertToUtf8(input string) string {

	if !utf8.ValidString(input) {
		converted, err := iconv.ConvertString(input, "iso-8859-2", "utf-8")
		if err != nil {
			return input
		}
		if utf8.ValidString(converted) {
			return converted
		}
	}

	return input
}
