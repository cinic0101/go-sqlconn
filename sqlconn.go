package sqlconn

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Databases struct {}

type DBConn struct {
	Driver string
	DataSource string
}

var once sync.Once
var dbInstances map[string]*DBConn

func (d *Databases) NewInstance(db string) *DBConn {
	once.Do(func() {
		c, err := UnmarshalConfig(os.Getenv("CONFIG"))
		if err != nil {
			panic(err)
		}

		dbInstances =  make(map[string]*DBConn, len(c.Databases))
		for k, v := range c.Databases {
			dbInstances[strings.ToLower(k)] = &DBConn{
				Driver: v.Driver,
				DataSource: fmtDataSource(v),
			}
		}
	})

	return dbInstances[strings.ToLower(db)]
}

func (d *DBConn) Query(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := sql.Open(d.Driver, d.DataSource)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db.Query(query, args...)
}

func fmtDataSource(db Database) string {
	switch db.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s",
			db.User, db.Password, db.Host, db.Database)
	case "sqlserver":
		return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			db.Host, db.User, db.Password, 1433, db.Database)
	default:
		return fmt.Sprintf("Unknown driver: %s", db.Driver)
	}
}