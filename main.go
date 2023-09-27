package main

// TODO: iac Lambda Store by plumi

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	getCmd := NewGetCommand()
	setCmd := NewSetCommand()
	cmds := []Command{getCmd, setCmd}

	if err := ExecSubCommand(cmds); err != nil {
		log.Fatal(err)
	}
}

func ExecSubCommand(cmds []Command) error {
	subCommands := os.Args[1:]
	for _, c := range cmds {
		if c.Name() == subCommands[0] {
			if err := c.Parse(); err != nil {
				return fmt.Errorf("failed flag parse: %v", err)
			}
			if err := c.RunCommand(); err != nil {
				return fmt.Errorf("failed execute command: %v", err)
			}
			break
		}
	}
	return nil
}

type Command interface {
	Name() string
	RunCommand() error
	Parse() error
}

type GetCommand struct {
	FlagSet *flag.FlagSet
	Key     *string
}

func NewGetCommand() *GetCommand {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	key := cmd.String("key", "", "get key")
	return &GetCommand{
		FlagSet: cmd,
		Key:     key,
	}
}

func (g *GetCommand) Name() string {
	return g.FlagSet.Name()
}
func (g *GetCommand) RunCommand() error {
	fmt.Println("run get")
	fmt.Printf("set key: %s", *g.Key)
	// TODO: kvs
	return nil
}
func (g *GetCommand) Parse() error {
	return g.FlagSet.Parse(os.Args[2:])
}

type SetCommand struct {
	FlagSet *flag.FlagSet
	Key     *string
	Value   *string
}

func NewSetCommand() *SetCommand {
	cmd := flag.NewFlagSet("set", flag.ExitOnError)
	key := cmd.String("key", "", "set key")
	value := cmd.String("value", "", "set value")
	return &SetCommand{
		FlagSet: cmd,
		Key:     key,
		Value:   value,
	}
}

func (s *SetCommand) Name() string {
	return s.FlagSet.Name()
}
func (s *SetCommand) RunCommand() error {
	fmt.Println("executed set")
	fmt.Printf("set key: %s, value: %s", *s.Key, *s.Value)
	// TODO: kvs
	return nil
}
func (s *SetCommand) Parse() error {
	return s.FlagSet.Parse(os.Args[2:])
}
