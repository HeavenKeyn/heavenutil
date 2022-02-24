package confutil

type DBProper interface {
	GetType() DBType
	GetHost() string
	GetPort() int
	GetUsername() string
	GetPassword() string
	GetDatabase() string
	GetNetwork() string
	GetDSN() string
}
