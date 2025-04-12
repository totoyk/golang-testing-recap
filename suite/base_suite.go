package suite

import (
	"github.com/stretchr/testify/suite"
	"github.com/totoyk/golang-testing-recap/suite/mock"
)

type BaseSuite struct {
	suite.Suite
	MockData        *mock.MockData
	MockRedisClient *mock.MockRedisClient
	MockMailer      *mock.MockMailer
	MockLogger      *mock.MockLogger
}

func (s *BaseSuite) SetupTest() {
	s.MockData = &mock.MockData{}
	s.MockRedisClient = &mock.MockRedisClient{}
}

func (s *BaseSuite) TearDownTest() {
	// Clean up resources if needed
}
