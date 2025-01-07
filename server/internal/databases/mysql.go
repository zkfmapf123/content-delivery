package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysqlParams struct {
	host     string
	port     string
	user     string
	password string
	database string
}

type MysqlConn struct {
	db *sqlx.DB
}

func CreateDBConnection() mysqlParams {
	return mysqlParams{}
}

func (p mysqlParams) WithHost(host string) mysqlParams {
	p.host = host
	return p
}

func (p mysqlParams) WithPort(port string) mysqlParams {
	p.port = port
	return p
}

func (p mysqlParams) WithUser(username string) mysqlParams {
	p.user = username
	return p
}

func (p mysqlParams) WithPassword(password string) mysqlParams {
	p.password = password
	return p
}

func (p mysqlParams) WithDatabase(database string) mysqlParams {
	p.database = database
	return p
}

func (p mysqlParams) Build() (MysqlConn, error) {
	params := fmt.Sprintf("%s:%s@(%s:%s)/%s", p.user, p.password, p.host, p.port, p.database)
	fmt.Println("db params : ", params)
	db, err := sqlx.Connect("mysql", params)

	// DB Connection Error
	if err == nil {
		return MysqlConn{}, err
	}

	// DB Ping Error
	err = db.Ping()
	return MysqlConn{
		db: db,
	}, err	
}


// func (m MysqlConn) Select(query string, args map[string]interface{}) (interface{}, error) {
// 	rows, err := m.db.NamedQuery(query, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }

// func (m MysqlConn) Insert(query string, args map[string]interface{}) (interface{}, error) {
// 	result, err := m.db.NamedExec(query, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (m MysqlConn) Update(query string, args map[string]interface{}) (interface{}, error) {
// 	result, err := m.db.NamedExec(query, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (m MysqlConn) Delete(query string, args map[string]interface{}) (interface{}, error) {
// 	result, err := m.db.NamedExec(query, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
