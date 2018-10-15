package cmd

import (
	"reflect"
)

func createResource(createFunc func() (interface{}, error)) {
	resource, err := createFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
}

func deleteResource(deleteFunc func() error) {
	if err := deleteFunc(); err != nil {
		er(err)
	}
}

func getResource(getFunc func() (interface{}, error)) {
	resource, err := getFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
}

func listResource(listFunc func(pageToken string) (interface{}, string, error)) {
	var (
		pageToken string
		resources []interface{}
	)
	for {
		t, nextPageToken, err := listFunc(pageToken)
		if err != nil {
			er(err)
		}

		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)

			for i := 0; i < s.Len(); i++ {
				resources = append(resources, s.Index(i).Interface())
			}
		}

		if nextPageToken == "" {
			break
		}
		pageToken = nextPageToken
	}
	pr(resources)
}

func updateResource(updateFunc func() (interface{}, error)) {
	resource, err := updateFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
}
