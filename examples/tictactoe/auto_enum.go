/************************************
 *
 * This file contains auto-generated methods to help configure enums.
 * It was generated by autoreader.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/

package tictactoe

import (
	"github.com/jkomoros/boardgame/enum"
)

var Enums = enum.NewSet()

var PhaseEnum = Enums.MustAdd("Phase", map[int]string{
	PhaseAfterFirstMove:  "After First Move",
	PhaseBeforeFirstMove: "Before First Move",
})
