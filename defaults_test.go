package toml

import (
	"bytes"
	"os"
	"testing"
)

type Person struct {
	Name    string `comment:"Name comment" default:"Dmitry"`
	Surname string `comment:"Surname comment" default:"Yakutkin"`
}

func TestDefaults(t *testing.T) {
	data := struct {
		Person Person
	}{Person: Person{"Dmitry", "Yakutkin"}}

	var Buffer bytes.Buffer
	e := NewEncoder(&Buffer)
	e.Encode(data)

	_, err := os.Stderr.Write(Buffer.Bytes())
	if err != nil {
		panic("Cant write to stdout: " + err.Error())
	}
}
