package pclient

import (
	"os"
	"testing"
)

func TestCollection_all_cases(t *testing.T) {
	c := NewTokenClient(os.Getenv("Authorization"))

	userCollection, err := c.ViewCollection("users", ViewCollectionParams{
		Fields: "",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if userCollection.Name != "users" {
		t.Fatalf("unexpected response: %+v", userCollection)
	}

	collections, err := c.ListCollection(ListCollectionParams{
		Page:      0,
		PerPage:   0,
		Sort:      "",
		Filter:    "",
		Fields:    "",
		SkipTotal: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(collections.Items) == 0 {
		t.Fatalf("unexpected response: %+v", collections)
	}

	if collections.Items[0].Name != "users" {
		t.Fatalf("unexpected response: %+v", collections)
	}
}
