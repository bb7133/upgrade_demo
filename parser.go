package upgrade_demo

// Parser is parser interface
type Parser interface {
	// should be implement
	ParserStatement(file string) ([]*Statement, error)
}

// Statement is statement
type Statement struct {
	version string
	action  Action
}

// Action is action
type Action struct {
	name string
	args []string
}

// getIPFromCfg using name in files to gets IP from cat config
func getIPFromCfg(name string, cfg *Config) (string, error) {
	return "1.1.1.1", nil
}
