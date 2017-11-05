package moves

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/moves/moveinterfaces"
)

//ShuffleStack is a move, typically used in SetUp phases, that simply shuffles
//a given stack. The struct you embed this in should implement SourceStacker.
type ShuffleStack struct {
	Base
}

//We don't need a Legal method because the pass-through to moves.Base is sufficient.

func (s *ShuffleStack) Apply(state boardgame.MutableState) error {
	embeddingMove := s.TopLevelStruct()

	stacker, ok := embeddingMove.(moveinterfaces.SourceStacker)

	if !ok {
		return errors.New("Embedding move unexpectedly did not implement SourceStacker")
	}

	stack := stacker.SourceStack(state)

	if stack == nil {
		return errors.New("PrimaryStack returned a nil stack")
	}

	return stack.Shuffle()
}

func (s *ShuffleStack) ValidConfiguration() error {
	testMove := s.TopLevelStruct()

	_, ok := testMove.(moveinterfaces.SourceStacker)

	if !ok {
		return errors.New("The embedding Move doesn't implement SourceStacker")
	}

	return nil
}
