package vongo

import (
	"fmt"
	"log"
	"reflect"
)

// GetCollectionName method
func getCollectionName(i interface{}) (string, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Slice {
		inner := v.Type().Elem()
		switch inner.Kind() {
		case reflect.Struct:
			t = inner
			v = reflect.New(inner).Elem()
		}
	}

	if t.Kind() == reflect.Ptr {
		t = reflect.Indirect(reflect.ValueOf(i)).Type()
		v = reflect.ValueOf(i).Elem()

		if v.Kind() == reflect.Slice {
			return getCollectionName(reflect.New(t.Elem()).Interface())
		}
	}

	n := t.Name()
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("Only type struct can be used as document model (passed type %s is not struct) t: %v", n, t.Kind())
	}
	var idx = -1
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)
		if ft.Type.ConvertibleTo(reflect.TypeOf(DocumentModel{})) {
			idx = i
			break
		}
	}

	if idx == -1 {
		return "", fmt.Errorf("A document model must embed a DocumentModel type field (passed type %s does not have)", n)
	}

	coll := initializeTags(t, v)
	if coll == "" {
		return "", fmt.Errorf("The document model does not have a collection name (passed type %s)", n)
	}

	return coll, nil
}

func initializeTags(t reflect.Type, v reflect.Value) string {
	var coll = ""

	for i := 0; i < v.NumField(); i++ {
		// f := v.Field(i)
		ft := t.Field(i)
		// n := "_" + ft.Name
		switch ft.Type.Kind() {
		case reflect.Struct:
			if ft.Type.ConvertibleTo(reflect.TypeOf(DocumentModel{})) {
				coll = extractColl(ft)
				break
			}
			fallthrough
		default:
			logBadColl(ft)
		}
	}

	return coll
}

func logBadColl(sf reflect.StructField) {
	if extractColl(sf) != "" {
		log.Printf("Tag collection used outside DocumentModel is ignored (field: %s)", sf.Name)
	}
}

func extractColl(sf reflect.StructField) string {
	coll := sf.Tag.Get("coll")
	if coll == "" {
		coll = sf.Tag.Get("collection")
	}

	return coll
}
