package extractor

import (
	"encoding/json"
	"io"
	"os"

	"imagine_everything_exercise/model"
)

// ExtractAuditJSON extracts JSON that has the audit structure.
func ExtractAuditJSON() (model.AuditRecords, error) {
	var records model.AuditRecords

	jsonFile, err := os.Open("extractor/test_files/audit.json")
	if err != nil {
		return model.AuditRecords{}, err
	}

	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		return model.AuditRecords{}, err
	}

	json.Unmarshal(byteJson, &records)

	return records, nil
}
