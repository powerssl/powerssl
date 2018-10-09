package resource

import (
	"time"

	"github.com/gogo/protobuf/types"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
)

func DecodeGRPCObjectMeta(objectMeta *apiv1.ObjectMeta) (api.ObjectMeta, error) {
	creationTimestamp, err := types.TimestampFromProto(objectMeta.GetCreationTimestamp())
	if err != nil {
		return api.ObjectMeta{}, err
	}

	var deletionTimestamp time.Time
	if objectMeta.GetDeletionTimestamp() != nil {
		deletionTimestamp, err = types.TimestampFromProto(objectMeta.GetDeletionTimestamp())
		if err != nil {
			return api.ObjectMeta{}, err
		}
	}

	return api.ObjectMeta{
		CreationTimestamp: creationTimestamp,
		DeletionTimestamp: &deletionTimestamp,
		Labels:            objectMeta.GetLabels(),
		Name:              objectMeta.GetName(),
		ResourceVersion:   objectMeta.GetResourceVersion(),
		UID:               objectMeta.GetUid(),
	}, nil
}

func DecodeGRPCTypeMeta(typeMeta *apiv1.TypeMeta) (api.TypeMeta, error) {
	return api.TypeMeta{
		APIVersion: typeMeta.GetApiVersion(),
		Kind:       typeMeta.GetKind(),
	}, nil
}

func EncodeGRPCObjectMeta(objectMeta api.ObjectMeta) (*apiv1.ObjectMeta, error) {
	creationTimestamp, err := types.TimestampProto(objectMeta.CreationTimestamp)
	if err != nil {
		return nil, err
	}

	var deletionTimestamp *types.Timestamp
	if objectMeta.DeletionTimestamp != nil {
		deletionTimestamp, err = types.TimestampProto(*objectMeta.DeletionTimestamp)
		if err != nil {
			return nil, err
		}
	}

	return &apiv1.ObjectMeta{
		CreationTimestamp: creationTimestamp,
		DeletionTimestamp: deletionTimestamp,
		Labels:            objectMeta.Labels,
		Name:              objectMeta.Name,
		ResourceVersion:   objectMeta.ResourceVersion,
		Uid:               objectMeta.UID,
	}, nil
}

func EncodeGRPCTypeMeta(typeMeta api.TypeMeta) (*apiv1.TypeMeta, error) {
	return &apiv1.TypeMeta{
		ApiVersion: typeMeta.APIVersion,
		Kind:       typeMeta.Kind,
	}, nil
}
