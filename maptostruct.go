package maptostruct

import (
	"errors"
	"fmt"
	"reflect"
)

// Do converts a map to a structure
func Do(m map[string]interface{}, s interface{}) error {
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
