package vongo

import (
	"context"
	"time"
)

// BeforeSaveHook method
type BeforeSaveHook interface {
	BeforeSave() error
}

// AfterSaveHook method
type AfterSaveHook interface {
	AfterSave() error
}

// BeforeDeleteHook method
type BeforeDeleteHook interface {
	BeforeDelete() error
}

// AfterDeleteHook method
type AfterDeleteHook interface {
	AfterDelete() error
}

// AfterFindHook method
type AfterFindHook interface {
	AfterFind() error
}

// ValidateHook method
type ValidateHook interface {
	Validate() []error
}

// ValidationError method
type ValidationError struct {
	Errors []error
}

// NewTrackerHook method
type NewTrackerHook interface {
	IsNew() bool
}

// TimeCreatedAtTrackerHook method
type TimeCreatedAtTrackerHook interface {
	GetCreatedAt() time.Time
	SetCreatedAt(t time.Time)
}

// TimeUpdatedAtTrackerHook method
type TimeUpdatedAtTrackerHook interface {
	GetUpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
}

// ConnectionContextWrapperHook method
type ConnectionContextWrapperHook interface {
	WrapContext(ctx context.Context) context.Context
}

func callToBeforeSaveHook(d Document, isNew bool) error {
	if hook, ok := d.(BeforeSaveHook); ok {
		if err := hook.BeforeSave(); err != nil {
			return err
		}
	}

	// Add created/modified time. Also set on the model itself if it has those fields.
	now := time.Now()
	if isNew {
		d.SetCreatedAt(now)
	}
	d.SetUpdatedAt(now)

	return nil
}

func callToAfterSaveHook(d Document) error {
	if hook, ok := d.(AfterSaveHook); ok {
		if err := hook.AfterSave(); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeRemoveHook(d Document) error {
	if hook, ok := d.(BeforeDeleteHook); ok {
		err := hook.BeforeDelete()
		if err != nil {
			return err
		}
	}

	return nil
}

func callToAfterRemoveHook(d Document) error {
	if hook, ok := d.(AfterDeleteHook); ok {
		if err := hook.AfterDelete(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterFindHook(d Document) error {
	if hook, ok := d.(AfterFindHook); ok {
		if err := hook.AfterFind(); err != nil {
			return err
		}
	}

	return nil
}
