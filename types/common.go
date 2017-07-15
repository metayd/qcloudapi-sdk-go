package types

import (
	"fmt"
	"net/url"
	"reflect"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message, omitempty"`
}

type CommonArgs struct {
	Action          string
	Region          string
	Timestamp       uint
	Nonce           uint
	SecretId        string
	SignatureMethod string
}

type RequestSignature string

func EncodeArgs(key string, i interface{}, v *url.Values) error {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for index := 0; index < val.Len(); index++ {
			elem := val.Index(index)
			switch elem.Kind() {
			case reflect.Struct:
				{
					elemType := elem.Type()
					for ptr := 0; ptr < elem.NumField(); ptr++ {
						v.Set(
							fmt.Sprintf("%s.%d.%s", key, index, elemType.Field(ptr).Name),
							fmt.Sprint(elem.Field(ptr)),
						)
					}
				}
			default:
				v.Set(
					fmt.Sprintf("%s.%d", key, index),
					fmt.Sprint(elem),
				)
			}
		}
		return nil
	} else if val.Kind() == reflect.Struct {
		valType := val.Type()
		for ptr := 0; ptr < val.NumField(); ptr++ {
			v.Set(
				fmt.Sprintf("%s.%s", key, valType.Field(ptr).Name),
				fmt.Sprint(val.Field(ptr)),
			)
		}
		return nil
	}
	return nil
}

func EncodeFlattenStruct(key string, i interface{}, v *url.Values) error {
	return nil
}
