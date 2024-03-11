package entity

type City struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
type PropertyType struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
type ConfectionaryType struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
type Unit struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
type Packaging struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
type District struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	City `db:"city"`
}
