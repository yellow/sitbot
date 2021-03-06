package msl

import (
	"fmt"
)

type Command interface {
	Emit() string
}

// used in stack
type commandUnion struct {
	IfCommand
	WhileCommand
	Block
	Statement

	Result Command
	Line   int
}

type IfCommand struct {
	ifCond    string
	ifCommand Command

	elseCond     []string
	elseCommands []Command
}

type WhileCommand struct {
	cond    string
	command Command
}

type Block struct {
	commands []Command
}

func (b *Block) Add(c Command) {
	if c == nil {
		panic("undefined command")
	}
	b.commands = append(b.commands, c)
}

func (b *Block) Emit() (ret string) {
	for _, c := range b.commands {
		if c == nil {
			ret += "???\n"
			continue
		}
		ret += c.Emit() + "\n"
	}
	return ret
}

type Statement struct {
	Values []string
}

func (m *Statement) Emit() string {
	return fmt.Sprintf("stmt(%q)", m.Values)
}
