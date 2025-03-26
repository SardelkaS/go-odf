package odt

type Element interface {
	GetStyle() string
	GetContent() string
}
