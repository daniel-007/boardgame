package boardgame

//Move's are how all modifications are made to Game States after
//initialization. Packages define structs that implement Move for all
//modifications.
type Move interface {
	//Legal returns nil if this proposed move is legal, or an error if the
	//move is not legal
	Legal(state StatePayload) error
	//TODO: figure out how to get a string describing why it's not legal out

	//Apply applies the move to the state and returns a new state object. It
	//should not be called directly; use Game.ApplyMove.
	Apply(state StatePayload) StatePayload

	//Copy creates a new move based on this one.
	Copy() Move
	GameNamer
	PropertyReader
	JSONer
}
