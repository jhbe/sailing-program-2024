package main

type Class int

const (
	IOM Class = iota
	RM
	TenRater
	AClass
)

func (c *Class) Next() {
	switch *c {
	case IOM:
		*c = RM
	case RM:
		*c = TenRater
	case TenRater:
		*c = AClass
	case AClass:
		*c = IOM
	}
}

func (c Class) ToString() string {
	switch c {
	case IOM:
		return "IOM"
	case RM:
		return "RM"
	case TenRater:
		return "10R"
	}
	return "A-Class"
}
