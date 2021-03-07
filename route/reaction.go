package route

type ReactionEvent uint8

const (
	ReactionAdd ReactionEvent = iota
	ReactionRemove
)

type Reaction struct {
	Name  string
	Run   ReactionRunFunc
	Event ReactionEvent
}
