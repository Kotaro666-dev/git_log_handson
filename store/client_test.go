package store

import (
	"encoding/hex"
	"testing"
)

func TestClient_GetObject(t *testing.T) {
	client, err := NewClient("../testrepo")
	if err != nil {
		t.Fatal(err)
	}
	hashString := "06fdbe2834fba5a7231fafa2524d67eddbd38e65"
	hash,err := hex.DecodeString(hashString)
	if err != nil{
		t.Fatal(err)
	}
	object, err := client.GetObject(hash)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(string(object.Data))
}
