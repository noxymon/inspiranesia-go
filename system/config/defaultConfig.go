package config

type DefaultConfig struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	Port int
	RootName string `mapstructure:"root-name"`
	HttpHandler string `mapstructure:"http-handler"`
}



