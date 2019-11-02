package edm

import (
	"errors"
	"time"
)

// Base data mapper resource definition. Includes the most important and common needed fields.
type Resource struct {
	Id                      string    `json:"id,omitempty" yaml:"id,omitempty"`
	InternalUID             string    `json:"internalUID,omitempty" yaml:"internalUID,omitempty"`
	Name                    string    `json:"name,omitempty"		yaml:"name,omitempty"`
	Parent                  string    `json:"parent,omitempty"		yaml:"parent,omitempty"`
	CreateTime              time.Time `json:"createTime,omitempty"	yaml:"createTime,omitempty"`
	UpdateTime              time.Time `json:"updateTime,omitempty"	yaml:"updateTime,omitempty"`
	PersistenceAttrsChanged bool
}

// The interface which every edm usable resource needs to implement. Type Resource implements
// everything by default.
type ResourceInterface interface {
	GetCreateTime() time.Time
	GetInternalUID() string
	GetName() string
	GetParent() string
	GetUpdateTime() time.Time
	GetPersistenceAttrsChanged() bool
	IsPersisted() bool
	Kind() string
	SetCreateTime(value time.Time) error
	SetInternalUID(value string) error
}

func (r Resource) GetCreateTime() time.Time {
	return r.CreateTime
}

func (r Resource) GetInternalUID() string {
	return r.InternalUID
}

func (r Resource) GetName() string {
	return r.Name
}

func (r Resource) GetParent() string {
	return r.Parent
}

func (r Resource) GetUpdateTime() time.Time {
	return r.UpdateTime
}

func (r Resource) IsPersisted() bool {
	return r.CreateTime != time.Time{} && r.InternalUID != ""
}

func (r Resource) GetPersistenceAttrsChanged() bool {
	return r.PersistenceAttrsChanged
}

// Sets the CreateTime of the resource and raises error if it already was persisted.
func (r *Resource) SetCreateTime(value time.Time) error {
	if r.IsPersisted() {
		return errors.New("SetCreateTime already persisted (run ResetPersistenceValues() first)")
	} else {
		r.CreateTime = value
		r.PersistenceAttrsChanged = true

		return nil
	}
}

// Sets the InternalUID of the resource and raises error if it already was persisted.
func (r *Resource) SetInternalUID(value string) error {
	if r.IsPersisted() {
		return errors.New("InternalUID already persisted (run ResetPersistenceValues() first)")
	} else {
		r.InternalUID = value
		r.PersistenceAttrsChanged = true

		return nil
	}
}
