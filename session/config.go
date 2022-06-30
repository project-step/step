package session

// MySQLConfig configures a MySQL compatible database connection.
type MySQLConfig struct {
	Host    string
	Port    int
	User    string
	Pass    string
	Db      string
	Options string
}

// SessionManagerConfig is the configuration for the step.
type SessionManagerConfig struct {
	User  string
	Pass  string
	Port  int
	MySQL MySQLConfig
}
