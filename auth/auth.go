package auth

type Authentication struct {
	Key    string
	Secret string
}

// GetKey retrieves the authentication key
func (a *Authentication) GetKey() string {
	return a.Key
}

// GetSecret retrieves the authentication secret
func (a *Authentication) GetSecret() string {
	return a.Secret
}

// SetKey sets the passed string as the authentication key
func (a *Authentication) SetKey(k string) {
	a.Key = k
}

// SetSecret  sets the passed string as the authentication secret
func (a *Authentication) SetSecret(s string) {
	a.Secret = s
}
