package system

type window struct {
	Width     int32
	Height    int32
	Monitor   int32
	Mode      byte
	TargetFPS int32
}

type connection struct {
	Ip   string
	Port string
}

type settings struct {
	SoundVolume float32
	MusicVolume float32
}

type controls struct {
	Left  int32
	Right int32
	Up    int32
	Down  int32
	Menu  int32
}

// Configuration is the global object used to hold game settings
type Configuration struct {
	Window     window     `mapstructure:"window"`
	Connection connection `mapstructure:"connection"`
	Settings   settings   `mapstructure:"settings"`
	Controls   controls   `mapstructure:"controls"`
}
