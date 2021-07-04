package server

import (
	"fmt"
	"io/ioutil"
	"net"

	httpapi "github.com/jemmycalak/mall-tangsel/api/http"
	resourceProduct "github.com/jemmycalak/mall-tangsel/resource/product"
	resourceUser "github.com/jemmycalak/mall-tangsel/resource/user"

	serviceProduct "github.com/jemmycalak/mall-tangsel/service/product"
	serviceUser "github.com/jemmycalak/mall-tangsel/service/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"gopkg.in/yaml.v2"
)

func Mains() error {

	file, err := ioutil.ReadFile("../config/DB.yaml")
	if err != nil {
		fmt.Println("Error 1", err)
		return err
	}

	if err := yaml.Unmarshal(file, &Config); err != nil {
		fmt.Println("Error 2")
		return err
	}

	master := Config.DatabaseMaster
	conM := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", master.Address, master.User, master.Password, master.Database)

	slave := Config.DatabaseSlave
	conS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", slave.Address, slave.User, slave.Password, slave.Database)

	masterDB, err := sqlx.Connect("postgres", conM)
	if err != nil {
		fmt.Println("Error 3")
		return err
	}
	slaveDB, err := sqlx.Connect("postgres", conS)
	if err != nil {
		fmt.Println("Error 4")
		return err
	}

	//user
	resUser := resourceUser.New(masterDB, slaveDB)
	serUser := serviceUser.New(resUser)

	//product
	resProduct := resourceProduct.New(masterDB, slaveDB)
	serProduct := serviceProduct.New(resProduct)

	// create a new Listener for http and grpc server
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("error Listener ", err)
		return err
	}

	httpserver := httpapi.Server{
		UserService:    serUser,
		ProductService: serProduct,
	}

	return httpserver.Serve(listener)
}
