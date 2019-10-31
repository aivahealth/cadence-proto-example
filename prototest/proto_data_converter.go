package prototest

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"go.uber.org/cadence/encoded"
)

var (
	pbMarshaler   = jsonpb.Marshaler{}
	pbUnmarshaler = jsonpb.Unmarshaler{}
)

type protoDataConverter struct {
	defaultDataConverter encoded.DataConverter
}

func NewProtoDataConverter(defaultDataConverter encoded.DataConverter) encoded.DataConverter {
	return &protoDataConverter{defaultDataConverter: defaultDataConverter}
}

func (dc *protoDataConverter) ToData(objs ...interface{}) ([]byte, error) {
	if len(objs) != 1 {
		return dc.defaultDataConverter.ToData(objs)
	}
	obj := objs[0]
	if !isProtoMessagePtr(obj) {
		fmt.Printf("\n\n[protoDataConverter.ToData()] Not proto argument: %v", reflect.ValueOf(obj).Type())

		return dc.defaultDataConverter.ToData(objs)
	}
	fmt.Printf("\n\n[protoDataConverter.ToData()] Encoding a proto.Message: %T\n\n", obj)
	result := &bytes.Buffer{}
	pbObj := toProtoMessageObject(obj)
	if pbObj == nil {
		return nil, fmt.Errorf("missing proto.Message argument of type %T", obj)
	}
	err := pbMarshaler.Marshal(result, pbObj)
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("missing argument of type %T", obj)
		}
		return nil, fmt.Errorf(
			"unable to encode argument: %v, with jsonpb error: %v", reflect.TypeOf(obj), err,
		)
	}
	// fmt.Printf("\n\n[protoDataConverter.ToData()] jsonpb wrote bytes:\n%s\n\n", curBytes.String())
	return result.Bytes(), nil
}

func (dc *protoDataConverter) FromData(data []byte, objs ...interface{}) error {
	if len(objs) != 1 {
		return dc.defaultDataConverter.FromData(data, objs)
	}
	obj := objs[0]
	if !isProtoMessagePtr(obj) && !isProtoMessagePtrPtr(obj) {
		fmt.Printf("\n\n[protoDataConverter.FromData()] Not a proto message\n", obj)

		return dc.defaultDataConverter.FromData(data, objs)
	}
	fmt.Printf("\n\n[protoDataConverter.FromData()] Decoding a proto.Message: %T\n\n", obj)
	q := reflect.ValueOf(obj)
	if isProtoMessagePtrPtr(obj) {
		q = q.Elem()
	}
	if q.IsNil() {
		v := reflect.New(q.Type().Elem())
		q.Set(v)
	}
	spew.Dump(q)
	fmt.Printf("-------- isNil=%v\n", q.IsNil())
	ii := q.Interface()
	spew.Dump(ii)
	fmt.Printf("-----------------------\n")

	asProtoMsg := ii.(proto.Message)
	spew.Dump(asProtoMsg)
	err := pbUnmarshaler.Unmarshal(bytes.NewReader(data), asProtoMsg)
	if err != nil {
		return fmt.Errorf(
			"unable to encode argument:  %v, with jsonpb error: %v", reflect.TypeOf(obj), err,
		)
	}

	fmt.Printf("\n\n[protoDataConverter.FromData()] Decoding result: %s\n\n", spew.Sdump(asProtoMsg))
	return nil
}

var (
	protoMessageIface = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

// This is a safe check of Type objects only
func isProtoMessagePtr(obj interface{}) bool {
	objType := reflect.TypeOf(obj)
	for objType.Kind() != reflect.Ptr {
		return false
	}
	if objType.Implements(protoMessageIface) {
		return true
	}
	fmt.Printf("\n\nNot a proto: %v\n\n", spew.Sdump(obj))

	return false
}

func isProtoMessagePtrPtr(obj interface{}) bool {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return false
	}
	if objType.Elem().Implements(protoMessageIface) {
		return true
	}
	return false
}

// This digs into values and returns the proto.Message that we know is lurking
func toProtoMessageObject(obj interface{}) proto.Message {
	return reflect.ValueOf(obj).Interface().(proto.Message)
}

func findConcreteType(obj interface{}) reflect.Type {
	typ := reflect.TypeOf(obj)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}
