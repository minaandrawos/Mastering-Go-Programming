package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type crewMember struct {
	id           int
	name         string
	secClearance int
	position     string
}

type Crew []crewMember

func main() {
	db, err := sql.Open("sqlite3", "hydra.db")
	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS
						Personnel(id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, SecurityClearance INTEGER, Position TEXT) `)

	//CSVToMySQL(db)

	cw := GetCrewByPositions(db, []string{"'Mechanic'", "'Biologist'"})
	fmt.Println(cw)

	fmt.Println(GetCrewMemberById(db, 11))

	AddCrewMember(db, crewMember{name: "Kaya Gal", secClearance: 4, position: "Biologist"})

	fmt.Println(GetCrewMemberByPosition(db, "Chemist"))

	cr := Crew{
		crewMember{name: "Adam stler", secClearance: 4, position: "Chemist"},
		crewMember{name: "Zach Garph", secClearance: 5, position: "Biologist"},
	}
	CreateCrewMembersByTransaction(db, cr)

}

func GetCrewByPositions(db *sql.DB, positions []string) Crew {

	Qs := fmt.Sprintf("SELECT id,Name,SecurityClearance,Position from Personnel where Position in (%s);", strings.Join(positions, ","))

	rows, err := db.Query(Qs)
	if err != nil {
		log.Fatal("Could not get data from the Personnel table ", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	for rows.Next() {
		member := crewMember{}
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err != nil {
			log.Fatal("Error scanning row", err)
		}
		retVal = append(retVal, member)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return retVal
}

func GetCrewMemberById(db *sql.DB, id int) (cm crewMember) {
	row := db.QueryRow("Select * from Personnel where id = ?", id)

	err := row.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetCrewMemberByPosition(db *sql.DB, position string) (cr Crew) {

	stmt, err := db.Prepare("Select * from Personnel where Position = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(position)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		cm := crewMember{}
		err = rows.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)
		if err != nil {
			log.Fatal(err)
		}
		cr = append(cr, cm)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func AddCrewMember(db *sql.DB, cm crewMember) int64 {

	res, err := db.Exec("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES (?,?,?)", cm.name, cm.secClearance, cm.position)
	if err != nil {
		log.Fatal(err)
	}
	ra, _ := res.RowsAffected()
	id, _ := res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted id", id)
	return id
}

func CreateCrewMembersByTransaction(db *sql.DB, cr Crew) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Could not begin tx", err)
	}
	stmt, err := tx.Prepare("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES (?,?,?);")
	if err != nil {
		tx.Rollback()
		log.Fatal("Could not do select statement ", err)
	}
	defer stmt.Close()
	for _, person := range cr {
		_, err := stmt.Exec(person.name, person.secClearance, person.position)
		if err != nil {
			tx.Rollback()
			log.Fatal("Could not query positions ", err)
		}
	}

	tx.Commit()
	return
}

func CSVToMySQL(db *sql.DB) {
	file, err := os.Open("Crews.csv")
	if err != nil {
		log.Fatal("Could not open CSV file", err)
	}
	defer file.Close()

	vargs := []interface{}{}
	sargs := []string{}

	r := csv.NewReader(file)
	r.Comment = '#'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		for _, rec := range record {
			vargs = append(vargs, rec)
		}

		sargs = append(sargs, "(?,?,?,?)")
	}

	insertStmt := fmt.Sprintf("INSERT INTO Personnel (id,Name,SecurityClearance,Position) VALUES %s ", strings.Join(sargs, ","))

	_, err = db.Exec(insertStmt, vargs...)
	if err != nil {
		log.Fatalf("Could not execute insert statement %s with args %s, error %s \n", insertStmt, vargs, err.Error())
	}

}
