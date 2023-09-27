package main

// TODO: iac upstash Store by plumi

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shoet/go-cli-with-kvs/config"
	"github.com/shoet/go-cli-with-kvs/store"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to new config: %v", err)
	}
	kvs, err := store.NewRedisKVS(cfg)
	if err != nil {
		log.Fatalf("failed to new redis kvs: %v", err)
	}

	getCmd := NewGetCommand(kvs)
	setCmd := NewSetCommand(kvs)
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
	KVS     KVS
	FlagSet *flag.FlagSet
	Key     *string
}

func NewGetCommand(kvs KVS) *GetCommand {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	key := cmd.String("key", "", "get key")
	return &GetCommand{
		FlagSet: cmd,
		Key:     key,
		KVS:     kvs,
	}
}

func (g *GetCommand) Name() string {
	return g.FlagSet.Name()
}
func (g *GetCommand) RunCommand() error {
	val, err := g.KVS.Get(context.Background(), *g.Key)
	if err != nil {
		return fmt.Errorf("failed to get value: %w", err)
	}
	fmt.Println(val)
	return nil
}
func (g *GetCommand) Parse() error {
	return g.FlagSet.Parse(os.Args[2:])
}

type SetCommand struct {
	FlagSet *flag.FlagSet
	Key     *string
	Value   *string
	KVS     KVS
}

func NewSetCommand(kvs KVS) *SetCommand {
	cmd := flag.NewFlagSet("set", flag.ExitOnError)
	key := cmd.String("key", "", "set key")
	value := cmd.String("value", "", "set value")
	return &SetCommand{
		FlagSet: cmd,
		Key:     key,
		Value:   value,
		KVS:     kvs,
	}
}

func (s *SetCommand) Name() string {
	return s.FlagSet.Name()
}
func (s *SetCommand) RunCommand() error {
	fmt.Println("executed set")
	fmt.Printf("set key: %s, value: %s", *s.Key, *s.Value)
	if err := s.KVS.Set(context.Background(), *s.Key, *s.Value); err != nil {
		return fmt.Errorf("failed to set value: %w", err)
	}
	return nil
}
func (s *SetCommand) Parse() error {
	return s.FlagSet.Parse(os.Args[2:])
}

type KVS interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
}
