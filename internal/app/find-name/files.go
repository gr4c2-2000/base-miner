package findname

import (
	"encoding/csv"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func ParseNamesToMap() map[string]struct{} {

	namesMap := map[string]struct{}{}

	fn := func(val []string) {
		namesMap[val[0]] = struct{}{}
	}

	parse(fn, "../../static/female_name_2024.csv", "../../static/female_name_2024.csv")
	return namesMap
}

func ParseSurNameToMap() map[string]struct{} {

	namesMap := map[string]struct{}{}

	fn := func(val []string) {
		namesMap[val[0]] = struct{}{}
	}

	parse(fn, "../../static/female_surname_2024.csv", "../../static/female_surname_2024.csv")
	return namesMap
}

func parse(fn func([]string), paths ...string) {
	start := time.Now()
	defer func() { log.Infof("Parsing took %v", time.Since(start)) }()

	for _, path := range paths {
		names, err := readCsvFile(path)
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, val := range names {
			fn(val)
		}
	}
}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
