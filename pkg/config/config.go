package config

type Config struct {
	Postgres Postgres `yaml:"postgres" envPrefix:"POSTGRES_"`
}

type Postgres struct {
	Host     string `yaml:"host" env:"HOST" required:"true" default:"localhost"`
	Port     int    `yaml:"port" env:"PORT" required:"true" default:"5432"`
	User     string `yaml:"user" env:"USER" required:"true"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname" env:"DBNAME" required:"true"`
	SSLMode  string `yaml:"sslmode"`
}
