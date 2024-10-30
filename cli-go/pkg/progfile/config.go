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
