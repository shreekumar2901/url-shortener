package validator

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/shreekumar2901/url-shortener/dto"
)

func RegisterUserValidator(dto *dto.UserRequestDTO) map[string][]string {
	errorMsgs := make(map[string][]string)
	errorMsgs["errors"] = []string{}

	// username should have atleast 5 character
	if len(dto.UserName) < 5 {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Username should be at least 5 characters")
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Invalid email address")
	}

	/* Password validation
	at least 7 characters and at most 15 characters
	at least 1 special char
	at least 1  number
	at least 1 uppercase
	at least 1 lowercase
	*/

	password := strings.TrimSpace(dto.Password)

	if len(password) < 7 {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password must contain at least 7 characters")
	}

	if len(password) > 15 {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password can not contain more than  15 characters")
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password must contain at least 1 digit")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password must contain at least 1 uppercase character")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password must contain at least 1 lowercase character")
	}

	if !regexp.MustCompile(`[#,@,_,+,.,$,=]`).MatchString(password) {
		errorMsgs["errors"] = append(errorMsgs["errors"], "Password must contain at least 1 special characters among #,@,_,+,.,$,=")
	}

	return errorMsgs
}
