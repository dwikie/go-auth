package config

type config struct {
	App struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"app"`
}
