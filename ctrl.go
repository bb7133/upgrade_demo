package upgrade_demo

import (
	"github.com/juju/errors"
)
// Ctrl is the controller 
type Ctrl struct{
	cfg *Config
	st []*Statement
	upgrade Upgrade
	parser Parser
	statementFile string
}

// Starts start controller 
func (c *Ctrl) Start() error {
	c.cfg, _ = c.upgrade.GetClusterInfo()
	c.st,_ = c.parser.ParserStatement(c.statementFile)
	return nil
}

// Run runs all statement
func (c *Ctrl) Run() error {
	for _, st := range c.st {
		err := c.RunStatement(st)
		if err != nil {
			return err
		}
	}
	return nil
}

// RunStatement run  statement
func (c *Ctrl) RunStatement(st *Statement) error {
	switch st.action.name {
	case "kill":
		// args[0] can be pod name of tidb cluster
		ip,_ := getIPFromCfg(st.action.args[0], c.cfg)
		err := c.upgrade.DoKill(ip, "tidb")
		if err != nil {
			return err
		}
	case "upgrade":
		ip,_ := getIPFromCfg(st.action.args[0], c.cfg)
		// args[1] can be version
		version := st.action.args[1]
		err := c.upgrade.UpgradeBinary(ip, version, "tidb")
		if err != nil {
			return err
		}

		state, err := c.upgrade.State(ip, "tidb")
		if err != nil {
			return err
		}

		if state.IsTimeout() {
			return errors.New("timeout")
		}
	}
	return nil
}