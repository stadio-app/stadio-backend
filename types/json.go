package types

type Tokens struct {
	JwtKey string `json:"jwtKey"`
	// To add other tokens create a struct and add them here,
	// make sure to also update ./tokens.json
}
