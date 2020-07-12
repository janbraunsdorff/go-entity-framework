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

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/janbraunsdorff/demo/pkg/database/access"
)

func main() {
	manager := access.NewDbManager(&access.DbConfig{
		UserName: "root",
		Password: "123",
		Protocol: "tcp",
		Address:  "localhost:3306",
		DbName:   "",
	})

	manager.DropDatabase()
	manager.CreateDatabase()
	manager.CreateAllTables()

}
