package mock

type MockMailer struct {
	SendEmailFunc func(to string, subject string, body string) error
}

func (m *MockMailer) SendEmail(to string, subject string, body string) error {
	if m.SendEmailFunc != nil {
		return m.SendEmailFunc(to, subject, body)
	}
	return nil
}
