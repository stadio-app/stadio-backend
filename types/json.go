package types

type SendGridTemplates struct {
	EmailVerification string `json:"emailVerification"`
}

type SendGridTokens struct {
	ApiKey       string            `json:"apiKey"`
	RecoveryCode string            `json:"recoveryCode,omitempty"`
	ApiKeyId     string            `json:"apiKeyId,omitempty"`
	Templates    SendGridTemplates `json:"templates,omitempty"`
}

type CloudflareTokens struct {
	ApiKey   string `json:"apiKey"`
	ApiEmail string `json:"apiEmail"`
}

type Tokens struct {
	JwtKey     string           `json:"jwtKey"`
	SendGrid   SendGridTokens   `json:"sendGrid"`
	Cloudflare CloudflareTokens `json:"cloudflare"`
	// To add other tokens create a struct and add them here,
	// make sure to also update ./tokens.json
}
