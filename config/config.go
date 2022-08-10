package config

type AppConfig struct {
	Name string
	Port string
}

type DBConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Schema string
}
