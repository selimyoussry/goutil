package goutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// SetField can convert a map[string]interface{} to a struct
func SetField(obj interface{}, name string, value interface{}) error {
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

// DumbRestruct uses JSON marshalling to re-cast a map[string]interface{} to a struct
func JsonRestruct(src interface{}, dest interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return err
	}

	return nil
}

// How to use
// type MyStruct struct {
//     Name string
//     Age  int64
// }
//
// func (s *MyStruct) FillStruct(m map[string]interface{}) error {
//     for k, v := range m {
//         err := SetField(s, k, v)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }
//
// func main() {
//     myData := make(map[string]interface{})
//     myData["Name"] = "Tony"
//     myData["Age"] = int64(23)
//
//     result := &MyStruct{}
//     err := result.FillStruct(myData)
//     if err != nil {
//         fmt.Println(err)
//     }
//     fmt.Println(result)
// }
