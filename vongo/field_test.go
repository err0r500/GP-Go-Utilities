package vongo

import (
	"fmt"
	"reflect"
	"testing"
)

type VongoDoc struct {
	DocumentModel `bson:",inline" coll:"vongo-registry"` // The VongoDoc will be stored in the vongo-registry collection
	Name          string
}

type NonDocumentModel struct {
	Name string
}

type MissingCollectionNameDoc struct {
	DocumentModel `bson:",inline"` // The VongoDoc will be stored in the vongo-registry collection
	Name          string
}

func TestInitializeTags(t *testing.T) {
	tOf := reflect.TypeOf(VongoDoc{})
	vOf := reflect.ValueOf(VongoDoc{})
	collectionName := initializeTags(tOf, vOf)

	if collectionName != "vongo-registry" {
		t.Errorf("collectionName isn't the one expected: %s", collectionName)
	}
}

func TestGetCollectionName(t *testing.T) {

	type test struct {
		object         interface{}
		collectionName string
		expectedErr    error
	}

	type VongoDocList []VongoDoc

	tests := []test{
		{
			object:         VongoDoc{},
			collectionName: "vongo-registry",
			expectedErr:    nil,
		},
		{
			object:         []VongoDoc{{}, {}},
			collectionName: "vongo-registry",
			expectedErr:    nil,
		},
		{
			object:         &VongoDocList{},
			collectionName: "vongo-registry",
			expectedErr:    nil,
		},
		{
			object:         NonDocumentModel{},
			collectionName: "",
			expectedErr:    fmt.Errorf("A document model must embed a DocumentModel type field (passed type NonDocumentModel does not have)"),
		},
		{
			object:         MissingCollectionNameDoc{},
			collectionName: "",
			expectedErr:    fmt.Errorf("The document model does not have a collection name (passed type MissingCollectionNameDoc)"),
		},
	}

	for index, tc := range tests {
		got, err := getCollectionName(tc.object)
		if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
			t.Fatalf("\n%d - collectionName: %s - err: %v - expectedErr: %v", index, got, err, tc.expectedErr)
		} else if (err != nil && tc.expectedErr == nil) || (err == nil && tc.expectedErr != nil) {
			t.Fatalf("\n%d - collectionName: %s - err: %v - expectedErr: %v", index, got, printOptional(err), printOptional(tc.expectedErr))
		} else if got != tc.collectionName {
			t.Fatalf("\n%d - collectionName: %s - expected: %s", index, got, tc.collectionName)
		}
	}
}
