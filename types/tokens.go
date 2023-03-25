package types

type GoogleKeys struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	CallbackUrl string `json:"callbackUrl"`
}

type AppleKeys struct {
	ClientId string `json:"clientId"`
	Secret string `json:"secret"`
	CallbackUrl string `json:"callbackUrl"`
}

type SendgridKeys struct {
	ApiKey string `json:"apiKey"`
}

type Tokens struct {
	JwtKey string `json:"jwtKey"`
	Google GoogleKeys `json:"google"`
	Apple AppleKeys `json:"apple"`
	Sendgrid SendgridKeys `json:"sendgrid"`
	// To add other tokens create a struct and add them here,
	// make sure to also update tokens.json
}