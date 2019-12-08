package gopopu

import (
	"errors"
	"log"
	"reflect"
)

// Populate populate raw's data to target
func Populate(target interface{}, raw interface{}) error {
	var rvPtr = reflect.ValueOf(target)

	if rvPtr.Kind() != reflect.Ptr || rvPtr.IsNil() {
		return errors.New("param target must Pointer")
	}

	rawMap, err := struct2map(raw)
	if err != nil {
		return err
	}

	v := rvPtr.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name

		if value, ok := rawMap[name]; ok {
			log.Println(name, value)
			if reflect.ValueOf(value).Type() == v.FieldByName(name).Type() {
				v.FieldByName(name).Set(reflect.ValueOf(value))
			}
		}
	}
	return nil
}

// struct2map
// type of obj to map
func struct2map(obj interface{}) (map[string]interface{}, error) {
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("rv must Pointer")
	}

	v := rv.Elem()
	t := v.Type()
	data := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[t.Field(i).Name] = field.Interface()
	}
	return data, nil
}