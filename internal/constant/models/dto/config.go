package dto

type Config struct {
	Db        DbConfig  `mapstructure:"database"`
	Server    Server    `mapstructure:"server"`
	Migration Migration `mapstucture:"migration"`
	Redis     Server    `mapstucture:"redis"`
	Authz     Authz     `mapstucture:"authz"`
}

type DbConfig struct {
	Url   string `mapstructure:"url"`
	Name  string `mapstructure:"db"`
	PgUrl string `mapstructure:"pgurl"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Migration struct {
	Path string `mapstructure:"path"`
}

type Authz struct {
	Model   string `mapstructure:"model"`
	Adapter string `mapstructure:"adapter"`
}
