package main

import (
	"fmt"
	"strings"
)

//Payload struct that holds the name of the payload
type Payload struct {
	name string
}

//Add adds a new value to the payload string
func (p *Payload) Add(parameterName string, parameter string) {
	p.name = fmt.Sprintf("%s&%s=%s", p.name, parameterName, parameter)
}

//GetPayload converts the string to a Reader
func (p *Payload) GetPayload() *strings.Reader {
	return strings.NewReader(p.name)
}
