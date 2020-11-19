package vongo

import (
	"context"
	"testing"
)

type HookedDoc struct {
	DocumentModel `bson:",inline" coll:"hooked-registry"`

	beforeSavingCalled bool
	afterSavingCalled  bool

	beforeDeletingCalled bool
	afterDeletingCalled  bool
}

func (s *HookedDoc) BeforeSave() error {
	s.beforeSavingCalled = true
	return nil
}

func (s *HookedDoc) AfterSave() error {
	s.afterSavingCalled = true
	return nil
}

func TestDocumentHooks(t *testing.T) {
	setupDefConnection()
	ut := HookedDoc{}

	_, err := getCollectionName(ut)
	if err != nil {
		t.Errorf("HookedDoc getCollectionName - Err: %v", err)
	}

	if ut.beforeSavingCalled || ut.afterSavingCalled {
		t.Errorf("HookedDoc beforeSavingCalled and afterSavingCalled should be false")
	}

	err = Save(context.TODO(), &ut)

	if !ut.beforeSavingCalled || !ut.afterSavingCalled {
		t.Errorf("HookedDoc beforeSavingCalled and afterSavingCalled should be true")
	}
}
