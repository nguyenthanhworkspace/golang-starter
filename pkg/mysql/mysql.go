package mysql

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Mysql -.
type Mysql struct {
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	DB      *sql.DB
}

func New(url string) (*Mysql, error) {
	mysql := &Mysql{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	mysql.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	for mysql.connAttempts > 0 {
		DB, err := sql.Open("mysql", url)

		if err == nil {
			mysql.DB = DB
			break
		}

		log.Printf("MySQL error: %d", err)
		log.Printf("MySQL is trying to connect, attempts left: %d", mysql.connAttempts)

		time.Sleep(mysql.connTimeout)

		mysql.connAttempts--
	}

	return mysql, nil
}

// Close -.
func (p *Mysql) Close() {
	if p.DB != nil {
		err := p.DB.Close()

		if err != nil {
			return
		}
	}
}
