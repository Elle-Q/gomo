package config

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

//设置logger
func (e Logger) Setup()  {
	//log.Logger.s
}