package mask

import (
	"errors"
	"fmt"
	"strings"
)

// MaskEmail masks the email address
// max.mustermann@example.com to  ma************@*******.com
// tries to keep the first two characters of the local part
// keeps the @ as is
// keeps dot and the tld as is
// if the local part is shorter than 3 characters, the first character is kept
func MaskEmail(email string) (string, error) {

	if email == "" {
		return "", errors.New("email is empty")
	}

	localPartIndex := strings.Index(email, "@")

	if localPartIndex == -1 {
		return "", errors.New("email faulty, no @ found")
	}

	localPart := email[:localPartIndex]
	domainPart := email[localPartIndex+1:]

	tldIndex := strings.LastIndex(domainPart, ".")
	if tldIndex == -1 {
		return "", errors.New("email faulty, domain faulty")
	}

	tld := domainPart[tldIndex+1:]

	maskedLocalPart := localPart[:3] + strings.Repeat("*", len(localPart)-3)
	maskedDomainPart := strings.Repeat("*", len(domainPart)-len(tld)-1)

	return fmt.Sprintf("%v@%v.%v", maskedLocalPart, maskedDomainPart, tld), nil
}
