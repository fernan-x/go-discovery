package tokenissuer

type DummyTokenIssuer struct {}

var _ TokenIssuer = (*DummyTokenIssuer)(nil)

func NewDummyTokenIssuer() *DummyTokenIssuer {
	return &DummyTokenIssuer{}
}

func (t *DummyTokenIssuer) GenerateAccessToken(userId string) (string, error) {
	return "access-token", nil
}

func (t *DummyTokenIssuer) GenerateRefreshToken(userId string) (string, error) {
	return "refresh-token", nil
}

func (t *DummyTokenIssuer) GenerateRandomToken(size int) (string, error) {
	return "random-token", nil
}

func (t *DummyTokenIssuer) Parse(token string) (map[string]any, error) {
	return map[string]any{"sub": "12345"}, nil
}