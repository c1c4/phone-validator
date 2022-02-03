package repositories

import (
	"api/app/repositories"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SuiteTest struct {
	suite.Suite
	DB           *gorm.DB
	customerRepo repositories.CustomerRepoInterface
}

func (s *SuiteTest) SetupSuite() {
	var (
		err error
	)

	s.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	
	require.NoError(s.T(), err)

	s.customerRepo = repositories.NewCustomerRepository(s.DB)
}

func TestSuite(t *testing.T) {
	suite.Run(t, &SuiteTest{})
}

func (s *SuiteTest) SetupTest() {
	sqlByte, err := ioutil.ReadFile("../../sql/create_customer_table.sql")
	if err != nil {
		log.Print(err)
	}

	s.DB.Exec(string(sqlByte))
}

func (s *SuiteTest) TearDownTest() {
	s.DB.Exec("DROP TABLE customer")
}

func (s *SuiteTest) TestGetAllCustomers_Success() {
	sqlByte, err := ioutil.ReadFile("../../sql/insert_customers.sql")
	if err != nil {
		log.Print(err)
	}

	s.DB.Exec(string(sqlByte))

	dbTask := s.customerRepo.GetAll()
	require.Equal(s.T(), len(dbTask), 41)
}

func (s *SuiteTest) TestGetAllCustomers_Empty() {
	dbTask := s.customerRepo.GetAll()
	require.Equal(s.T(), len(dbTask), 0)
}
