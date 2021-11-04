package repository

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type UpdateMaskSanitizer struct {
	m        proto.Message
	allowed  *fieldmaskpb.FieldMask
	internal *fieldmaskpb.FieldMask
}

func NewUpdateMaskSanitizer(m proto.Message) *UpdateMaskSanitizer {
	return &UpdateMaskSanitizer{m: m}
}

func (s *UpdateMaskSanitizer) Allowed(paths ...string) *UpdateMaskSanitizer {
	var err error
	if s.allowed, err = fieldmaskpb.New(s.m, paths...); err != nil {
		panic(err)
	}
	return s
}

func (s *UpdateMaskSanitizer) Internal(paths ...string) *UpdateMaskSanitizer {
	var err error
	if s.internal, err = fieldmaskpb.New(s.m, paths...); err != nil {
		panic(err)
	}
	return s
}

func (s *UpdateMaskSanitizer) Sanitize(fm *fieldmaskpb.FieldMask, internal bool) *fieldmaskpb.FieldMask {
	if internal {
		return fieldmaskpb.Intersect(fm, s.allowed, s.internal)
	}
	return fieldmaskpb.Intersect(fm, s.allowed)
}
