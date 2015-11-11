package toml

import (
	"bytes"
	"os"
	"testing"
)

func TestDefaults(t *testing.T) {
	data := struct {
		Name    []string `comment:"Name comment" default:"['one', 'two']"`
		Surname string   `comment:"Surname comment"`
	}{[]string{"one", "tthree"}, "Yakutkin"}

	var Buffer bytes.Buffer
	e := NewEncoder(&Buffer)
	e.Encode(data)

	_, err := os.Stderr.Write(Buffer.Bytes())
	if err != nil {
		panic("Cant write to stdout: " + err.Error())
	}
}
