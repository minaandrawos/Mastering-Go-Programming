package hydragob

import (
	"bytes"
	"encoding/gob"
	"io"
)

func EncodeAndWriteGob(obj interface{}, w io.Writer) error {

	return gob.NewEncoder(w).Encode(obj)
}

func DecodeAndReadGob(b []byte) (interface{}, error) {
	r := bytes.NewReader(b)
	//var obj interface{}
	obj := new(Ship)
	err := gob.NewDecoder(r).Decode(obj)
	return obj, err
}
