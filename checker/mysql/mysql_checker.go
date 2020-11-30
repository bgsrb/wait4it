package mysql

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"wait4it/config"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//MySQLChecker ...
type MySQLChecker struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

//BuildContext...
func (ch *MySQLChecker) BuildContext(cx config.CheckContext) {
	ch.Port = cx.Port
	ch.Host = cx.Host
	ch.Username = cx.Username
	ch.Password = cx.Password
	ch.DatabaseName = cx.DatabaseName
}

//Validate ...
func (ch *MySQLChecker) Validate() error {
	if len(ch.Host) == 0 || len(ch.Username) == 0 {
		return errors.New("host or username can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("invalid port range for mysql")
	}

	return nil
}

//Check ...
func (ch *MySQLChecker) Check() (bool, bool, error) {

	dsl := ch.buildConnectionString()

	db, err := sql.Open("mysql", dsl)

	// if there is an error opening the connection, handle it
	if err != nil {
		return false, true, err
	}

	err = mysql.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
	if err != nil {
		return false, true, err
	}

	err = db.Ping()
	if err != nil {
		// todo: need a logger
		return false, false, nil
	}
	_ = db.Close()

	return true, true, nil
}

func (ch MySQLChecker) buildConnectionString() string {
	dsl := ""

	if len(ch.Password) == 0 {
		dsl = dsl + ch.Username
	} else {
		dsl = dsl + ch.Username + ":" + ch.Password
	}

	dsl = dsl + "@tcp(" + ch.Host + ":" + strconv.Itoa(ch.Port) + ")/"

	if len(ch.DatabaseName) > 0 {
		dsl = dsl + ch.DatabaseName
	}

	return dsl
}
