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
// Generated at: 2020-07-12 21:13:28.933355 +0200 CEST m=+0.002246539

package entity

type Person struct {
	PersonId  int
	LastName  string
	FirstName string
	Age       int
	Orders    []Orders
}
