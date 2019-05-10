package upgrade_demo

import (
	"time"
)
// TestState is test state
type TestState string

// TestState
const (
	TestStatePending   TestState = "PENDING"
	TestStateCreating            = "CREATING"
	TestStateRunning             = "RUNNING"
	TestStateFailed              = "FAILED"
	TestStateStopped             = "STOPPED"
	TestStateFinished            = "FINISHED"
	TestStateDestroyed           = "DESTORYED"
)

// Config defines the config of the client
type Config struct {
	Self       *TestInfo       `json:"self"`
	Cat        *CatInfo        `json:"cat"`
	Tests      []*TestInfo     `json:"tests"`
}

// CatInfo defines the information of cat
type CatInfo struct {
	ID             int64
	LBService      string                  `json:"lb_service"`
	PdService      string                  `json:"pd_service"`
	PDs            []*PodInfo              `json:"pds"`
	TiKVs          []*PodInfo              `json:"tikvs"`
	TiDBs          []*PodInfo              `json:"tidbs"`
}

// TestInfo defines the infomation of a test
type TestInfo struct {
	Name   string          `json:"name"`
	ID     int64           `json:"id"`
	IP     string          `json:"ip"`
	State  TestState `json:"state"`
}

// PodInfo defines the information of pod
type PodInfo struct {
	Name        string `json:"name"`
	IP          string `json:"ip"`
	AgentPort   int    `json:"agent_port"`
	NodeIP      string `json:"node_ip"`
}

// Status is the status of binary
type Status struct {
	StartTime 	time.Time
	StateTime 	time.Time
}

// IsTimeout determines if task is timeout
func (s *Status) IsTimeout() bool {
	return true
}