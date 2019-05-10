package upgrade_demo

// Upgrade is interface that case can use
type Upgrade interface {
	GetClusterInfo() (*Config, error)
	DoKill(ip string, service string) error
	State(ip string, service string) (*Status, error)
	SetInjection(ip string, injectMethod string, args ...string) error
	RecoverInjection(ip string, injectMethod string, args ...string) error
	// Sync or async?
	UpgradeBinary(ip string, version string, name string) error
	// need case self implment this two action
	PauseCase(ip string, name string) error
	ResumeCase(ip string, name string) error
	
	StopCase(ip string, name string) error
	StartCase(ip string, name string) error
}