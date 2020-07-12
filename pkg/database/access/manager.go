/*
/$$$$$$$   /$$$$$$        /$$   /$$  /$$$$$$  /$$$$$$$$       /$$$$$$$$ /$$$$$$$  /$$$$$$ /$$$$$$$$
| $$__  $$ /$$__  $$      | $$$ | $$ /$$__  $$|__  $$__/      | $$_____/| $$__  $$|_  $$_/|__  $$__/
| $$  \ $$| $$  \ $$      | $$$$| $$| $$  \ $$   | $$         | $$      | $$  \ $$  | $$     | $$
| $$  | $$| $$  | $$      | $$ $$ $$| $$  | $$   | $$         | $$$$$   | $$  | $$  | $$     | $$
| $$  | $$| $$  | $$      | $$  $$$$| $$  | $$   | $$         | $$__/   | $$  | $$  | $$     | $$
| $$  | $$| $$  | $$      | $$\  $$$| $$  | $$   | $$         | $$      | $$  | $$  | $$     | $$
| $$$$$$$/|  $$$$$$/      | $$ \  $$|  $$$$$$/   | $$         | $$$$$$$$| $$$$$$$/ /$$$$$$   | $$
|_______/  \______/       |__/  \__/ \______/    |__/         |________/|_______/ |______/   |__/
*/
// Generated at: 2020-07-12 21:13:28.932896 +0200 CEST m=+0.001787744

package access

import (
	"database/sql"
	"log"
)

type DbManager struct {
	db *sql.DB
}

func NewDbManager(config *DbConfig) DbManager {
	manager := DbManager{}
	manager.openConnection(config)
	return manager
}

func (manager *DbManager) openConnection(config *DbConfig) {
	db, err := sql.Open("mysql", config.getDatasource())
	if err != nil {
		log.Fatal("could not connect to database", err)
	}
	manager.db = db
}

func (manager *DbManager) performDDL(cmd string) {
	tx, err := manager.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Exec(cmd)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (manager *DbManager) DropDatabase() {
	manager.performDDL("DROP DATABASE IF EXISTS Shop;")
}

func (manager *DbManager) CreateDatabase() {
	manager.performDDL("CREATE DATABASE IF NOT EXISTS Shop DEFAULT CHARACTER SET UTF8MB4;")
	manager.performDDL("USE Shop")
}

func (manager *DbManager) CreateAllTables() {
	manager.CreateTablePerson()
	manager.CreateTableOrders()

	manager.AddForeignKeyPerson()
	manager.AddForeignKeyOrders()

}

func (manager *DbManager) CreateTablePerson() {
	manager.performDDL(`
        CREATE TABLE Person (
             PersonId int PRIMARY KEY,
		     LastName varchar(60) NOT NULL,
		     FirstName varchar(60) NOT NULL,
		     Age int CHECK (Age >= 18)
		    
        );
    `)
}

func (manager *DbManager) AddForeignKeyPerson() {

}

func (manager *DbManager) CreateTableOrders() {
	manager.performDDL(`
        CREATE TABLE Orders (
             OrderId int PRIMARY KEY,
		     OrderNumber varchar(10) NOT NULL,
		     PersonId int NOT NULL
		    
        );
    `)
}

func (manager *DbManager) AddForeignKeyOrders() {
	manager.performDDL(`ALTER TABLE Orders ADD FOREIGN KEY (PersonId) REFERENCES Person(PersonId)`)

}
