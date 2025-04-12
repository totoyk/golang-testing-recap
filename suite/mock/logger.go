package mock

import "fmt"

// MockLogger is a simple mock implementation of a logger.
type MockLogger struct {
	Messages []string
}

// NewMockLogger creates a new instance of MockLogger.
func NewMockLogger() *MockLogger {
	return &MockLogger{
		Messages: []string{},
	}
}

// Println mocks the Println method of a logger.
func (m *MockLogger) Println(v ...interface{}) {
	message := fmt.Sprintln(v...)
	m.Messages = append(m.Messages, message)
}

// Printf mocks the Printf method of a logger.
func (m *MockLogger) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	m.Messages = append(m.Messages, message)
}

// LastMessage returns the last logged message.
func (m *MockLogger) LastMessage() string {
	if len(m.Messages) == 0 {
		return ""
	}
	return m.Messages[len(m.Messages)-1]
}
