package moves

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"reflect"
)

//The interface that moves that can be handled by DefaultConfig implement.
type defaultConfigMoveType interface {
	//The name for the move type
	MoveTypeName(manager *boardgame.GameManager) string
	//The name for the HelpText
	MoveTypeHelpText(manager *boardgame.GameManager) string
	//Whether the move should be a fix up.
	MoveTypeIsFixUp(manager *boardgame.GameManager) bool
}

//MustDefaultConfig is a wrapper around DefaultConfig that if it errors will
//panic. Only suitable for being used during setup.
func MustDefaultConfig(manager *boardgame.GameManager, exampleStruct boardgame.Move) *boardgame.MoveTypeConfig {
	result, err := DefaultConfig(manager, exampleStruct)

	if err != nil {
		panic("Couldn't DefaultConfig: " + err.Error())
	}

	return result
}

//DefaultConfig is a powerful default MoveTypeConfig generator. In many cases
//you'll implement moves that are very thin embeddings of moves in this
//package. Generating a MoveTypeConfig for each is a pain. This method auto-
//generates the MoveTypeConfig based on an example zero type of your move to
//install. It does some magic to create a more fleshed out move, then consults
//move.MoveTypeName and move.MoveTypeHelpText to generate the name and
//helptext. Moves in this package return reasonable values for those methods,
//based on the configuration you set on the rest of your move. See the package
//doc for an example of use.
func DefaultConfig(manager *boardgame.GameManager, exampleStruct boardgame.Move) (*boardgame.MoveTypeConfig, error) {

	if exampleStruct == nil {
		return nil, errors.New("nil struct provided")
	}

	//We'll create a throw-away move type config first to get a fully-
	//initialized and expanded move (e.g. with all tag-based autoinflation)
	//that we can then pass to the MoveType* methods, so they'll have more to work with.

	throwAwayConfig := newMoveTypeConfig("Temporary Move", "Temporary Move Help Text", false, exampleStruct)

	throwAwayMoveType, err := throwAwayConfig.NewMoveType(manager)

	if err != nil {
		return nil, errors.New("Couldn't create temporary move type: " + err.Error())
	}

	actualExample := throwAwayMoveType.NewMove(manager.ExampleState())

	defaultConfig, ok := actualExample.(defaultConfigMoveType)

	if !ok {
		return nil, errors.New("Example struct didn't have MoveTypeName and MoveTypeHelpText.")
	}

	name := defaultConfig.MoveTypeName(manager)
	helpText := defaultConfig.MoveTypeHelpText(manager)
	isFixUp := defaultConfig.MoveTypeIsFixUp(manager)

	return newMoveTypeConfig(name, helpText, isFixUp, exampleStruct), nil

}

func newMoveTypeConfig(name, helpText string, isFixUp bool, exampleStruct boardgame.Move) *boardgame.MoveTypeConfig {
	val := reflect.ValueOf(exampleStruct)

	//We can accept either pointer or struct types.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	return &boardgame.MoveTypeConfig{
		Name:     name,
		HelpText: helpText,
		MoveConstructor: func() boardgame.Move {
			return reflect.New(typ).Interface().(boardgame.Move)
		},
		IsFixUp: isFixUp,
	}
}
