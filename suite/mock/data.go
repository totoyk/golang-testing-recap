package mock

type MockData struct{}

func (m *MockData) NewHelloWolrd() string {
	return "Hello, World!"
}
