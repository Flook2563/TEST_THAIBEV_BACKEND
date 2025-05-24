package appconfig

type AppConfig struct {
	Database Database
	Server   struct {
		Port string
	}
	Log struct {
		Level string
	}
}

type Database struct {
	Host       string
	DBName     string
	Port       string
	SSLMode    string
	User       string
	Password   string
	Timezone   string
	SearchPath string
}
