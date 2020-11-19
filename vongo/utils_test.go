package vongo

import (
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	type Test struct {
		given interface{}
		want  string
	}
	tests := []Test{
		{given: &VongoDoc{}, want: "VongoDoc"},
		{given: &[]VongoDoc{}, want: "VongoDoc"},
	}

	for index, tc := range tests {
		got := interfaceName(tc.given)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("\n%d \nexpected: %v \ngot %v", index, tc.want, got)
		}
	}
}
