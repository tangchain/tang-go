package tangtoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable tangtoml client.
type MockClient struct {
	mock.Mock
}

// GetTangToml is a mocking a method
func (m *MockClient) GetTangToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetTangTomlByAddress is a mocking a method
func (m *MockClient) GetTangTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
