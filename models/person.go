package models

import (
	db "gin_exercise/database"
	"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	stmtIn, err := db.SqlDB.Prepare("INSERT INTO person(firstname, lastname) VALUES(?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmtIn.Exec(p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	stmtOut, err := db.SqlDB.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	res, err := stmtOut.Exec(p.Id)
	if err != nil {
		log.Println(err)
	}

	ra, _ = res.RowsAffected()

	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmtInt, err := db.SqlDB.Prepare("UPDATE person SET firstname=?, lastname=? WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	res, err := stmtInt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		log.Println(err)
	}

	ra, _ = res.RowsAffected()

	return
}

func (p *Person) GetPerson() (person Person, err error) {
	stmtOut, err := db.SqlDB.Prepare("SELECT id, firstname, lastname FROM person WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	err = stmtOut.QueryRow(p.Id).Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		log.Println(err)
	}

	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)

	stmsOut, err := db.SqlDB.Prepare("SELECT id, firstname, lastname FROM person")
	if err != nil {
		log.Println(err)
	}

	rows, err := stmsOut.Query()
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return
}
