package config

// ENV ...
type ENV struct {
	Env     string `env:"ENV"`
	MongoDB struct {
		Host      string `env:"MONGO_HOST,required"`
		DBName    string `env:"MONGO_DBName,required"`
		Mechanism string `env:"MONGO_MECHANISM"`
		Source    string `env:"MONGO_SOURCE"`
		User      string `env:"MONGO_USER"`
		Password  string `env:"MONGO_PWD"`
	}

	Port       string `env:"ADMIN_PORT,required"`
	AuthSecret string `env:"AUTH_SECRET,required"`

	Email struct {
		FromName       string `env:"EMAIL_FROM_NAME,required"`
		FromAddress    string `env:"EMAIL_FROM_ADDRESS,required"`
		SendGridApiKey string `env:"EMAIL_SENDGRID_API_KEY,required"`
	}
}
