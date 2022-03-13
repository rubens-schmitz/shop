package access

import (
	"database/sql"
	"log"

	"github.com/rubens-schmitz/shop/util"
	"github.com/sethvargo/go-password/password"
)

func CreateAccess(class string) int {
	code := generateCode()
	query := `insert into access (class, code) values ($1, $2) returning id`
	row := util.DB.QueryRow(query, class, code)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func GetCode(id int) string {
	query := `select code from access where id = $1`
	row := util.DB.QueryRow(query, id)
	var code string
	err := row.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}
	return code
}

func generateCode() string {
	var code string
	var err error
	for {
		code, err = password.Generate(64, 10, 10, false, false)
		if err != nil {
			log.Fatal(err)
		}
		query := `select id from access where code = $1`
		row := util.DB.QueryRow(query, code)
		var id int
		err := row.Scan(&id)
		if err == sql.ErrNoRows {
			break
		}
	}
	return code
}
