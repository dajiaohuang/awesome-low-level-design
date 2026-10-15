package loggingframework

import "testing"

func TestNewDatabaseAppenderRegistersMySQLDriver(t *testing.T) {
	appender, err := NewDatabaseAppender("tcp(127.0.0.1:3306)/logs", "user", "password")
	if err != nil {
		t.Fatalf("NewDatabaseAppender: %v", err)
	}
	if err := appender.db.Close(); err != nil {
		t.Fatalf("close database handle: %v", err)
	}
}
