//go:generate gobin -m -run github.com/go-bindata/go-bindata/go-bindata -fs -modtime 726710400 -pkg migration -prefix ../../db/migrations ../../db/migrations

package migration
