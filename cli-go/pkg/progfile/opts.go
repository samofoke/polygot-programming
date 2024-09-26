package progfile

import "github.com/hellflame/argparse"

type Opts struct {
	Args   []string
	Config string
	Pwd    string
}

func GetOps() (*Opts, error) {
	parser := argparse.NewParser("progfile", "get all values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Default:    "",
		Required:   false,
	})

	config := parser.String("p", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	pwd := parser.String("c", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Opts{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}
