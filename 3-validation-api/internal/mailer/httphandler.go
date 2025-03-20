package mailer

import (
	"encoding/json"
	"fmt"
	"go-training/validation-api/configs"
	"go-training/validation-api/internal/files"
	"go-training/validation-api/pkg/request"
	"go-training/validation-api/pkg/response"
	"net/http"
)

type MailerDeps struct {
	SmtpConfig       configs.SmtpConfig
	SendEmailHandler SendEmailHandler
}

type MailerHttpHandler struct {
	MailerDeps
	jsonStorageFile string
}

func NewMailerHttpHandler(router *http.ServeMux, mailerDeps MailerDeps) {
	handler := &MailerHttpHandler{
		MailerDeps:      mailerDeps,
		jsonStorageFile: "email_to_otp.json",
	}
	router.HandleFunc("POST /email/send-verify-code", handler.SendVerificationCodeByEmail())
	router.HandleFunc("GET /email/{email}/verify/{hash}", handler.VerifyEmail())

}

func (handler *MailerHttpHandler) SendVerificationCodeByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		sendEmailPayload, err := request.Decode[SendEmailPayload](req.Body)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		err = request.Validate(*sendEmailPayload)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		otp := GenerateOtp()
		data := make(map[string]string)
		data[sendEmailPayload.Email] = otp
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		err = files.WriteFile(jsonBytes, handler.jsonStorageFile)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		success, err := handler.SendEmailHandler.Send(
			sendEmailPayload.Email,
			"Verification code", "Verfication link <a href='http://localhost:8081/email/"+sendEmailPayload.Email+"/verify/"+otp+"'>Click to verify</a>",
		)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		response.JsonResponse(w, success, 201)
	}
}

func (handler *MailerHttpHandler) VerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		email := req.PathValue("email")
		data, err := files.ReadFile(handler.jsonStorageFile)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		parsedJson := make(map[string]string)
		err = json.Unmarshal(data, &parsedJson)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		fmt.Println(parsedJson, parsedJson[email])

		// не удалось провалидировать
		if parsedJson[email] != hash {
			response.JsonResponse(w, false, 201)
			return
		}

		// Валдидация успешна, чистим json
		delete(parsedJson, email)
		jsonBytes, err := json.Marshal(parsedJson)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		err = files.WriteFile(jsonBytes, handler.jsonStorageFile)
		if err != nil {
			response.JsonResponse(w, err.Error(), 500)
			return
		}
		response.JsonResponse(w, true, 201)
	}
}
