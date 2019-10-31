package prototest
//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io"
//	"reflect"
//
//	"github.com/davecgh/go-spew/spew"
//	"github.com/golang/protobuf/jsonpb"
//	"github.com/golang/protobuf/proto"
//	"go.uber.org/cadence/encoded"
//)
//
//var (
//	pbMarshaler   = jsonpb.Marshaler{}
//	pbUnmarshaler = jsonpb.Unmarshaler{}
//)
//
//type CustomDataConverter struct{}
//
//func (dc *CustomDataConverter) ToData(objs ...interface{}) ([]byte, error) {
//	result := &bytes.Buffer{}
//	for i, obj := range objs {
//		// Try jsonpb...
//		if isProtoMessage(obj) {
//			fmt.Printf("\n\n[CustomDataConverter.ToData()] Encoding a proto.Message: %T\n\n", obj)
//			curBytes := &bytes.Buffer{}
//			pbObj := findProtoMessageObject(obj)
//			if pbObj == nil {
//				return nil, fmt.Errorf("missing proto.Message argument at index %d of type %T", i, obj)
//			}
//			err := pbMarshaler.Marshal(curBytes, pbObj)
//			if err != nil {
//				if err == io.EOF {
//					return nil, fmt.Errorf("missing argument at index %d of type %T", i, obj)
//				}
//				return nil, fmt.Errorf(
//					"unable to encode argument: %d, %v, with jsonpb error: %v", i, reflect.TypeOf(obj), err,
//				)
//			}
//			// fmt.Printf("\n\n[CustomDataConverter.ToData()] jsonpb wrote bytes:\n%s\n\n", curBytes.String())
//			_, _ = io.Copy(result, curBytes)
//		} else {
//			// ... falling back to the default DC
//			fmt.Printf("\n\n[CustomDataConverter.ToData()] Encoding a normal type: %T\n\n", obj)
//			curBytes, err := encoded.GetDefaultDataConverter().ToData(obj)
//			if err != nil {
//				return nil, err
//			}
//			_, _ = io.Copy(result, bytes.NewReader(curBytes))
//		}
//	}
//	return result.Bytes(), nil
//}
//
//func (dc *CustomDataConverter) FromData(data []byte, objs ...interface{}) error {
//	dec := json.NewDecoder(bytes.NewBuffer(data))
//	for i, obj := range objs {
//		var intermediate json.RawMessage
//		if err := dec.Decode(&intermediate); err != nil {
//			return fmt.Errorf(
//				"unable to decode argument: %d, %v, with json error: %v", i, reflect.TypeOf(obj), err,
//			)
//		}
//
//		// fmt.Printf("\n\n[CustomDataConverter.FromData()] Ready to decode bytes:\n%s\n\n", string(intermediate))
//
//		// Try jsonpb...
//		if isProtoMessage(obj) {
//			fmt.Printf("\n\n[CustomDataConverter.FromData()] Decoding a proto.Message: %T\n\n", obj)
//
//			// Find the proto.Message!
//			q := reflect.ValueOf(obj)
//			for !q.Type().Implements(protoMessageIface) {
//				q = q.Elem()
//			}
//
//			// So now q is a (Value: *SomeStruct)
//			// which implements proto.Message
//
//			// If our proto.Message points to nil, we need to create a new
//			// zero value object, or jsonpb will freak out during unmarshal
//			if q.IsNil() {
//				concTyp :=findConcreteType(obj)
//				blankObjVal := reflect.Zero(concTyp)
//				blankObjPtrVal := reflect.New(concTyp)
//				blankObjPtrVal.Elem().Set(blankObjVal)
//				asIface := q.Interface().(proto.Message)
//				qq := reflect.ValueOf(&asIface).Elem()
//				qq.Set(blankObjPtrVal) // HELP! This isn't setting q to point to the new blank object! Why not?!
//			}
//
//			asProtoMsg := q.Interface().(proto.Message)
//			spew.Dump(asProtoMsg)
//			err := pbUnmarshaler.Unmarshal(bytes.NewReader(intermediate), asProtoMsg)
//			if err != nil {
//				return fmt.Errorf(
//					"unable to encode argument: %d, %v, with jsonpb error: %v", i, reflect.TypeOf(obj), err,
//				)
//			}
//
//			fmt.Printf("\n\n[CustomDataConverter.FromData()] Decoding result: %s\n\n", spew.Sdump(asProtoMsg))
//
//		} else {
//			// ... falling back to the default DC
//			fmt.Printf("\n\n[CustomDataConverter.FromData()] Decoding to a normal type: %T\n\n", obj)
//			err := encoded.GetDefaultDataConverter().FromData(intermediate, obj)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//var (
//	protoMessageIface = reflect.TypeOf((*proto.Message)(nil)).Elem()
//)
//
//// This is a safe check of Type objects only
//func isProtoMessage(obj interface{}) bool {
//	curTyp := reflect.TypeOf(obj)
//	prevTyp := curTyp
//	for curTyp.Kind() == reflect.Ptr {
//		prevTyp = curTyp
//		curTyp = curTyp.Elem()
//	}
//
//	if prevTyp.Implements(protoMessageIface) {
//		return true
//	}
//
//	return false
//}
//
//// This digs into values and returns the proto.Message that we know is lurking
//func findProtoMessageObject(obj interface{}) proto.Message {
//	// Strip all pointers away
//	curObj := obj
//	prevObj := obj
//	for reflect.TypeOf(curObj).Kind() == reflect.Ptr {
//		prevObj = curObj
//		if reflect.ValueOf(curObj).IsNil() {
//			break
//		}
//		curObj = reflect.ValueOf(curObj).Elem().Interface()
//	}
//
//	// prevObj should be a pointer to the concrete type
//	if reflect.TypeOf(prevObj).Implements(protoMessageIface) {
//		return prevObj.(proto.Message)
//	}
//
//	return nil
//}
//
//func findConcreteType(obj interface{}) reflect.Type {
//	typ := reflect.TypeOf(obj)
//	for typ.Kind() == reflect.Ptr {
//		typ = typ.Elem()
//	}
//	return typ
//}
