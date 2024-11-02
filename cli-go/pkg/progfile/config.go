package progfile

import (
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

func getPwd(opts Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getConfig(opts Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}
	customConfig, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(customConfig, "projector", "projector.json"), nil
}

func getOperations(opts Opts) Operation {

}

func getArgs(opts Opts) ([]string, error) {

}

func NewConfig() (*Config, error) {

}
