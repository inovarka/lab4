package engine

// Command represents actions that can be performed in a single event loop iteration.
type Command interface {
	Execute(handler Handler)
}
	// Handler allows to send commands to an event loop it's associated with.
type Handler interface {
	Post(cmd Command)
} 
	