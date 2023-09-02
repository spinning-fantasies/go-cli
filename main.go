package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dixonwille/wmenu/v5"
)

func main() {

	// Connect to database
	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)
	// defer close
	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}

}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {

	case 0:

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a first name: ")
		firstName, _ := reader.ReadString('\n')
		fmt.Print("Enter a last name: ")
		lastName, _ := reader.ReadString('\n')
		fmt.Print("Enter an email address: ")
		email, _ := reader.ReadString('\n')
		fmt.Print("Enter an IP address: ")
		ipAddress, _ := reader.ReadString('\n')

		newPerson := person{
			first_name: firstName,
			last_name:  lastName,
			email:      email,
			ip_address: ipAddress,
		}

		addPerson(db, newPerson)

		break
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
