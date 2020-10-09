package balbi;

import "testing"

func TestGreeting(t *testing.T) {
	if (Greeting("Code.education Rocks!") != "<b>Code.education Rocks!</b>") {
		t.Error("Falha, esperava o texto: Code.education Rocks! dentro das tags <b></b>")
	}
}