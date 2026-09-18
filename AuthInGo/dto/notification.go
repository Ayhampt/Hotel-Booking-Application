package dto

type MailPayload struct {
	To         string     `json:"to"`
	Subject    string     `json:"subject"`
	TemplateID string     `json:"templateId"`
	Params     MailParams `json:"params"`
}

type MailParams struct {
	Token           string `json:"token"`
	VerificationURL string `json:"verificationUrl"`
}
