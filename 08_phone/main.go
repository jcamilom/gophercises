package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gophercises_phone"
)

type phoneNumber struct {
	id     int
	number string
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(createPhoneNumbersTable(db))
	_, err = insertPhone(db, "1234567890")
	must(err)
	_, err = insertPhone(db, "123 456 7891")
	must(err)
	_, err = insertPhone(db, "(123) 456 7892")
	must(err)
	id, err := insertPhone(db, "(123) 456-7893")
	must(err)
	_, err = insertPhone(db, "123-456-7894")
	must(err)
	_, err = insertPhone(db, "123-456-7890")
	must(err)
	_, err = insertPhone(db, "1234567892")
	must(err)
	_, err = insertPhone(db, "(123)456-7892")
	must(err)

	number, err := getPhone(db, id)
	must(err)
	fmt.Println("Number is...", number)

	phones, err := getAllPhones(db)
	must(err)
	for _, phone := range phones {
		fmt.Printf("Workign on... %+v\n", phone)
		number := normalize(phone.number)
		if number != phone.number {
			fmt.Println("Updating or removing...", number)
			existing, err := findPhone(db, number)
			must(err)
			if existing != nil {
				must(deletePhone(db, phone.id))
			} else {
				phone.number = number
				must(updatPhone(db, phone))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func updatPhone(db *sql.DB, phone phoneNumber) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, phone.id, phone.number)
	return err
}

func findPhone(db *sql.DB, number string) (*phoneNumber, error) {
	var phone phoneNumber
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number).Scan(&phone.id, &phone.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &phone, nil
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id).Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func getAllPhones(db *sql.DB) ([]phoneNumber, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []phoneNumber
	for rows.Next() {
		var phone phoneNumber
		if err := rows.Scan(&phone.id, &phone.number); err != nil {
			return nil, err
		}
		ret = append(ret, phone)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`
	_, err := db.Exec(statement)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

/*func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
*/

/* func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
*/
