package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")

	// Auth service'e istek gönderin
	// Burada auth service URL'sini ve token'ı kullanarak kullanıcının profilini alabilirsiniz.
	profile, err := GetUserFromAuthService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı bilgileri alınamadı"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateUserProfile(c *gin.Context) {
	id := c.Param("id")

	var updatedUser map[string]interface{}
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek verisi"})
		return
	}

	// Auth service'e kullanıcı güncelleme isteği gönderin.
	authServiceURL := "http://auth_service:8080/" + id

	// Güncellenen kullanıcı bilgileri ile PUT isteği oluştur
	body, err := json.Marshal(updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Güncellenen kullanıcı verisi oluşturulamadı"})
		return
	}

	req, err := http.NewRequest(http.MethodPut, authServiceURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yeni istek oluşturulamadı"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Service-Authorization", "USER_SERVICE_API_KEY") //os.Getenv("USER_SERVICE_API_KEY")

	// Kullanıcı token'ını ekleyin
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Auth service'e bağlanılamadı"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla güncellendi"})
}

func ResetUserPassword(c *gin.Context) {
	id := c.Param("id")

	var resetData map[string]string
	if err := c.ShouldBindJSON(&resetData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek verisi"})
		return
	}

	// Auth service'e şifre sıfırlama isteği gönderin.
	authServiceURL := "http://auth_service:8080/" + id + "/reset-password"

	body, err := json.Marshal(resetData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre sıfırlama verisi oluşturulamadı"})
		return
	}

	req, err := http.NewRequest(http.MethodPut, authServiceURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yeni istek oluşturulamadı"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Service-Authorization", "USER_SERVICE_API_KEY")
	// Kullanıcı token'ını ekleyin
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Auth service'e bağlanılamadı"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre sıfırlanamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Şifre başarıyla sıfırlandı"})
}

func DeleteUserProfile(c *gin.Context) {
	id := c.Param("id")

	// Auth service'e kullanıcı silme isteği gönderin.
	authServiceURL := "http://auth_service:8080/" + id

	// Yeni DELETE isteği oluştur
	req, err := http.NewRequest(http.MethodDelete, authServiceURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yeni istek oluşturulamadı"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Service-Authorization", "USER_SERVICE_API_KEY")

	// Kullanıcı token'ını ekleyin
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Auth service'e bağlanılamadı"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı silinemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla silindi"})
}

func GetUserFromAuthService(id string) (interface{}, error) {
	authServiceURL := "http://auth_service:8080/" + id

	// İsteği oluştur
	req, err := http.NewRequest("GET", authServiceURL, nil)
	if err != nil {
		return nil, err
	}

	// Header'a User Service API Key'i ekle
	req.Header.Set("User-Service-Authorization", "USER_SERVICE_API_KEY")

	// HTTP client kullanarak isteği gönder
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth service error: %s", resp.Status)
	}

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
