package main

import (
	"encoding/json"
	"testing"
)

func TestApplyRecordPatchPreservesOmittedFields(t *testing.T) {
	record := Record{
		Artist: "Geese",
		Title:  "3D Country",
	}

	patch := RecordPatch{
		"Title": json.RawMessage(`"Getting Killed"`),
	}

	if err := applyRecordPatch(&record, patch); err != nil {
		t.Fatal(err)
	}

	if record.Title != "Getting Killed" {
		t.Errorf("title was not updated: %q", record.Title)
	}

	if record.Artist != "Geese" {
		t.Errorf("artist changed unexpectedly: %q", record.Artist)
	}
}

func TestApplyRecordPatchClearsPressingYear(t *testing.T) {
	year := 2023
	record := Record{PressingYear: &year}

	patch := RecordPatch{
		"PressingYear": json.RawMessage(`null`),
	}

	if err := applyRecordPatch(&record, patch); err != nil {
		t.Fatal(err)
	}

	if record.PressingYear != nil {
		t.Error("expected pressing year to be cleared")
	}
}

func TestApplyRecordPatchRejectsID(t *testing.T) {
	record := Record{ID: 1}
	patch := RecordPatch{
		"ID": json.RawMessage(`999`),
	}

	if err := applyRecordPatch(&record, patch); err == nil {
		t.Fatal("expected an error when patching ID")
	}

	if record.ID != 1 {
		t.Error("record ID changed")
	}
}
