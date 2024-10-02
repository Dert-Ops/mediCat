package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var SecretKey string
var LogrusLogger *logrus.Logger

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		LogrusLogger.Warnf("Error loading .env file")
	}

	SecretKey = os.Getenv("JWT_SECRET_KEY")
}

type GMT3JSONFormatter struct {
	logrus.JSONFormatter
}

func (f *GMT3JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Zaman damgasını GMT+3'e dönüştür
	entry.Time = entry.Time.In(time.FixedZone("GMT+3", 3*3600))
	return f.JSONFormatter.Format(entry)
}

func InitLogrusLogger() {
	LogrusLogger = logrus.New()

	// Logs dizinini kontrol et ve oluştur
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", os.ModePerm)
	}

	// Log dosyasını aç
	file, err := os.OpenFile("/logs/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		LogrusLogger.Warn("Unable to open log file, logging to stdout instead.")
		LogrusLogger.Out = os.Stdout
	} else {
		LogrusLogger.Out = file
	}

	// GMT+3 formatında log ayarları
	LogrusLogger.SetFormatter(&GMT3JSONFormatter{
		JSONFormatter: logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05", // İstediğin zaman formatı
			DisableHTMLEscape: true,
		},
	})

	LogrusLogger.SetLevel(logrus.TraceLevel)
}
