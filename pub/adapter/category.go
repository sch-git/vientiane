package adapter

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var (
	Cthulhu = &Category{Id: CategoryCthulhuId, Name: CategoryCthulhuName}
)

const (
	CategoryCthulhuId = iota + 1

	CategoryCthulhuName = "克苏鲁"
)
