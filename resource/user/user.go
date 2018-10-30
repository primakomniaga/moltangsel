package user

import (
	"fmt"

	"github.com/jemmycalak/mall-tangsel/service/user"
	"github.com/jmoiron/sqlx"
	// "github.com/mitchellh/mapstructure"
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

func (r *Resource) GetUser(userId int) (map[string]interface{}, error) {
	querys := `select name, username, email, phonenumber, birthday, phoneverified,
	 profilepicture, gender, level from t_user where userid = $1 and isdelete = $2`

	/*
		*belum terealisasi
		var m user.User
		if err := r.masterDB.QueryRow(querys, userId, false).Scan(&m.Name, &m.Username, &m.Email, &m.Phone, &m.Birthday,
			&m.PhoneVerified, &m.ProfilePicture, &m.Gender, &m.Level); err != nil {
			fmt.Println("err ", err)
			return nil, err
		}
	*/

	result := make(map[string]interface{})
	row, err := r.masterDB.Queryx(querys, userId, false)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err = row.MapScan(result)
		if err != nil {
			return nil, err
		}
	}
	// if err := mapstructure.Decode(result, &m); err != nil {
	// 	fmt.Println("error ", err)
	// 	return nil, err
	// }
	return result, nil
}

func (r *Resource) Register(m *user.User) error {
	query := `insert into t_user 
		(
		name, email, password, phonenumber, phoneverified, createat, isdelete, level
		) 
		values 
		($1, $2, $3, $4, $5, $6, $7, $8)`

	stmn, err := r.masterDB.Prepare(query)
	if err != nil {
		fmt.Println("error prepare query")
		return err
	}
	defer stmn.Close()

	_, err = stmn.Exec(m.Name, m.Email, m.Password, m.Phone, m.PhoneVerified, m.CreateAt, m.IsDelete, m.Level)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resource) Login(m *user.Login) (*user.User, error) {
	querys := `select userid, password from t_user where email = $1`
	var model user.User

	if err := r.masterDB.QueryRow(querys, m.Email).Scan(&model.ID, &model.Password); err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *Resource) ValidEmailPhone(m *user.User) bool {
	var c user.User
	querys := `select email, phonenumber from t_user where email = $1 or phonenumber = $2`
	err := r.masterDB.QueryRowx(querys, m.Email, m.Phone).Scan(&c.Email, &c.Phone)
	if err != nil {
		fmt.Println("false")
		return false
	}
	return true
}
