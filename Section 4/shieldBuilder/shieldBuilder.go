package shieldBuilder

import "strings"

type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shBuidler struct {
	code string
}

//constructor for a new shield builder
func NewShieldBuilder() *shBuidler {
	return new(shBuidler)
}

func (sh *shBuidler) RaiseFront() *shBuidler {
	sh.code += "F"
	return sh
}

func (sh *shBuidler) RaiseBack() *shBuidler {
	sh.code += "B"
	return sh
}

func (sh *shBuidler) RaiseRight() *shBuidler {
	sh.code += "R"
	return sh
}

func (sh *shBuidler) RaiseLeft() *shBuidler {
	sh.code += "L"
	return sh
}

func (sh *shBuidler) Build() *shield {
	code := sh.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
