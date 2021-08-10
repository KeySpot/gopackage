package keyspot

import (
	"testing"
	"os"
)

var accessKey string = "6113019833f1fff67d3ba3b5"

func TestGetRecord(t *testing.T) {
	record, err := GetRecord(accessKey)
	
	if err != nil {
		t.Error(err)
	}

	if record["VAR1"] != "val1" ||
		record["VAR2"] != "val2" ||
		record["VAR3"] != "val3" {
			t.Fatalf("The record doesn't match")
	}
}

func TestSetEnvironment(t *testing.T) {
	err := SetEnvironment(accessKey)

	if err != nil {
		t.Error(err)
	}

	if os.Getenv("VAR1") != "val1" ||
		os.Getenv("VAR2") != "val2" ||
		os.Getenv("VAR3") != "val3" {
			t.Fatalf("The record doesn't match")
	}
}