package main

import (
	"encoding/json"
	"fmt"
)

type RecordPatch map[string]json.RawMessage

func applyRecordPatch(record *Record, patch RecordPatch) error {
	for field, value := range patch {
		switch field {
		case "Artist", "Title", "OriginalYear", "CoverURL",
			"PressingLabel", "SizeInches", "ReleaseType",
			"DiscCount", "PressingCountry":
			if string(value) == "null" {
				return fmt.Errorf("%s cannot be null", field)
			}
		case "PressingYear", "IsReissue":
		default:
			return fmt.Errorf("unknown or non-editable field: %s", field)
		}
	}

	data, err := json.Marshal(patch)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, record)
}
