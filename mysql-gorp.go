package main

/*
 * USAGE: 
 *     go run mysql-gorp.go
 * 
 * REQUIREMENT: 
 *     go get github.com/go-gorp/gorp
 *     go get github.com/go-sql-driver/mysql
 * 
 * :: SQL ::
 * CREATE DATABASE test
 * CREATE TABLE `todo` (
 *   `id` bigint(20) NOT NULL AUTO_INCREMENT,
 *   `description` varchar(255) DEFAULT NULL,
 *   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 *   PRIMARY KEY (`id`)
 * ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8
 */

import (
	"database/sql"
	"fmt"
	"time"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id int              `db:id`
	Description string  `db:description`
	CreatedAt time.Time `db:created_at`
}


func main() {
	db, err := sql.Open("mysql", "root:admin@/test")
	// also "root:admin@tcp(localhost:3306)/test"
	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		os.Exit(1)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	rows, _ := dbMap.Select(Todo{}, "select * from todo")
	for _, row := range rows {
		todo := row.(*Todo)
		fmt.Println(todo.Description)
	}
	db.Close()
}
