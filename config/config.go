package config

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	cfg := Config{}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = ":8001"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "davlatbek"
	cfg.PostgresDatabase = "salesdb"
	cfg.PostgresPassword = "0313"
	cfg.PostgresPort = "5432"
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	return cfg
}
