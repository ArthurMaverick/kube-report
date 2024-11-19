package output

import (
	"encoding/json"
	"errors"
	"os"
)

func (j *JsonOutput) JsonFile() error {
	infos := j.fmtClient.FormatJSONData()

	file, err := os.Create("output.json")
	if err != nil {
		return errors.New("error creating file")
	}

	jsonData, err := json.MarshalIndent(infos, "", "    ")
	if err != nil {
		return errors.New("error marshalling data")
	}
	file.Write(jsonData)
	return nil
}
