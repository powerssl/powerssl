package transport

import (
	"time"

	"powerssl.io/pkg/api"
	pb "powerssl.io/pkg/api/v1"
)

func DecodeGRPCObjectMeta(objectMeta *pb.ObjectMeta) api.ObjectMeta {
	return api.ObjectMeta{
		CreationTimestamp: time.Time{}, // TODO: objectMeta.CreationTimestamp,
		DeletionTimestamp: nil,         // TODO: objectMeta.DeletionTimestamp,
		Labels:            objectMeta.Labels,
		Name:              objectMeta.Name,
		ResourceVersion:   objectMeta.ResourceVersion,
		UID:               objectMeta.Uid,
	}
}

func DecodeGRPCTypeMeta(typeMeta *pb.TypeMeta) api.TypeMeta {
	return api.TypeMeta{
		APIVersion: typeMeta.ApiVersion,
		Kind:       typeMeta.Kind,
	}
}

func EncodeGRPCObjectMeta(objectMeta api.ObjectMeta) *pb.ObjectMeta {
	return &pb.ObjectMeta{
		CreationTimestamp: nil, // TODO: objectMeta.CreationTimestamp,
		DeletionTimestamp: nil, // TODO: objectMeta.DeletionTimestamp,
		Labels:            objectMeta.Labels,
		Name:              objectMeta.Name,
		ResourceVersion:   objectMeta.ResourceVersion,
		Uid:               objectMeta.UID,
	}
}

func EncodeGRPCTypeMeta(typeMeta api.TypeMeta) *pb.TypeMeta {
	return &pb.TypeMeta{
		ApiVersion: typeMeta.APIVersion,
		Kind:       typeMeta.Kind,
	}
}
