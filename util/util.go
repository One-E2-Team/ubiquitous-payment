package util

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/smtp"
	"os"
	"strconv"
)

func SendMail(sendTo string, subject string, mailMessage string) {
	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	to := []string{sendTo}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	msg := []byte("To: " + sendTo + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + mailMessage + "\r\n")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func Uint2String(input uint) string {
	return strconv.FormatUint(uint64(input), 10)
}

func String2Uint(input string) uint {
	u64, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return uint(u64)
}

func GetLoggingStringFromID(id uint) string {
	return "profileId: '" + Uint2String(id) + "'"
}

func MongoID2String(mongoID primitive.ObjectID) string {
	return mongoID.Hex()
}

func String2MongoID(stringID string) primitive.ObjectID {
	mongoID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		return primitive.ObjectID{}
	}
	return mongoID
}

func Contains(element interface{}, list []interface{}) bool {
	for _, el := range list {
		if el == element {
			return true
		}
	}
	return false
}
