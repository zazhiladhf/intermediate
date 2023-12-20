package category

import "errors"

var (
	ErrRepository     = errors.New("error repository")
	ErrInternalServer = errors.New("unknown error")
)

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
