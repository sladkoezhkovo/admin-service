package config

type Config struct {
	App AppConfig `yaml:"app"`
	Pg  SqlConfig `yaml:"pg"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

type SqlConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   string `yaml:"db"`
	SSL  string `yaml:"ssl" env-default:"disable"`
}
