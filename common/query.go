package common

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func EncodeStruct(i interface{}, v *url.Values) error {
	val := reflect.ValueOf(i)
	return encodeStructWithPrefix("", val, v)
}

func encodeStructWithPrefix(prefix string, val reflect.Value, v *url.Values) error {
	switch val.Kind() {
	case reflect.Struct:
		{
			typ := val.Type()
			for index := 0; index < val.NumField(); index++ {
				encodeStructWithPrefix(
					strings.Join(
						[]string{prefix, parseTag(typ.Field(index).Tag.Get("url"))},
						".",
					),
					val.Field(index),
					v,
				)
			}
		}
	case reflect.Array, reflect.Slice:
		{
			for index := 0; index < val.Len(); index++ {
				encodeStructWithPrefix(
					strings.Join(
						[]string{prefix, fmt.Sprint(index)},
						".",
					),
					val.Index(index),
					v,
				)
			}
		}
	case reflect.Ptr:
		encodeStructWithPrefix(prefix, val.Elem(), v)
	case reflect.String: {
		if !(val.Len() == 0) {
			v.Set(strings.TrimLeft(prefix, "."), val.String())
		}
	}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:{
		if !(val.Int() == 0) {
			v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
		}
	}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:{
		if !(val.Uint() == 0) {
			v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
		}
	}
	case reflect.Float32, reflect.Float64:{
		if !(val.Float() == 0) {
			v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
		}
	}
	case reflect.Bool:{
		if !val.Bool() {
			v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
		}
	}
	default:
	}
	return nil
}

func parseTag(tag string) string {
	return strings.Split(tag, ",")[0]
}
