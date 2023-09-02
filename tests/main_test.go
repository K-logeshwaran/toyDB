package main

import (
	"testing"

	"github.com/K-logeshwaran/goDb/Driver"
)

func TestWhere(t *testing.T) {

	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Initialize a mock database with the temporary directory
	db := Driver.NewDB(tempDir, "test.log", Driver.NewCollection(tempDir))

	// Create a collection
	collectionName := "test_collection"
	err := db.CreateCollection(collectionName)
	if err != nil {
		t.Fatalf("Error creating collection: %v", err)
	}

	// Populate records in the collection
	data := []byte(`{"name": "John"}`)
	for i := 0; i < 3; i++ {
		db.PopulateRecords(collectionName, data)
	}

	// Test cases for different scenarios
	tests := []struct {
		name   string
		field  string
		value  interface{}
		expect int // Expected number of results
	}{
		{"TestWhereFound", "name", "John", 3},
		{"TestWhereNotFound", "name", "Jane", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			results, err := db.Where(collectionName, tc.field, tc.value)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(results) != tc.expect {
				t.Errorf("Expected %d results, got %d", tc.expect, len(results))
			}
		})
	}
}
