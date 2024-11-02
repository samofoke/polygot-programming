package progfile

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

}

func getConfig(opts Opts) (string, error) {

}

func getOperations(opts Opts) Operation {

}

func getArgs(opts Opts) ([]string, error) {

}

func NewConfig() (*Config, error) {

}
