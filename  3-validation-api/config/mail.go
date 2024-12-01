package config

type Mail struct {
	Email    string `env:"EMAIL"`
	Password string `env:"PASSWORD"`
	Address  string `env:"ADDRESS"`
	Host     string `env:"HOST"`
}
