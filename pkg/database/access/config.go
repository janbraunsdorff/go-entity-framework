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
// Generated at: 2020-07-12 21:13:28.932531 +0200 CEST m=+0.001422262
package access

import (
	"fmt"
)

type DbConfig struct {
	UserName string
	Password string
	Protocol string
	Address  string
	DbName   string
}

func (conf *DbConfig) getDatasource() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", conf.UserName, conf.Password, conf.Protocol, conf.Address, conf.DbName)
}
