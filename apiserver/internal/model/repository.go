package model

import (
	"fmt"
	"reflect"

	fieldmaskutils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/genproto/protobuf/field_mask"
)

func setUpdateParams(fm *field_mask.FieldMask, src interface{}, dst interface{}) (err error) {
	var mask fieldmaskutils.Mask
	if mask, err = fieldmaskutils.MaskFromProtoFieldMask(fm, camelCase); err != nil {
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

// copied from github.com/golang/protobuf/protoc-gen-go/generator

// camelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func camelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
