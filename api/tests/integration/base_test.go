package integration

import (
	"api/app"
	"api/app/database"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
}

var (
	baseURL = "http://localhost:8080/v1"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &SuiteTest{})
}

func (s *SuiteTest) SetupSuite() {
	serverReady := make(chan bool)

	app := app.App{
		ServerReady: serverReady,
	}

	go app.StartApp()
	<-serverReady

	db, err := database.Database.DB()
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxOpenConns(1)
}

func (s *SuiteTest) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *SuiteTest) SetupTest() {
	sqlByte, err := ioutil.ReadFile("../sql/create_customer_table.sql")
	if err != nil {
		log.Print(err)
	}

	database.Database.Exec(string(sqlByte))
}

func (s *SuiteTest) TearDownTest() {
	database.Database.Exec("DROP TABLE customer")
}

func (s *SuiteTest) seedMultipleCustomers() {
	sqlByte, err := ioutil.ReadFile("../sql/insert_customers.sql")
	if err != nil {
		log.Print(err)
	}

	database.Database.Exec(string(sqlByte))
}
