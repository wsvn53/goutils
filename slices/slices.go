package slices

import (
	"reflect"
)

func Remove(s interface{}, index int) interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		panic("input not a slice")
	}
	sleft := reflect.MakeSlice(v.Type(), index, index)
	reflect.Copy(sleft, v)

	l := v.Len()
	index += 1
	if index > l {
		panic("index out of bounds")
	}
	for ; index < l; index++ {
		sleft = reflect.Append(sleft, v.Index(index))
	}
	return sleft.Interface()
}
