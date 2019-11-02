package edm

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testResource struct {
	Resource
	CustomAttr string
}

func (t testResource) Kind() string {
	return "test-resource"
}

func TestResourceIntegration(t *testing.T) {
	testCustomAttr := "Test value"
	instance := &testResource{
		CustomAttr: testCustomAttr,
	}

	assert.Equal(t, instance.Id, "")
	assert.Equal(t, instance.InternalUID, "")
	assert.Equal(t, instance.Name, "")
	assert.Equal(t, instance.Parent, "")
	assert.Equal(t, instance.CreateTime, time.Time{})
	assert.Equal(t, instance.UpdateTime, time.Time{})
	assert.Equal(t, instance.CustomAttr, testCustomAttr)
}

func TestResourceInterfaceImplementation(t *testing.T) {
	assert.Implements(t, (*ResourceInterface)(nil), &testResource{})
}

func TestIsPersisted(t *testing.T) {
	t.Run("CreateTime not present", func(t *testing.T) {
		i := &testResource{}
		i.InternalUID = NewUID()

		assert.False(t, i.IsPersisted())
	})

	t.Run("InternalUID and CreateTime present", func(t *testing.T) {
		i := &testResource{}
		i.InternalUID = NewUID()
		i.CreateTime = time.Now()

		assert.True(t, i.IsPersisted())
	})
}

func TestSetCreateTime(t *testing.T) {
	t.Run("IsPersisted true", func(t *testing.T) {
		i := &testResource{}
		i.InternalUID = NewUID()
		i.CreateTime = time.Now()

		err := i.SetCreateTime(time.Now())
		assert.EqualError(t, err, "SetCreateTime already persisted (run ResetPersistenceValues() first)")
	})

	t.Run("IsPersisted false", func(t *testing.T) {
		newTime := time.Now()
		i := &testResource{}

		err := i.SetCreateTime(newTime)
		if assert.NoError(t, err) {
			assert.Equal(t, newTime, i.CreateTime)
		}
	})
}

func TestSetInternalUID(t *testing.T) {
	t.Run("IsPersisted true", func(t *testing.T) {
		i := &testResource{}
		i.InternalUID = NewUID()
		i.CreateTime = time.Now()

		err := i.SetInternalUID(NewUID())
		assert.EqualError(t, err, "InternalUID already persisted (run ResetPersistenceValues() first)")
	})

	t.Run("IsPersisted false", func(t *testing.T) {
		newUID := NewUID()
		i := &testResource{}

		err := i.SetInternalUID(newUID)
		if assert.NoError(t, err) {
			assert.Equal(t, newUID, i.InternalUID)
			assert.True(t, i.PersistenceAttrsChanged)
		}
	})
}

func TestResourceFullKey(t *testing.T) {
	i := &testResource{}
	i.Name = "hello world"

	fmt.Println(i)

	matched, _ := regexp.MatchString(`^\/([a-z0-9]{1,1})([a-z0-9-_]{0,})*[a-z0-9]$`, "/ab-c")
	fmt.Println(matched)
	// ^\/(?:[a-z0-9]{1,1})([a-z0-9-_]{0,})
}
