package validation

import "regexp"

func IsValidEmail(email string) bool {
	// Basit regex kullanarak e-posta doğrulaması
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
