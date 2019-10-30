package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/cadence/encoded"
)

var (
	pbMarshaler   = jsonpb.Marshaler{}
	pbUnmarshaler = jsonpb.Unmarshaler{}
)

type CustomDataConverter struct{}

func (dc *CustomDataConverter) ToData(objs ...interface{}) ([]byte, error) {
	result := &bytes.Buffer{}
	for i, obj := range objs {
		// Try jsonpb...
		if pbObj := findProtoMessage(obj); pbObj != nil {
			fmt.Printf("\n\n[CustomDataConverter.ToData()] Encoding a proto.Message: %T\n\n", obj)
			curBytes := &bytes.Buffer{}
			err := pbMarshaler.Marshal(curBytes, pbObj)
			if err != nil {
				if err == io.EOF {
					return nil, fmt.Errorf("missing argument at index %d of type %T", i, obj)
				}
				return nil, fmt.Errorf(
					"unable to encode argument: %d, %v, with jsonpb error: %v", i, reflect.TypeOf(obj), err,
				)
			}
			// fmt.Printf("\n\n[CustomDataConverter.ToData()] jsonpb wrote bytes:\n%s\n\n", curBytes.String())
			_, _ = io.Copy(result, curBytes)
		} else {
			// ... falling back to the default DC
			fmt.Printf("\n\n[CustomDataConverter.ToData()] Encoding a normal type: %T\n\n", obj)
			curBytes, err := encoded.GetDefaultDataConverter().ToData(obj)
			if err != nil {
				return nil, err
			}
			_, _ = io.Copy(result, bytes.NewReader(curBytes))
		}
	}
	return result.Bytes(), nil
}

func (dc *CustomDataConverter) FromData(data []byte, objs ...interface{}) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	for i, obj := range objs {
		var intermediate json.RawMessage
		if err := dec.Decode(&intermediate); err != nil {
			return fmt.Errorf(
				"unable to decode argument: %d, %v, with json error: %v", i, reflect.TypeOf(obj), err,
			)
		}

		// fmt.Printf("\n\n[CustomDataConverter.FromData()] Ready to decode bytes:\n%s\n\n", string(intermediate))

		// Try jsonpb...
		if pbObj := findProtoMessage(obj); pbObj != nil {
			fmt.Printf("\n\n[CustomDataConverter.FromData()] Decoding to a proto.Message: %T\nOriginal obj:\n----------\n%s\n----------\nCasted pbObj:\n----------\n%s\n\n", obj, spew.Sprintf("%#+v", obj), spew.Sprintf("%#+v", pbObj))
			err := pbUnmarshaler.Unmarshal(bytes.NewReader(intermediate), pbObj)
			if err != nil {
				return fmt.Errorf(
					"unable to encode argument: %d, %v, with jsonpb error: %v", i, reflect.TypeOf(obj), err,
				)
			}
		} else {
			// ... falling back to the default DC
			fmt.Printf("\n\n[CustomDataConverter.FromData()] Decoding to a normal type: %T\n\n", obj)
			err := encoded.GetDefaultDataConverter().FromData(intermediate, obj)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

var (
	protoMessageIface = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

func findProtoMessage(obj interface{}) proto.Message {
	// Strip all pointers away
	curObj := obj
	prevObj := obj
	for reflect.TypeOf(curObj).Kind() == reflect.Ptr {
		prevObj = curObj
		curObj = reflect.ValueOf(curObj).Elem().Interface()
	}

	// prevObj should be a pointer to the concrete type
	if reflect.TypeOf(prevObj).Implements(protoMessageIface) {
		return prevObj.(proto.Message)
	}

	// It's possible that curObj was a concrete type from the start,
	// so we need to make a pointer to the original value and test that
	ptrType := reflect.PtrTo(reflect.TypeOf(curObj))
	if ptrType.Implements(protoMessageIface) {
		// TODO: Now what?
		// Ideally we would get the original value and slap a pointer on it, and return that.
		// But how?
	}

	return nil
}
