package storage

import (
	"github.com/evassilyev/chgk-telebot-2/internal/core"
	"github.com/kelseyhightower/envconfig"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

type PostgresStorageTestSuite struct {
	suite.Suite
	ps   *PostgresStorage
	rand *rand.Rand
}

func TestPostgresServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresStorageTestSuite))
}

func (s *PostgresStorageTestSuite) TearDownSuite() {
	err := s.ps.db.Close()
	assert.Nil(s.T(), err)
}

func (s *PostgresStorageTestSuite) SetupSuite() {
	var c core.Configuration
	err := envconfig.Process("test", &c)
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), c)
	s.T().Logf("Configuration: %+v\n", c)
	s.ps, err = NewPostgresStorage(
		c.DbUser,
		c.DbPass,
		c.DbHost,
		c.DbName,
		c.DbPort,
	)

	assert.Nil(s.T(), err)
	err = s.ps.db.Ping()
	assert.Nil(s.T(), err)

	s.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (s *PostgresStorageTestSuite) TestAddGroup() {
	// group default values
	packageSize := int64(15)
	questionTypes := pq.Int64Array{1, 2}
	timer := int64(60)
	nextQuestionOnTimer := false
	var gId int64
	gId = s.rand.Int63()
	err := s.ps.AddGroup(gId)
	assert.Nil(s.T(), err)
	g, err := s.ps.GetGroup(gId)
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), g)
	assert.Equal(s.T(), packageSize, g.PackageSize)
	assert.Equal(s.T(), questionTypes, g.QuestionsTypes)
	assert.Equal(s.T(), timer, g.Timer)
	assert.Equal(s.T(), nextQuestionOnTimer, g.NextQuestionOnTimer)
	earliestYear, err := g.EarliestYear.Value()
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), earliestYear)
}
