/************************************
 *
 * This file contains auto-generated methods to help configure enums.
 * It was generated by autoreader.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/

package examplepkg

import (
	"github.com/jkomoros/boardgame/enum"
)

var Enums = enum.NewSet()

var ColorEnum = Enums.MustAdd("Color", map[int]string{
	ColorUnknown: "Unknown",
	ColorBlue:    "Blue",
	ColorGreen:   "Green",
	ColorRed:     "Red",
})