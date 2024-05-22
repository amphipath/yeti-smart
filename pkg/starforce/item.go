package starforce

type Item interface {
	GetRoundedLevel() int
}

type item struct {
	level int
}

func NewItem(level int) Item {
	return &item{level: level}
}

func (i *item) GetRoundedLevel() int {
	return int((i.level + 2) / 10 * 10)
}