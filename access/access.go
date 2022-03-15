package access

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
	"github.com/sethvargo/go-password/password"
)

func CreateAdminHandler(w http.ResponseWriter, r *http.Request) {
	acc := CreateAccess("admin")
	code := GetCode(acc)
	qrcode := util.EncodeQRCode(code)
	res := types.CreateAdminResponse{Qrcode: qrcode}
	util.WriteAsJSON(w, res)
}

func AdminExistHandler(w http.ResponseWriter, r *http.Request) {
	query := `select id from access where class = 'admin'`
	row := util.DB.QueryRow(query)
	var id int
	err := row.Scan(&id)
	var res types.AdminExistResponse
	if err == sql.ErrNoRows {
		res.Sucess = false
	} else {
		res.Sucess = true
	}
	util.WriteAsJSON(w, res)
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	qrcode := r.FormValue("qrcode")
	code := util.DecodeQRCode(qrcode)
	query := `select id from access where class = $1 and code = $2`
	row := util.DB.QueryRow(query, "admin", code)
	var id int
	err = row.Scan(&id)
	var res types.AdminLoginResponse
	if err == sql.ErrNoRows {
		res.Sucess = false
	} else {
		res.Sucess = true
		addAdminCookie(w, code)
	}
	util.WriteAsJSON(w, res)
}

func addAdminCookie(w http.ResponseWriter, code string) {
	cookie := &http.Cookie{Name: "admin", Value: code, Path: "/"}
	http.SetCookie(w, cookie)
}

func IsAdmin(r *http.Request) bool {
	code, err := r.Cookie("admin")
	if err != nil {
		return false
	}
	query := `select id from access where class = 'admin' and code = $1`
	row := util.DB.QueryRow(query, code.Value)
	var id int
	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

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
