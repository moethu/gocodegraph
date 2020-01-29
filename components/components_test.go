package components

import "testing"

func TestInitTypeRegistry(t *testing.T) {
	InitTypeRegistry()
	if _, ok := typeRegistry["components.Addition"]; !ok {
		t.Error("Addition Component not loaded")
	}
	if _, ok := typeRegistry["components.Number"]; !ok {
		t.Error("Addition Component not loaded")
	}
	if _, ok := typeRegistry["components.Multiplication"]; !ok {
		t.Error("Addition Component not loaded")
	}
}

func TestGetComponents(t *testing.T) {
	InitTypeRegistry()
	c := GetComponents()
	if len(c) == 0 {
		t.Error("Could not get any components")
	}
}
