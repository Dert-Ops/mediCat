package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/models"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/internal/jwt"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Şifreyi hash'le
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre hashlenemedi"})
		return
	}

	// Hashlenmiş şifreyi kullanıcı modeline kaydet
	user.PasswordHash = string(hashedPassword)

	// Kullanıcıyı veritabanına ekleyin
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı kaydedilemedi"})
		return
	}

	config.LogrusLogger.Tracef("Kullanıcı oluşturuldu: %v", user)
	// message := "Kullanıcı kaydınız başarıyla tamamlandı. Hoş geldiniz " + user.Username + "!"
	// rabbitmq.PublishMessage("email_queue", message)

	c.JSON(http.StatusCreated, user)
}

func SignIn(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var username, password string // username ve password değişkenlerini burada tanımlıyoruz

	// Önce JSON verisini kontrol et
	if err := c.ShouldBindJSON(&loginData); err == nil {
		// JSON verisi varsa, username ve password'u buradan al
		username = loginData.Username
		password = loginData.Password
	} else {
		// JSON verisi yoksa, URL'den veri al
		username = c.Query("username")
		password = c.Query("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş bilgileri"})
			return
		}
	}

	var user models.User
	// Kullanıcıyı veritabanında ara
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// Şifreyi doğrula
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		config.LogrusLogger.Warnf("Şifre doğrulama hatası: girilen şifre %s - veritabanındaki hashlenmiş şifre %s", password, user.PasswordHash)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz şifre"})
		return
	}

	// JWT token oluştur
	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token oluşturulamadı"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		// HttpOnly: true,
		// Secure:   true, // Sadece HTTPS üzerinden iletilsin
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Başarıyla giriş yaptı
	c.JSON(http.StatusOK, gin.H{"message": "Başarıyla giriş yapıldı", "token": token})
	config.LogrusLogger.Tracef("Bir Kullanıcı Giriş Yaptı")
}

// GetUser kullanıcı bilgilerini getirir.
func GetUser(c *gin.Context) {
	var user models.User
	username := c.Param("id")

	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ------------- [!!!!!!!!]     getuser token validation ypailcak
// ------------- [!!!!!!!!]     jwt token generate kontrol edilcek
// ------------- [!!!!!!!!]     

// func GetUser(c *gin.Context) {
//     // Çerezden token'ı al
//     tokenCookie, err := c.Request.Cookie("token")
//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Token bulunamadı"})
//         return
//     }

//     tokenString := tokenCookie.Value

//     // Token'ı doğrula
//     token, err := jwt.ValidateToken(tokenString)
//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
//         return
//     }

//     // Token geçerli mi?
//     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//         username := claims["username"].(string) // Token'dan username'i al

//         var user models.User
//         // Veritabanında kullanıcıyı ara
//         if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
//             return
//         }

//         // Başarılı durumda kullanıcıyı döndür
//         c.JSON(http.StatusOK, user)
//     } else {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Token doğrulaması başarısız"})
//     }
// }

func UpdateUser(c *gin.Context) {
	var updatedUser struct {
		ProfilePicture  string `json:"profile_picture"`
		FullName        string `json:"fullname"`
		Age             int    `json:"age"`
		Bio             string `json:"bio"`
		GithubAccount   string `json:"github_account"`
		LinkedinAccount string `json:"linkedin_account"`
		GoogleAccount   string `json:"google_account"`
		Job             string `json:"job"`
		FavEmail        string `json:"fav_email"`
		Location        string `json:"location"`
	}

	username := c.Param("id")

	// Kullanıcı güncelleme verisini al
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Veritabanında kullanıcıyı bul
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// Alanları güncelle
	if updatedUser.ProfilePicture != "" {
		user.ProfilePicture = updatedUser.ProfilePicture
	}
	if updatedUser.FullName != "" {
		user.FullName = updatedUser.FullName
	}
	if updatedUser.Age > 0 {
		user.Age = updatedUser.Age
	}
	if updatedUser.Bio != "" {
		user.Bio = updatedUser.Bio
	}
	if updatedUser.GithubAccount != "" {
		user.GithubAccount = updatedUser.GithubAccount
	}
	if updatedUser.LinkedinAccount != "" {
		user.LinkedinAccount = updatedUser.LinkedinAccount
	}
	if updatedUser.GoogleAccount != "" {
		user.GoogleAccount = updatedUser.GoogleAccount
	}
	if updatedUser.Job != "" {
		user.Job = updatedUser.Job
	}
	if updatedUser.FavEmail != "" {
		// Email formatı doğrulama
		if !validation.IsValidEmail(updatedUser.FavEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz e-posta formatı"})
			return
		}
		user.FavEmail = updatedUser.FavEmail
	}
	if updatedUser.Location != "" {
		user.Location = updatedUser.Location
	}

	// Kullanıcıyı kaydet
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	username := c.Param("id")

	// Kullanıcıyı veritabanında ara ve sil
	if err := config.DB.Where("username = ?", username).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla silindi"})
}

// Şifreyi belirli kurallara göre doğrulayan fonksiyon

func ResetUserPassword(c *gin.Context) {
	username := c.Param("id")

	// Yeni şifre verisini al
	var passwordData struct {
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&passwordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz şifre formatı"})
		return
	}

	// Şifreyi doğrula
	if err := validation.ValidatePassword(passwordData.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Veritabanında kullanıcıyı bul
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// Şifreyi hash'le ve güncelle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre hashlenemedi"})
		return
	}
	user.PasswordHash = string(hashedPassword)

	// Kullanıcıyı kaydet
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Şifre başarıyla güncellendi"})
}

func ListUsers(c *gin.Context) {
	var users []models.User

	// Tüm kullanıcıları veritabanından al
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcılar alınamadı"})
		return
	}

	c.JSON(http.StatusOK, users)
}
