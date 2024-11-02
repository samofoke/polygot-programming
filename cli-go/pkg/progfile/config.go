package progfile

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
)

type Config struct {
	Args      []string
	Operation Operation
	Pwd       string
	Config    string
}

func getPwd(opts *Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getConfig(opts *Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}
	customConfig, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(customConfig, "projector", "projector.json"), nil
}

func getOperations(opts *Opts) Operation {
	if len(opts.Args) == 0 {
		return Print
	}
	if opts.Args[0] == "add" {
		return Add
	}
	return Print
}

func getArgs(opts *Opts) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}
	operation := getOperations(opts)
	if operation == Add {
		if len(opts.Args) != 3 {
			return nil, fmt.Errorf("requires 2 args, but recieved %v", len(opts.Args)-1)
		}
		return opts.Args[1:], nil
	}
	if len(opts.Args) > 1 {
		return nil, fmt.Errorf("print requires 0 or 1 arguments %v", len(opts.Args))
	}
	return opts.Args, nil
}

func NewConfig(opts *Opts) (*Config, error) {
	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}

	custConfig, err := getConfig(opts)
	if err != nil {
		return nil, err
	}

	custArg, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Pwd:       pwd,
		Config:    custConfig,
		Args:      custArg,
		Operation: getOperations(opts),
	}, nil

}
