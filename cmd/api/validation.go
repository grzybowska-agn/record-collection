package main

import (
	"fmt"
	"strings"
	"time"
)

func validateRecord(record Record) error {
	if strings.TrimSpace(record.Artist) == "" {
		return fmt.Errorf("Artist is required")
	}

	if strings.TrimSpace(record.Title) == "" {
		return fmt.Errorf("Title is required")
	}

	switch record.SizeInches {
	case Size7Inch, Size10INch, Size12Inch:
	default:
		return fmt.Errorf("Size must be 7, 10, or 12 inches")
	}

	switch record.ReleaseType {
	case ReleaseTypeEP, ReleaseTypeLP, ReleaseTypeSingle:
	default:
		return fmt.Errorf("Release type must be LP, EP, or Single")
	}

	if record.DiscCount < 1 {
		return fmt.Errorf("Disc count must be at least 1")
	}

	maxYear := time.Now().Year() + 1

	if record.OriginalYear < 1900 || record.OriginalYear > maxYear {
		return fmt.Errorf("original year must be between 1900 and %d", maxYear)
	}

	if record.PressingYear != nil {
		year := *record.PressingYear

		if year < record.OriginalYear || year > maxYear {
			return fmt.Errorf("pressing year must be between %d and %d",
				record.OriginalYear, maxYear)
		}
	}

	return nil
}
