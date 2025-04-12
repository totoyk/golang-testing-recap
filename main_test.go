package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
	common_suite "github.com/totoyk/golang-testing-recap/suite"
)

type TestSuite struct {
	common_suite.BaseSuite
}

func (s *TestSuite) TestAdd() {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 5, 3, 8},
		{"Add negative numbers", -2, -3, -5},
		{"Add positive and negative number", 7, -4, 3},
		{"Add zero", 0, 5, 5},
		{"Add two zeros", 0, 0, 0},
		{"Add large numbers", 1000000, 2000000, 3000000},
		{"Add small numbers", -1000000, -2000000, -3000000},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := Add(tt.a, tt.b)
			s.Equal(tt.expected, result, "Add(%d, %d) should be %d", tt.a, tt.b, tt.expected)
		})
	}
}

func TestAddTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
