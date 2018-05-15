package data_test

import (
	"testing"

	data "github.com/fourcube/goiban-data"
)

func TestStoreInMemory(t *testing.T) {
	store := data.NewInMemoryStore()
	data := data.BankInfo{Bankcode: "12345", Country: "DE"}

	ok, err := store.Store(data)

	if !ok || err != nil {
		t.Errorf("Test failed: %v %v", ok, err)
	}

	record, _ := store.Find("DE", "12345")

	if *record != data {
		t.Errorf("Failed to find data in store")
	}
}

func TestClearInMemory(t *testing.T) {
	store := data.NewInMemoryStore()
	data := data.BankInfo{Bankcode: "12345", Country: "DE", Source: "Foo"}

	ok, err := store.Store(data)

	if !ok || err != nil {
		t.Errorf("Test failed: %v %v", ok, err)
	}

	store.Clear("Foo")

	record, _ := store.Find("DE", "12345")

	if record != nil {
		t.Errorf("Data not deleted!")
	}
}
