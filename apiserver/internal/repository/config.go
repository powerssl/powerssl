package repository

type Config struct {
	ConnString string `flag:"connString;;;db conn string" validate:"required"`
}
