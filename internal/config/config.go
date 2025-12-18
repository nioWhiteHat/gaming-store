package config



type DBConfig struct {
	
	DSNForApp string

	DSNForMigrate string
}

func LoadDBConfig() DBConfig {
	
	baseDSN := "postgres:236691@localhost:5432/fake_keys?sslmode=disable"

	return DBConfig{
		DSNForApp:     "postgres://" + baseDSN,
		DSNForMigrate: "postgres://postgres:236691@localhost:5432/fake_keys?sslmode=disable",
	}
}
