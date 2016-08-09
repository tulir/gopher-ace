package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// Function and variable names
const (
	Ace = "ace"
)

// Edit turns the object with the given ID into an Ace editor
func Edit(id string) Editor {
	return Editor{js.Global.Get(Ace).Call("edit", id)}
}

// EditDOM turns the given object into an Ace editor
func EditDOM(obj interface{}) Editor {
	return Editor{js.Global.Get(Ace).Call("edit", obj)}
}
