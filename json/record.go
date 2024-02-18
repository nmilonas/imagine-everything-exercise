package json

import (
	"encoding/json"

	"imagine_everything_exercise/model"
)

func ExtractJSON(recordsJSON string) model.Records {
	var records model.Records

	json.Unmarshal([]byte(recordsJSON), &records)

	return records
}
