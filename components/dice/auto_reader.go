/************************************
 *
 * This file contains auto-generated methods to help certain structs
 * implement boardgame.SubState and boardgame.MutableSubState. It was
 * generated by autoreader.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/

package dice

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/enum"
)

// Implementation for Value

var __ValueReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"Faces": boardgame.TypeIntSlice,
}

type __ValueReader struct {
	data *Value
}

func (v *__ValueReader) Props() map[string]boardgame.PropertyType {
	return __ValueReaderProps
}

func (v *__ValueReader) Prop(name string) (interface{}, error) {
	props := v.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return v.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return v.BoolSliceProp(name)
	case boardgame.TypeEnum:
		return v.EnumProp(name)
	case boardgame.TypeInt:
		return v.IntProp(name)
	case boardgame.TypeIntSlice:
		return v.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return v.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return v.PlayerIndexSliceProp(name)
	case boardgame.TypeStack:
		return v.StackProp(name)
	case boardgame.TypeString:
		return v.StringProp(name)
	case boardgame.TypeStringSlice:
		return v.StringSliceProp(name)
	case boardgame.TypeTimer:
		return v.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (v *__ValueReader) SetProp(name string, value interface{}) error {
	props := v.Props()
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
		return v.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return v.SetBoolSliceProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return v.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return v.SetIntSliceProp(name, val)
	case boardgame.TypeEnum:
		val, ok := value.(enum.MutableVal)
		if !ok {
			return errors.New("Provided value was not of type enum.MutableVal")
		}
		return v.SetMutableEnumProp(name, val)
	case boardgame.TypeStack:
		val, ok := value.(boardgame.MutableStack)
		if !ok {
			return errors.New("Provided value was not of type boardgame.MutableStack")
		}
		return v.SetMutableStackProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return v.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return v.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return v.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return v.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type boardgame.Timer")
		}
		return v.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (v *__ValueReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (v *__ValueReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (v *__ValueReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (v *__ValueReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (v *__ValueReader) EnumProp(name string) (enum.Val, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (v *__ValueReader) SetMutableEnumProp(name string, value enum.MutableVal) error {

	return errors.New("No such MutableEnum prop: " + name)

}

func (v *__ValueReader) MutableEnumProp(name string) (enum.MutableVal, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (v *__ValueReader) IntProp(name string) (int, error) {

	return 0, errors.New("No such Int prop: " + name)

}

func (v *__ValueReader) SetIntProp(name string, value int) error {

	return errors.New("No such Int prop: " + name)

}

func (v *__ValueReader) IntSliceProp(name string) ([]int, error) {

	switch name {
	case "Faces":
		return v.data.Faces, nil

	}

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (v *__ValueReader) SetIntSliceProp(name string, value []int) error {

	switch name {
	case "Faces":
		v.data.Faces = value
		return nil

	}

	return errors.New("No such IntSlice prop: " + name)

}

func (v *__ValueReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (v *__ValueReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndex prop: " + name)

}

func (v *__ValueReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (v *__ValueReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (v *__ValueReader) StackProp(name string) (boardgame.Stack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (v *__ValueReader) SetMutableStackProp(name string, value boardgame.MutableStack) error {

	return errors.New("No such MutableStack prop: " + name)

}

func (v *__ValueReader) MutableStackProp(name string) (boardgame.MutableStack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (v *__ValueReader) StringProp(name string) (string, error) {

	return "", errors.New("No such String prop: " + name)

}

func (v *__ValueReader) SetStringProp(name string, value string) error {

	return errors.New("No such String prop: " + name)

}

func (v *__ValueReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (v *__ValueReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (v *__ValueReader) TimerProp(name string) (boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (v *__ValueReader) SetTimerProp(name string, value boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (v *Value) Reader() boardgame.PropertyReader {
	return &__ValueReader{v}
}

func (v *Value) ReadSetter() boardgame.PropertyReadSetter {
	return &__ValueReader{v}
}

// Implementation for DynamicValue

var __DynamicValueReaderProps map[string]boardgame.PropertyType = map[string]boardgame.PropertyType{
	"SelectedFace": boardgame.TypeInt,
	"Value":        boardgame.TypeInt,
}

type __DynamicValueReader struct {
	data *DynamicValue
}

func (d *__DynamicValueReader) Props() map[string]boardgame.PropertyType {
	return __DynamicValueReaderProps
}

func (d *__DynamicValueReader) Prop(name string) (interface{}, error) {
	props := d.Props()
	propType, ok := props[name]

	if !ok {
		return nil, errors.New("No such property with that name: " + name)
	}

	switch propType {
	case boardgame.TypeBool:
		return d.BoolProp(name)
	case boardgame.TypeBoolSlice:
		return d.BoolSliceProp(name)
	case boardgame.TypeEnum:
		return d.EnumProp(name)
	case boardgame.TypeInt:
		return d.IntProp(name)
	case boardgame.TypeIntSlice:
		return d.IntSliceProp(name)
	case boardgame.TypePlayerIndex:
		return d.PlayerIndexProp(name)
	case boardgame.TypePlayerIndexSlice:
		return d.PlayerIndexSliceProp(name)
	case boardgame.TypeStack:
		return d.StackProp(name)
	case boardgame.TypeString:
		return d.StringProp(name)
	case boardgame.TypeStringSlice:
		return d.StringSliceProp(name)
	case boardgame.TypeTimer:
		return d.TimerProp(name)

	}

	return nil, errors.New("Unexpected property type: " + propType.String())
}

func (d *__DynamicValueReader) SetProp(name string, value interface{}) error {
	props := d.Props()
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
		return d.SetBoolProp(name, val)
	case boardgame.TypeBoolSlice:
		val, ok := value.([]bool)
		if !ok {
			return errors.New("Provided value was not of type []bool")
		}
		return d.SetBoolSliceProp(name, val)
	case boardgame.TypeInt:
		val, ok := value.(int)
		if !ok {
			return errors.New("Provided value was not of type int")
		}
		return d.SetIntProp(name, val)
	case boardgame.TypeIntSlice:
		val, ok := value.([]int)
		if !ok {
			return errors.New("Provided value was not of type []int")
		}
		return d.SetIntSliceProp(name, val)
	case boardgame.TypeEnum:
		val, ok := value.(enum.MutableVal)
		if !ok {
			return errors.New("Provided value was not of type enum.MutableVal")
		}
		return d.SetMutableEnumProp(name, val)
	case boardgame.TypeStack:
		val, ok := value.(boardgame.MutableStack)
		if !ok {
			return errors.New("Provided value was not of type boardgame.MutableStack")
		}
		return d.SetMutableStackProp(name, val)
	case boardgame.TypePlayerIndex:
		val, ok := value.(boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type boardgame.PlayerIndex")
		}
		return d.SetPlayerIndexProp(name, val)
	case boardgame.TypePlayerIndexSlice:
		val, ok := value.([]boardgame.PlayerIndex)
		if !ok {
			return errors.New("Provided value was not of type []boardgame.PlayerIndex")
		}
		return d.SetPlayerIndexSliceProp(name, val)
	case boardgame.TypeString:
		val, ok := value.(string)
		if !ok {
			return errors.New("Provided value was not of type string")
		}
		return d.SetStringProp(name, val)
	case boardgame.TypeStringSlice:
		val, ok := value.([]string)
		if !ok {
			return errors.New("Provided value was not of type []string")
		}
		return d.SetStringSliceProp(name, val)
	case boardgame.TypeTimer:
		val, ok := value.(boardgame.Timer)
		if !ok {
			return errors.New("Provided value was not of type boardgame.Timer")
		}
		return d.SetTimerProp(name, val)

	}

	return errors.New("Unexpected property type: " + propType.String())
}

func (d *__DynamicValueReader) BoolProp(name string) (bool, error) {

	return false, errors.New("No such Bool prop: " + name)

}

func (d *__DynamicValueReader) SetBoolProp(name string, value bool) error {

	return errors.New("No such Bool prop: " + name)

}

func (d *__DynamicValueReader) BoolSliceProp(name string) ([]bool, error) {

	return []bool{}, errors.New("No such BoolSlice prop: " + name)

}

func (d *__DynamicValueReader) SetBoolSliceProp(name string, value []bool) error {

	return errors.New("No such BoolSlice prop: " + name)

}

func (d *__DynamicValueReader) EnumProp(name string) (enum.Val, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (d *__DynamicValueReader) SetMutableEnumProp(name string, value enum.MutableVal) error {

	return errors.New("No such MutableEnum prop: " + name)

}

func (d *__DynamicValueReader) MutableEnumProp(name string) (enum.MutableVal, error) {

	return nil, errors.New("No such Enum prop: " + name)

}

func (d *__DynamicValueReader) IntProp(name string) (int, error) {

	switch name {
	case "Value":
		return d.data.Value, nil
	case "SelectedFace":
		return d.data.SelectedFace, nil

	}

	return 0, errors.New("No such Int prop: " + name)

}

func (d *__DynamicValueReader) SetIntProp(name string, value int) error {

	switch name {
	case "Value":
		d.data.Value = value
		return nil
	case "SelectedFace":
		d.data.SelectedFace = value
		return nil

	}

	return errors.New("No such Int prop: " + name)

}

func (d *__DynamicValueReader) IntSliceProp(name string) ([]int, error) {

	return []int{}, errors.New("No such IntSlice prop: " + name)

}

func (d *__DynamicValueReader) SetIntSliceProp(name string, value []int) error {

	return errors.New("No such IntSlice prop: " + name)

}

func (d *__DynamicValueReader) PlayerIndexProp(name string) (boardgame.PlayerIndex, error) {

	return 0, errors.New("No such PlayerIndex prop: " + name)

}

func (d *__DynamicValueReader) SetPlayerIndexProp(name string, value boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndex prop: " + name)

}

func (d *__DynamicValueReader) PlayerIndexSliceProp(name string) ([]boardgame.PlayerIndex, error) {

	return []boardgame.PlayerIndex{}, errors.New("No such PlayerIndexSlice prop: " + name)

}

func (d *__DynamicValueReader) SetPlayerIndexSliceProp(name string, value []boardgame.PlayerIndex) error {

	return errors.New("No such PlayerIndexSlice prop: " + name)

}

func (d *__DynamicValueReader) StackProp(name string) (boardgame.Stack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (d *__DynamicValueReader) SetMutableStackProp(name string, value boardgame.MutableStack) error {

	return errors.New("No such MutableStack prop: " + name)

}

func (d *__DynamicValueReader) MutableStackProp(name string) (boardgame.MutableStack, error) {

	return nil, errors.New("No such Stack prop: " + name)

}

func (d *__DynamicValueReader) StringProp(name string) (string, error) {

	return "", errors.New("No such String prop: " + name)

}

func (d *__DynamicValueReader) SetStringProp(name string, value string) error {

	return errors.New("No such String prop: " + name)

}

func (d *__DynamicValueReader) StringSliceProp(name string) ([]string, error) {

	return []string{}, errors.New("No such StringSlice prop: " + name)

}

func (d *__DynamicValueReader) SetStringSliceProp(name string, value []string) error {

	return errors.New("No such StringSlice prop: " + name)

}

func (d *__DynamicValueReader) TimerProp(name string) (boardgame.Timer, error) {

	return nil, errors.New("No such Timer prop: " + name)

}

func (d *__DynamicValueReader) SetTimerProp(name string, value boardgame.Timer) error {

	return errors.New("No such Timer prop: " + name)

}

func (d *DynamicValue) Reader() boardgame.PropertyReader {
	return &__DynamicValueReader{d}
}

func (d *DynamicValue) ReadSetter() boardgame.PropertyReadSetter {
	return &__DynamicValueReader{d}
}
