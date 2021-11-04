package repository

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	fieldmaskutils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func setUpdateParams(fm *fieldmaskpb.FieldMask, src interface{}, dst interface{}) (err error) {
	var mask fieldmaskutils.Mask
	if mask, err = fieldmaskutils.MaskFromProtoFieldMask(fm, generator.CamelCase); err != nil {
		return err
	}
	if err = fieldmaskutils.StructToStruct(mask, src, dst); err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			if err != nil {
				err = fmt.Errorf("%w: %v", err, e)
			} else {
				err = fmt.Errorf("%v", e)
			}
		}
	}()

	ps := reflect.ValueOf(dst)
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("kind of element is not a struct")
	}
	for name := range mask {
		name = fmt.Sprintf("Set%s", name)
		f := s.FieldByName(name)
		if !f.IsValid() {
			return fmt.Errorf("field %v is invalid", name)
		}
		if !f.CanSet() {
			return fmt.Errorf("field %v connot be set", name)
		}
		if f.Kind() == reflect.Bool {
			return fmt.Errorf("field %v kind is not bool", name)
		}
		f.SetBool(true)
	}
	return nil
}
