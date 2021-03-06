package game

import (
	"fmt"
)

type Action struct {
	Type ActionType
	Card *Card
}

type ActionType int

const (
	Pass ActionType = iota
	Play
	Attack
)

// For debugging and logging. Don't use this in the critical path.
func (a *Action) String() string {
	switch a.Type {
	case Pass:
		return "pass"
	case Play:
		return fmt.Sprintf("play %s", a.Card)
	case Attack:
		return fmt.Sprintf("attack with %s", a.Card)
	}
	panic("control should not reach here")
}
