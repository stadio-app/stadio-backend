package types

type SendGridTemplates struct {
	EmailVerification string `json:"emailVerification"`
}

type SendGridTokens struct {
	ApiKey string `json:"api_key"`
	RecoveryCode string `json:"recovery_code,omitempty"`
	ApiKeyId string `json:"api_key_id,omitempty"`
	Templates SendGridTemplates `json:"templates,omitempty"`
}

type Tokens struct {
	JwtKey string `json:"jwtKey"`
	SendGrid SendGridTokens `json:"sendGrid,omitempty"`
	// To add other tokens create a struct and add them here,
	// make sure to also update ./tokens.json
}
