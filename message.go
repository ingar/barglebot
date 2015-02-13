package barglebot

type Message interface {
	Text() string
	Respond(string)
	DebugDump() string
}
