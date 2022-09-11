package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Human struct {
	Age  int
	Name string
}
type Action struct {
	Human Human // встраивание
}

func (h *Human) AnswerOnQuestionByThemeWhoAreYou() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("I'm ")
	buffer.WriteString(h.Name)
	buffer.WriteString(" and I'm ")
	buffer.WriteString(strconv.Itoa(h.Age))
	buffer.WriteString(" years old")
	return buffer.String()
}

func main() {
	action := Action{Human: Human{60, "Qwerty"}}
	fmt.Println(action.Human.AnswerOnQuestionByThemeWhoAreYou())
}
