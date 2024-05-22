package mysql

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/rotisserie/eris"
)

func RowsScan(v interface{}, r *sql.Rows, strict bool) (outerr error) {
	vType := reflect.TypeOf(v)
	if k := vType.Kind(); k != reflect.Ptr {
		return fmt.Errorf("%q must be a pointer", k.String())
	}
	sliceType := vType.Elem()
	if reflect.Slice != sliceType.Kind() {
		return fmt.Errorf("%q must be a slice", sliceType.String())
	}
	sliceVal := reflect.Indirect(reflect.ValueOf(v))
	itemType := sliceType.Elem()

	cols, err := r.Columns()
	if err != nil {
		return err
	}

	isPrimitive := itemType.Kind() != reflect.Struct

	for r.Next() {
		sliceItem := reflect.New(itemType).Elem()

		var pointers []interface{}
		if isPrimitive {
			if len(cols) > 1 {
				return eris.New("ToNamyRows")
			}
			pointers = []interface{}{sliceItem.Addr().Interface()}
		} else {
			pointers = structPointers(sliceItem, cols, strict)
		}

		if len(pointers) == 0 {
			return nil
		}
		err := r.Scan(pointers...)
		if err != nil {
			return err
		}
		rewritePointersForNullableMysqlFields(pointers, sliceItem, cols, strict)
		sliceVal.Set(reflect.Append(sliceVal, sliceItem))
	}
	return r.Err()
}
func rewritePointersForNullableMysqlFields(pointers []interface{}, stct reflect.Value, cols []string, strict bool) {
	for k, v := range pointers {
		if value, ok := v.(*interface{}); ok {
			is, fv := fieldHasNullableTag(stct, cols[k])
			if is && *value != nil {
				val := *value
				Realfield := fieldByName(stct, cols[k], strict)
				setPointer(val, Realfield, fv)
				pointers[k] = Realfield.Addr().Interface()
			}
		}
	}
}
func setPointer(val interface{}, ref reflect.Value, nullableType string) {
	switch t := val.(type) {
	case []uint8:
		ref.SetString(string(t))
	case int64:
		ref.SetInt(t)
	case float64:
		ref.SetFloat(t)
	case bool:
		ref.SetBool(t)
	}
}
func structPointers(stct reflect.Value, cols []string, strict bool) []interface{} {
	pointers := make([]interface{}, 0, len(cols))
	for _, colName := range cols {
		_, fieldVal := nullableFieldByName(stct, colName, strict)
		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			var nothing interface{}
			pointers = append(pointers, &nothing)
			continue
		}
		pointers = append(pointers, fieldVal.Addr().Interface())
	}
	return pointers
}
func fieldByName(v reflect.Value, name string, strict bool) reflect.Value {
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag, ok := typ.Field(i).Tag.Lookup("db")
		if ok && tag == name {
			return v.Field(i)
		}
	}
	if strict {
		return reflect.ValueOf(nil)
	}
	return v.FieldByName(name)
}

func fieldHasNullableTag(v reflect.Value, name string) (bool, string) {
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag, ok := typ.Field(i).Tag.Lookup("db")
		Nullabletag, ok2 := typ.Field(i).Tag.Lookup("sql")
		if ok && tag == name && ok2 {
			return true, Nullabletag
		}
	}
	return false, ""
}
func nullableFieldByName(v reflect.Value, name string, strict bool) (bool, reflect.Value) {
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag, ok := typ.Field(i).Tag.Lookup("db")
		sqlTag, ok2 := typ.Field(i).Tag.Lookup("sql")
		if ok && tag == name && ok2 {
			return true, sqlNullable(v.Field(i), sqlTag)
		}
		if ok && tag == name {
			return false, v.Field(i)
		}

	}
	if strict {
		return false, reflect.ValueOf(nil)
	}
	return false, v.FieldByName(name)
}

func sqlNullable(field reflect.Value, tag string) reflect.Value {
	switch tag {
	case "NullBool":
		return reflect.ValueOf(sql.NullBool{})
	case "NullString":
		return reflect.ValueOf(sql.NullString{})
	case "NullFloat64":
		return reflect.ValueOf(sql.NullFloat64{})
	case "NullInt32":
		return reflect.ValueOf(sql.NullInt32{})
	case "NullInt64":
		return reflect.ValueOf(sql.NullInt64{})
	case "NullTime":
		return reflect.ValueOf(sql.NullTime{})
	default:
		return field
	}

}
