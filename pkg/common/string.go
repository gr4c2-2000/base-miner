package common

import (
	"fmt"
	"strconv"

	"github.com/gr4c2-2000/base-miner/pkg/error_handlers"
)

func InterfaceToString(i interface{}) (string, error) {

	switch v := i.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil
	case float32:
		return fmt.Sprintf("%f", v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "", nil
	case bool:
		if v {
			return "1", nil
		}
		return "0", nil
	default:
		return "", error_handlers.NO_IMPLEMENTATION_ERROR
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
