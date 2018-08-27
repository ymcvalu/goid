package goid

import (
	"reflect"
)

func getG() interface{}

func getOffset() uintptr {
	giface := getG()
	gtype := reflect.TypeOf(giface)
	field, ok := gtype.FieldByName("goid")
	if !ok {
		panic("goid not found")
	}
	return field.Offset
}

var offset = getOffset()

func Goid() int64
