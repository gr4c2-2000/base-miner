package error_handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/rotisserie/eris"
)

func CriticalErrorLog(err error) {
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
}
func LogError(err error) {
	if err != nil {
		formattedJSON := eris.ToJSON(err, true)
		json, errjs := json.Marshal(formattedJSON) // marshal to JSON and print
		if errjs != nil {
			log.Println(err)
		}
		pjson, errjs := PrettyString(json)
		if errjs != nil {
			log.Println(err)
		}
		log.Println(pjson)
	}
}

func PrettyString(by []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, by, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
