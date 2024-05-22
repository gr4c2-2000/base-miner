package logger

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/rotisserie/eris"
	log "github.com/sirupsen/logrus"
)

func LogError(err error) {
	if err != nil {
		formattedJSON := eris.ToJSON(err, true)
		log.WithFields(log.Fields{"errorStack": formattedJSON}).Error(err)
	}
}
func CriticalErrorLog(err error) {
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
}

func PrettyString(by []byte) (error, string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, by, "", "    "); err != nil {
		return err, ""
	}
	return nil, prettyJSON.String()
}
