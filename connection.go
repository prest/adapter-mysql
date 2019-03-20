package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/joelmdesouza/mysql/internal/connection"
)

// GetURI mysql connection URI
func GetURI(DBName string) string {
	return connection.GetURI(DBName)
}

// Get get mysql connection
func Get() (*sqlx.DB, error) {
	return connection.Get()
}

// GetPool of connection
func GetPool() *connection.ConnectionPool {
	return connection.GetPool()
}

// AddDatabaseToPool add connection to pool
func AddDatabaseToPool(name string, DB *sqlx.DB) {
	connection.AddDatabaseToPool(name, DB)
}

// MustGet get mysql connection
func MustGet() *sqlx.DB {
	return connection.MustGet()
}

// SetDatabase set current database in use
func SetDatabase(name string) {
	connection.SetDatabase(name)
}

// GetDatabase get current database in use
func GetDatabase() string {
	return connection.GetDatabase()
}
