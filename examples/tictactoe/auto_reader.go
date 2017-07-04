/************************************
 *
 * This file contains auto-generated methods to help certain structs
 * implement boardgame.SubState and boardgame.MutableSubState. It was
 * generated by autoreader.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/
package tictactoe

import (
	"errors"
	"github.com/jkomoros/boardgame"
)

// Implementation for playerToken

var __playerTokenReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"Value": boardgame.TypeString,
}

type __playerTokenReader struct {
	data *playerToken
}

func (p *__playerTokenReader) Props() map[string]boardgame.PropertyType {
	return __playerTokenReaderProps
}

func (p *__playerTokenReader) Prop(name string) (interface{}, error) {
	props := p.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return p.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return p.BoolSliceProp(name)
	case boardgame.TypeGrowableStack:
		return p.GrowableStackProp(name)
	case boardgame.TypeInt:
		return p.IntProp(name)
	case boardgame.TypeIntSlice:
		return p.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return p.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return p.PlayerIndexSliceProp(name)
	case boardgame.TypeSizedStack:
		return p.SizedStackProp(name)
	case boardgame.TypeString:
		return p.StringProp(name)
	case boardgame.TypeStringSlice:
		return p.StringSliceProp(name)
	case boardgame.TypeTimer:
		return p.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (p *__playerTokenReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (p *__playerTokenReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (p *__playerTokenReader) GrowableStackProp(name string) (*boardgame.GrowableStack, error) {

	return nil, errors.New("No such GrowableStack prop: " + name)

}

func (p *__playerTokenReader) IntProp(name string) (int, error) {

	return 0, errors.New("No such Int prop: " + name)

}

func (p *__playerTokenReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (p *__playerTokenReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (p *__playerTokenReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (p *__playerTokenReader) SizedStackProp(name string) (*boardgame.SizedStack, error) {

	return nil, errors.New("No such SizedStack prop: " + name)

}

func (p *__playerTokenReader) StringProp(name string) (string, error) {

	switch name {
	case "Value":
		return p.data.Value, nil

	}

	return "", errors.New("No such String prop: " + name)

}

func (p *__playerTokenReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (p *__playerTokenReader) TimerProp(name string) (*boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (p *playerToken) Reader() boardgame.PropertyReader {
	return &__playerTokenReader{p}
}

// Implementation for MovePlaceToken

var __MovePlaceTokenReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"Slot":              boardgame.TypeInt,
	"TargetPlayerIndex": boardgame.TypePlayerIndex,
}

type __MovePlaceTokenReader struct {
	data *MovePlaceToken
}

func (m *__MovePlaceTokenReader) Props() map[string]boardgame.PropertyType {
	return __MovePlaceTokenReaderProps
}

func (m *__MovePlaceTokenReader) Prop(name string) (interface{}, error) {
	props := m.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return m.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return m.BoolSliceProp(name)
	case boardgame.TypeGrowableStack:
		return m.GrowableStackProp(name)
	case boardgame.TypeInt:
		return m.IntProp(name)
	case boardgame.TypeIntSlice:
		return m.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return m.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return m.PlayerIndexSliceProp(name)
	case boardgame.TypeSizedStack:
		return m.SizedStackProp(name)
	case boardgame.TypeString:
		return m.StringProp(name)
	case boardgame.TypeStringSlice:
		return m.StringSliceProp(name)
	case boardgame.TypeTimer:
		return m.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (m *__MovePlaceTokenReader) SetProp(name string, value interface{}) error {
	props := m.Props()
	propType, ok := props[name]

	if !ok {
		return errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		val, ok := value.(bool)
		if !ok {
			return errors.New("Provided value was not of type bool")
		}
		return m.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return m.SetBoolSliceProp(name, val)
	case boardgame.TypeGrowableStack:
		val, ok := value.(*boardgame.GrowableStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.GrowableStack")
		}
		return m.SetGrowableStackProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return m.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return m.SetIntSliceProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return m.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return m.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeSizedStack:
		val, ok := value.(*boardgame.SizedStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.SizedStack")
		}
		return m.SetSizedStackProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return m.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return m.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(*boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.Timer")
		}
		return m.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (m *__MovePlaceTokenReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (m *__MovePlaceTokenReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (m *__MovePlaceTokenReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) GrowableStackProp(name string) (*boardgame.GrowableStack, error) {

	return nil, errors.New("No such GrowableStack prop: " + name)

}

func (m *__MovePlaceTokenReader) SetGrowableStackProp(name string, value *boardgame.GrowableStack) error {

	return errors.New("No such GrowableStack prop: " + name)

}

func (m *__MovePlaceTokenReader) IntProp(name string) (int, error) {

	switch name {
	case "Slot":
		return m.data.Slot, nil

	}

	return 0, errors.New("No such Int prop: " + name)

}

func (m *__MovePlaceTokenReader) SetIntProp(name string, value int) error {

	switch name {
	case "Slot":
		m.data.Slot = value
		return nil

	}

	return errors.New("No such Int prop: " + name)

}

func (m *__MovePlaceTokenReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	switch name {
	case "TargetPlayerIndex":
		return m.data.TargetPlayerIndex, nil

	}

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (m *__MovePlaceTokenReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	switch name {
	case "TargetPlayerIndex":
		m.data.TargetPlayerIndex = value
		return nil

	}

	return errors.New("No such PlayerIndex prop: " + name)

}

func (m *__MovePlaceTokenReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) SizedStackProp(name string) (*boardgame.SizedStack, error) {

	return nil, errors.New("No such SizedStack prop: " + name)

}

func (m *__MovePlaceTokenReader) SetSizedStackProp(name string, value *boardgame.SizedStack) error {

	return errors.New("No such SizedStack prop: " + name)

}

func (m *__MovePlaceTokenReader) StringProp(name string) (string, error) {

	return "", errors.New("No such String prop: " + name)

}

func (m *__MovePlaceTokenReader) SetStringProp(name string, value string) error {

	return errors.New("No such String prop: " + name)

}

func (m *__MovePlaceTokenReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (m *__MovePlaceTokenReader) TimerProp(name string) (*boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (m *__MovePlaceTokenReader) SetTimerProp(name string, value *boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (m *MovePlaceToken) ReadSetter() boardgame.PropertyReadSetter {
	return &__MovePlaceTokenReader{m}
}

// Implementation for MoveFinishTurn

var __MoveFinishTurnReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{}

type __MoveFinishTurnReader struct {
	data *MoveFinishTurn
}

func (m *__MoveFinishTurnReader) Props() map[string]boardgame.PropertyType {
	return __MoveFinishTurnReaderProps
}

func (m *__MoveFinishTurnReader) Prop(name string) (interface{}, error) {
	props := m.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return m.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return m.BoolSliceProp(name)
	case boardgame.TypeGrowableStack:
		return m.GrowableStackProp(name)
	case boardgame.TypeInt:
		return m.IntProp(name)
	case boardgame.TypeIntSlice:
		return m.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return m.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return m.PlayerIndexSliceProp(name)
	case boardgame.TypeSizedStack:
		return m.SizedStackProp(name)
	case boardgame.TypeString:
		return m.StringProp(name)
	case boardgame.TypeStringSlice:
		return m.StringSliceProp(name)
	case boardgame.TypeTimer:
		return m.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (m *__MoveFinishTurnReader) SetProp(name string, value interface{}) error {
	props := m.Props()
	propType, ok := props[name]

	if !ok {
		return errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		val, ok := value.(bool)
		if !ok {
			return errors.New("Provided value was not of type bool")
		}
		return m.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return m.SetBoolSliceProp(name, val)
	case boardgame.TypeGrowableStack:
		val, ok := value.(*boardgame.GrowableStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.GrowableStack")
		}
		return m.SetGrowableStackProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return m.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return m.SetIntSliceProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return m.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return m.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeSizedStack:
		val, ok := value.(*boardgame.SizedStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.SizedStack")
		}
		return m.SetSizedStackProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return m.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return m.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(*boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.Timer")
		}
		return m.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (m *__MoveFinishTurnReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (m *__MoveFinishTurnReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (m *__MoveFinishTurnReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) GrowableStackProp(name string) (*boardgame.GrowableStack, error) {

	return nil, errors.New("No such GrowableStack prop: " + name)

}

func (m *__MoveFinishTurnReader) SetGrowableStackProp(name string, value *boardgame.GrowableStack) error {

	return errors.New("No such GrowableStack prop: " + name)

}

func (m *__MoveFinishTurnReader) IntProp(name string) (int, error) {

	return 0, errors.New("No such Int prop: " + name)

}

func (m *__MoveFinishTurnReader) SetIntProp(name string, value int) error {

	return errors.New("No such Int prop: " + name)

}

func (m *__MoveFinishTurnReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (m *__MoveFinishTurnReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndex prop: " + name)

}

func (m *__MoveFinishTurnReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) SizedStackProp(name string) (*boardgame.SizedStack, error) {

	return nil, errors.New("No such SizedStack prop: " + name)

}

func (m *__MoveFinishTurnReader) SetSizedStackProp(name string, value *boardgame.SizedStack) error {

	return errors.New("No such SizedStack prop: " + name)

}

func (m *__MoveFinishTurnReader) StringProp(name string) (string, error) {

	return "", errors.New("No such String prop: " + name)

}

func (m *__MoveFinishTurnReader) SetStringProp(name string, value string) error {

	return errors.New("No such String prop: " + name)

}

func (m *__MoveFinishTurnReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (m *__MoveFinishTurnReader) TimerProp(name string) (*boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (m *__MoveFinishTurnReader) SetTimerProp(name string, value *boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (m *MoveFinishTurn) ReadSetter() boardgame.PropertyReadSetter {
	return &__MoveFinishTurnReader{m}
}

// Implementation for gameState

var __gameStateReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"CurrentPlayer": boardgame.TypePlayerIndex,
	"Slots":         boardgame.TypeSizedStack,
}

type __gameStateReader struct {
	data *gameState
}

func (g *__gameStateReader) Props() map[string]boardgame.PropertyType {
	return __gameStateReaderProps
}

func (g *__gameStateReader) Prop(name string) (interface{}, error) {
	props := g.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return g.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return g.BoolSliceProp(name)
	case boardgame.TypeGrowableStack:
		return g.GrowableStackProp(name)
	case boardgame.TypeInt:
		return g.IntProp(name)
	case boardgame.TypeIntSlice:
		return g.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return g.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return g.PlayerIndexSliceProp(name)
	case boardgame.TypeSizedStack:
		return g.SizedStackProp(name)
	case boardgame.TypeString:
		return g.StringProp(name)
	case boardgame.TypeStringSlice:
		return g.StringSliceProp(name)
	case boardgame.TypeTimer:
		return g.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (g *__gameStateReader) SetProp(name string, value interface{}) error {
	props := g.Props()
	propType, ok := props[name]

	if !ok {
		return errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		val, ok := value.(bool)
		if !ok {
			return errors.New("Provided value was not of type bool")
		}
		return g.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return g.SetBoolSliceProp(name, val)
	case boardgame.TypeGrowableStack:
		val, ok := value.(*boardgame.GrowableStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.GrowableStack")
		}
		return g.SetGrowableStackProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return g.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return g.SetIntSliceProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return g.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return g.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeSizedStack:
		val, ok := value.(*boardgame.SizedStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.SizedStack")
		}
		return g.SetSizedStackProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return g.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return g.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(*boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.Timer")
		}
		return g.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (g *__gameStateReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (g *__gameStateReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (g *__gameStateReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (g *__gameStateReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (g *__gameStateReader) GrowableStackProp(name string) (*boardgame.GrowableStack, error) {

	return nil, errors.New("No such GrowableStack prop: " + name)

}

func (g *__gameStateReader) SetGrowableStackProp(name string, value *boardgame.GrowableStack) error {

	return errors.New("No such GrowableStack prop: " + name)

}

func (g *__gameStateReader) IntProp(name string) (int, error) {

	return 0, errors.New("No such Int prop: " + name)

}

func (g *__gameStateReader) SetIntProp(name string, value int) error {

	return errors.New("No such Int prop: " + name)

}

func (g *__gameStateReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (g *__gameStateReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (g *__gameStateReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	switch name {
	case "CurrentPlayer":
		return g.data.CurrentPlayer, nil

	}

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (g *__gameStateReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	switch name {
	case "CurrentPlayer":
		g.data.CurrentPlayer = value
		return nil

	}

	return errors.New("No such PlayerIndex prop: " + name)

}

func (g *__gameStateReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (g *__gameStateReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (g *__gameStateReader) SizedStackProp(name string) (*boardgame.SizedStack, error) {

	switch name {
	case "Slots":
		return g.data.Slots, nil

	}

	return nil, errors.New("No such SizedStack prop: " + name)

}

func (g *__gameStateReader) SetSizedStackProp(name string, value *boardgame.SizedStack) error {

	switch name {
	case "Slots":
		g.data.Slots = value
		return nil

	}

	return errors.New("No such SizedStack prop: " + name)

}

func (g *__gameStateReader) StringProp(name string) (string, error) {

	return "", errors.New("No such String prop: " + name)

}

func (g *__gameStateReader) SetStringProp(name string, value string) error {

	return errors.New("No such String prop: " + name)

}

func (g *__gameStateReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (g *__gameStateReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (g *__gameStateReader) TimerProp(name string) (*boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (g *__gameStateReader) SetTimerProp(name string, value *boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (g *gameState) Reader() boardgame.PropertyReader {
	return &__gameStateReader{g}
}

func (g *gameState) ReadSetter() boardgame.PropertyReadSetter {
	return &__gameStateReader{g}
}

// Implementation for playerState

var __playerStateReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"TokenValue":            boardgame.TypeString,
	"TokensToPlaceThisTurn": boardgame.TypeInt,
	"UnusedTokens":          boardgame.TypeGrowableStack,
}

type __playerStateReader struct {
	data *playerState
}

func (p *__playerStateReader) Props() map[string]boardgame.PropertyType {
	return __playerStateReaderProps
}

func (p *__playerStateReader) Prop(name string) (interface{}, error) {
	props := p.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return p.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return p.BoolSliceProp(name)
	case boardgame.TypeGrowableStack:
		return p.GrowableStackProp(name)
	case boardgame.TypeInt:
		return p.IntProp(name)
	case boardgame.TypeIntSlice:
		return p.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return p.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return p.PlayerIndexSliceProp(name)
	case boardgame.TypeSizedStack:
		return p.SizedStackProp(name)
	case boardgame.TypeString:
		return p.StringProp(name)
	case boardgame.TypeStringSlice:
		return p.StringSliceProp(name)
	case boardgame.TypeTimer:
		return p.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (p *__playerStateReader) SetProp(name string, value interface{}) error {
	props := p.Props()
	propType, ok := props[name]

	if !ok {
		return errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		val, ok := value.(bool)
		if !ok {
			return errors.New("Provided value was not of type bool")
		}
		return p.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return p.SetBoolSliceProp(name, val)
	case boardgame.TypeGrowableStack:
		val, ok := value.(*boardgame.GrowableStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.GrowableStack")
		}
		return p.SetGrowableStackProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return p.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return p.SetIntSliceProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return p.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return p.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeSizedStack:
		val, ok := value.(*boardgame.SizedStack)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.SizedStack")
		}
		return p.SetSizedStackProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return p.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return p.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(*boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type *boardgame.Timer")
		}
		return p.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (p *__playerStateReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (p *__playerStateReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (p *__playerStateReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (p *__playerStateReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (p *__playerStateReader) GrowableStackProp(name string) (*boardgame.GrowableStack, error) {

	switch name {
	case "UnusedTokens":
		return p.data.UnusedTokens, nil

	}

	return nil, errors.New("No such GrowableStack prop: " + name)

}

func (p *__playerStateReader) SetGrowableStackProp(name string, value *boardgame.GrowableStack) error {

	switch name {
	case "UnusedTokens":
		p.data.UnusedTokens = value
		return nil

	}

	return errors.New("No such GrowableStack prop: " + name)

}

func (p *__playerStateReader) IntProp(name string) (int, error) {

	switch name {
	case "TokensToPlaceThisTurn":
		return p.data.TokensToPlaceThisTurn, nil

	}

	return 0, errors.New("No such Int prop: " + name)

}

func (p *__playerStateReader) SetIntProp(name string, value int) error {

	switch name {
	case "TokensToPlaceThisTurn":
		p.data.TokensToPlaceThisTurn = value
		return nil

	}

	return errors.New("No such Int prop: " + name)

}

func (p *__playerStateReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (p *__playerStateReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (p *__playerStateReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (p *__playerStateReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndex prop: " + name)

}

func (p *__playerStateReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (p *__playerStateReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (p *__playerStateReader) SizedStackProp(name string) (*boardgame.SizedStack, error) {

	return nil, errors.New("No such SizedStack prop: " + name)

}

func (p *__playerStateReader) SetSizedStackProp(name string, value *boardgame.SizedStack) error {

	return errors.New("No such SizedStack prop: " + name)

}

func (p *__playerStateReader) StringProp(name string) (string, error) {

	switch name {
	case "TokenValue":
		return p.data.TokenValue, nil

	}

	return "", errors.New("No such String prop: " + name)

}

func (p *__playerStateReader) SetStringProp(name string, value string) error {

	switch name {
	case "TokenValue":
		p.data.TokenValue = value
		return nil

	}

	return errors.New("No such String prop: " + name)

}

func (p *__playerStateReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (p *__playerStateReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (p *__playerStateReader) TimerProp(name string) (*boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (p *__playerStateReader) SetTimerProp(name string, value *boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (p *playerState) Reader() boardgame.PropertyReader {
	return &__playerStateReader{p}
}

func (p *playerState) ReadSetter() boardgame.PropertyReadSetter {
	return &__playerStateReader{p}
}
