package category

type category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
