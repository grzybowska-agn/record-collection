package main

import "testing"

func TestValidateRecordRejectsInvalidSize(t *testing.T) {
	record := Record{
		Artist:       "Geese",
		Title:        "3D Country",
		OriginalYear: 2023,
		SizeInches:   8,
		ReleaseType:  ReleaseTypeLP,
		DiscCount:    1,
	}

	err := validateRecord(record)
	if err == nil {
		t.Fatal("Expected an error for an invalid size")
	}
}

func TestValidateRecordAcceptsValidRecord(t *testing.T) {
	record := Record{
		Artist:       "Geese",
		Title:        "3D Country",
		OriginalYear: 2023,
		SizeInches:   Size12Inch,
		ReleaseType:  ReleaseTypeLP,
		DiscCount:    1,
	}

	err := validateRecord(record)
	if err != nil {
		t.Fatalf("Expected a valid record, got: %v", err)
	}
}
