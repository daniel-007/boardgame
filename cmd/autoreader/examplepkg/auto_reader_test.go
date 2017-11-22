/************************************
 *
 * This file contains auto-generated methods to help certain structs
 * implement boardgame.SubState and boardgame.MutableSubState. It was
 * generated by autoreader.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/

package examplepkg

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/enum"
)

// Implementation for testStruct

var __testStructReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"A": boardgame.TypeInt,
	"B": boardgame.TypeString,
}

type __testStructReader struct {
	data *testStruct
}

func (t *__testStructReader) Props() map[string]boardgame.PropertyType {
	return __testStructReaderProps
}

func (t *__testStructReader) Prop(name string) (interface{}, error) {
	props := t.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return t.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return t.BoolSliceProp(name)
	case boardgame.TypeEnum:
		return t.EnumProp(name)
	case boardgame.TypeInt:
		return t.IntProp(name)
	case boardgame.TypeIntSlice:
		return t.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return t.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return t.PlayerIndexSliceProp(name)
	case boardgame.TypeStack:
		return t.StackProp(name)
	case boardgame.TypeString:
		return t.StringProp(name)
	case boardgame.TypeStringSlice:
		return t.StringSliceProp(name)
	case boardgame.TypeTimer:
		return t.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (t *__testStructReader) SetProp(name string, value interface{}) error {
	props := t.Props()
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
		return t.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return t.SetBoolSliceProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return t.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return t.SetIntSliceProp(name, val)
	case boardgame.TypeEnum:
		return errors.New("SetProp does not allow setting mutable types. Use ConfigureProp instead.")
	case boardgame.TypeStack:
		return errors.New("SetProp does not allow setting mutable types. Use ConfigureProp instead.")
	case boardgame.TypeTimer:
		return errors.New("SetProp does not allow setting mutable types. Use ConfigureProp instead.")
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return t.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return t.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return t.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return t.SetStringSliceProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (t *__testStructReader) ConfigureProp(name string, value interface{}) error {
	props := t.Props()
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
		return t.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return t.SetBoolSliceProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return t.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return t.SetIntSliceProp(name, val)
	case boardgame.TypeEnum:
		val, ok := value.(enum.MutableVal)
		if !ok {
			return errors.New("Provided value was not of type enum.MutableVal")
		}
		return t.ConfigureMutableEnumProp(name, val)
	case boardgame.TypeStack:
		val, ok := value.(boardgame.MutableStack)
		if !ok {
			return errors.New("Provided value was not of type boardgame.MutableStack")
		}
		return t.ConfigureMutableStackProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(boardgame.MutableTimer)
		if !ok {
			return errors.New("Provided value was not of type boardgame.MutableTimer")
		}
		return t.ConfigureMutableTimerProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return t.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return t.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return t.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return t.SetStringSliceProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (t *__testStructReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (t *__testStructReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (t *__testStructReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (t *__testStructReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (t *__testStructReader) EnumProp(name string) (enum.Val, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (t *__testStructReader) ConfigureMutableEnumProp(name string, value enum.MutableVal) error {

	return errors.New("No such MutableEnum prop: " + name)

}

func (t *__testStructReader) MutableEnumProp(name string) (enum.MutableVal, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (t *__testStructReader) IntProp(name string) (int, error) {

	switch name {
	case "A":
		return t.data.A, nil

	}

	return 0, errors.New("No such Int prop: " + name)

}

func (t *__testStructReader) SetIntProp(name string, value int) error {

	switch name {
	case "A":
		t.data.A = value
		return nil

	}

	return errors.New("No such Int prop: " + name)

}

func (t *__testStructReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (t *__testStructReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (t *__testStructReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (t *__testStructReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndex prop: " + name)

}

func (t *__testStructReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (t *__testStructReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (t *__testStructReader) StackProp(name string) (boardgame.Stack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (t *__testStructReader) ConfigureMutableStackProp(name string, value boardgame.MutableStack) error {

	return errors.New("No such MutableStack prop: " + name)

}

func (t *__testStructReader) MutableStackProp(name string) (boardgame.MutableStack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (t *__testStructReader) StringProp(name string) (string, error) {

	switch name {
	case "B":
		return t.data.B, nil

	}

	return "", errors.New("No such String prop: " + name)

}

func (t *__testStructReader) SetStringProp(name string, value string) error {

	switch name {
	case "B":
		t.data.B = value
		return nil

	}

	return errors.New("No such String prop: " + name)

}

func (t *__testStructReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (t *__testStructReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (t *__testStructReader) TimerProp(name string) (boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (t *__testStructReader) ConfigureMutableTimerProp(name string, value boardgame.MutableTimer) error {

	return errors.New("No such MutableTimer prop: " + name)

}

func (t *__testStructReader) MutableTimerProp(name string) (boardgame.MutableTimer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (t *testStruct) Reader() boardgame.PropertyReader {
	return &__testStructReader{t}
}

func (t *testStruct) ReadSetter() boardgame.PropertyReadSetter {
	return &__testStructReader{t}
}

func (t *testStruct) ReadSetConfigurer() boardgame.PropertyReadSetConfigurer {
	return &__testStructReader{t}
}