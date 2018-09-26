package user

import (
	"fmt"

	"github.com/jemmycalak/mall-tangsel/service/user"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	masterDB *sqlx.DB
	slaveDB  *sqlx.DB
}

func New(masterDB, slaveDB *sqlx.DB) *Resource {
	r := Resource{
		masterDB: masterDB,
		slaveDB:  slaveDB,
	}
	return &r
}

func (r *Resource) GetUser(userId int64) (user.User, error) {
	u := user.User{
		ID:       userId,
		Name:     "jemmy calak",
		Username: "jemmy-calak",
		Email:    "jemmy@calak.com",
		Phone:    "082269219485",
	}

	return u, nil
}

func (r *Resource) Register(m *user.User) error {
	query := `insert into t_user 
		(
		name, email, password, phonenumber, phoneverified, createat, isdelete
		) 
		values 
		($1, $2, $3, $4, $5, $6, $7)`

	stmn, err := r.masterDB.Prepare(query)
	if err != nil {
		fmt.Println("error prepare query")
		return err
	}
	defer stmn.Close()

	_, err = stmn.Exec(m.Name, m.Email, m.Password, m.Phone, m.PhoneVerified, m.CreateAt, m.IsDelete)
	if err != nil {
		return err
	}

	return nil
}
