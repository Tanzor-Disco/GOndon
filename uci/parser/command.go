package parser

import ()

// Command interface is used to be able
// To return multiple structs from one parser function
type Command interface {
	isCommand()
}
