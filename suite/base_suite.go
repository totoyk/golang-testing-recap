package suite

import "github.com/stretchr/testify/suite"

type BaseSuite struct {
	suite.Suite
	MockFactory *MockFactory
}

func (s *BaseSuite) SetupTest() {
	s.MockFactory = &MockFactory{}
}

func (s *BaseSuite) TearDownTest() {
	// Clean up resources if needed
}
