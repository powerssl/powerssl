package internal

type Config struct {
	ConfigDir string   `flag:"configDir;;config;Config directory to load a set of yaml config files from" validate:"required"`
	Env       string   `flag:"env;e;development;Environment is one of the input params ex-development" validate:"required"`
	Services  []string `flag:"services;;frontend,history,matching,worker;Service(s) to start" validate:"gt=0,required"`
	Zone      string   `flag:"zone;;;Zone is another input param"`
}

func (cfg *Config) Defaults() {
	return
}
