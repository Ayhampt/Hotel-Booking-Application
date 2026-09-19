package dto

type MailPayload struct {
	To         string     `json:"to"`
	Subject    string     `json:"subject"`
	TemplateID string     `json:"templateId"`
	Params     MailParams `json:"params"`
}

type MailParams struct {
	Name            string `json:"name"`
	AppName         string `json:"appName"`
	VerificationURL string `json:"verificationUrl"`
}
