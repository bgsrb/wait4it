package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"wait4it/config"

	_ "github.com/lib/pq"
)

//PostgreSQLChecker ...
type PostgreSQLChecker struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	SSLMode      string
}

//BuildContext...
func (ch *PostgreSQLChecker) BuildContext(cx config.CheckContext) {
	ch.Port = cx.Port
	ch.Host = cx.Host
	ch.Username = cx.Username
	ch.Password = cx.Password
	ch.DatabaseName = cx.DatabaseName
	if len(cx.DBConf.SSLMode) < 1 {
		ch.SSLMode = "disable"
	} else {
		ch.SSLMode = cx.DBConf.SSLMode
	}
}

//Validate...
func (ch *PostgreSQLChecker) Validate() error {
	if len(ch.Host) == 0 || len(ch.Username) == 0 {
		return errors.New("host or username can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("invalid port range for PostgresSQL")
	}

	return nil
}

//Check ...
func (ch *PostgreSQLChecker) Check() (bool, bool, error) {
	dsl := ch.buildConnectionString()

	db, err := sql.Open("postgres", dsl)

	// if there is an error opening the connection, handle it
	if err != nil {
		return false, true, err
	}

	err = db.Ping()
	if err != nil {
		return false, false, nil
	}
	_ = db.Close()

	return true, true, nil
}

func (ch PostgreSQLChecker) buildConnectionString() string {
	dsl := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s dbname=%s ",
		ch.Host, ch.Port, ch.Username, ch.Password, ch.SSLMode, ch.DatabaseName)

	return dsl
}
