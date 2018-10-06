package transport

import (
	"github.com/gogo/protobuf/types"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
)

// TODO: Error handling
func DecodeGRPCObjectMeta(objectMeta *apiv1.ObjectMeta) api.ObjectMeta {
	creationTimestamp, _ := types.TimestampFromProto(objectMeta.GetCreationTimestamp())
	deletionTimestamp, _ := types.TimestampFromProto(objectMeta.GetDeletionTimestamp())

	return api.ObjectMeta{
		CreationTimestamp: creationTimestamp,
		DeletionTimestamp: &deletionTimestamp,
		Labels:            objectMeta.GetLabels(),
		Name:              objectMeta.GetName(),
		ResourceVersion:   objectMeta.GetResourceVersion(),
		UID:               objectMeta.GetUid(),
	}
}

func DecodeGRPCTypeMeta(typeMeta *apiv1.TypeMeta) api.TypeMeta {
	return api.TypeMeta{
		APIVersion: typeMeta.GetApiVersion(),
		Kind:       typeMeta.GetKind(),
	}
}

// TODO: Error handling
func EncodeGRPCObjectMeta(objectMeta api.ObjectMeta) *apiv1.ObjectMeta {
	creationTimestamp, _ := types.TimestampProto(objectMeta.CreationTimestamp)
	var deletionTimestamp *types.Timestamp
	if objectMeta.DeletionTimestamp != nil {
		deletionTimestamp, _ = types.TimestampProto(*objectMeta.DeletionTimestamp)
	}

	return &apiv1.ObjectMeta{
		CreationTimestamp: creationTimestamp,
		DeletionTimestamp: deletionTimestamp,
		Labels:            objectMeta.Labels,
		Name:              objectMeta.Name,
		ResourceVersion:   objectMeta.ResourceVersion,
		Uid:               objectMeta.UID,
	}
}

func EncodeGRPCTypeMeta(typeMeta api.TypeMeta) *apiv1.TypeMeta {
	return &apiv1.TypeMeta{
		ApiVersion: typeMeta.APIVersion,
		Kind:       typeMeta.Kind,
	}
}
