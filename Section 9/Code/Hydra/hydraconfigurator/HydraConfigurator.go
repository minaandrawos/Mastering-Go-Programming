package hydraconfigurator

import (
	"errors"
	"reflect"
)

//Configuration file types
const (
	CUSTOM uint8 = iota
	JSON
	XML
)

//error to return if obj is not a pointer to a struct
var errWrongType = errors.New("Type must be a pointer to a struct")

//GetConfiguration reads the supplied file then fills the supplied struct with configuration parameters
func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	//check if this is type pointer
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return errWrongType
	}
	//get and confirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return errWrongType
	}
	switch confType {
	case CUSTOM:
		err = marshalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}
	return err
}
