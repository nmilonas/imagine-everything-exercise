package extractor

import (
	"encoding/json"
	"io"
	"os"

	"imagine_everything_exercise/model"
)

func ExtractJSON() (model.Records, error) {
	var records model.Records

	jsonFile, err := os.Open("extractor/test_files/audit.json")
	if err != nil {
		return model.Records{}, err
	}

	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		return model.Records{}, err
	}

	json.Unmarshal(byteJson, &records)

	return records, nil
}
