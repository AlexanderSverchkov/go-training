package mailer

type SendEmailPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type SendEmailResponse struct {
	Success   bool    `json:"sucess"`
	ErrorText *string `json:"errorText"`
}

type VerifyResponse struct {
	Success   bool    `json:"sucess"`
	ErrorText *string `json:"errorText"`
}
