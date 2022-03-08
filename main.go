package main

import (
	"github.com/pkg/errors"
	"log"
	"reflect"
)

type inter struct {
}

func PrintStruct(in interface{}, val map[string]interface{}) error {
	if val == nil {
		return errors.New("Didn't found struct")
	}

	rvp := reflect.ValueOf(val)

	rvp = reflect.Indirect(rvp)

	for get := range val {
		rvt := reflect.TypeOf(get)
		log.Printf("typeIs %s", rvt)
	}

	//for i := 0; i < value.NumField(); i++ {
	//	typeField := value.Type().Field(i)
	//
	//	if typeField.Type.Kind() == reflect.Struct {
	//		log.Printf("Internal structer %v found", typeField.Name)
	//		//PrintStruct(value.Field(i).Interface(), val map[string]interface{})
	//		continue
	//	}
	//
	//	log.Printf("\tname=%s, type=%s, value=%v, tag=%s\n",
	//		typeField.Name,
	//		typeField.Type,
	//		value.Field(i),
	//		typeField.Tag,
	//	)
	//}
	return nil
}

func main() {
	var in inter

	mapToTest := map[string]interface{}{
		"id":    1,
		"name":  "nameless",
		"value": 123,
	}

	PrintStruct(in, mapToTest)
}
