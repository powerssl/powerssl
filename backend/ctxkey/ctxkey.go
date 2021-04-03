package ctxkey // import "powerssl.dev/backend/ctexkey"

import "fmt"

type Key interface {
	Name() string
}

func New(name string) Key {
	return &key{
		name: name,
	}
}

type key struct {
	name string
}

func (k *key) Name() string {
	return k.name
}

func (k key) String() string {
	return fmt.Sprintf("Key(%s)", k.Name())
}
