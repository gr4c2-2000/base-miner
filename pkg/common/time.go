package common

import (
	"strconv"
	"time"

	"github.com/rotisserie/eris"
)

func TimeParseToUnix(layout string, timeString string, timezone string) (int64, error) {
	timeLocation, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, eris.Wrapf(err, "")
	}
	result, err := time.ParseInLocation(layout, timeString, timeLocation)
	if err != nil {
		return 0, eris.Wrapf(err, "")
	}
	return result.Unix(), nil
}

func StringParseToTime(layout string, timeString string, timezone string) (*time.Time, error) {
	timeLocation, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	result, err := time.ParseInLocation(layout, timeString, timeLocation)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return &result, nil
}

func UnixStringToTime(UnixString string) (*time.Time, error) {
	i, err := strconv.ParseInt(UnixString, 10, 64)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	result := time.Unix(i, 0)
	return &result, nil

}
