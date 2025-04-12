package suite

type MockFactory struct{}

func (m *MockFactory) NewHelloWolrd() string {
	return "Hello, World!"
}
