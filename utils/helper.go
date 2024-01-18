package utils

import (
	"context"
	"fmt"
	"log"
	"mpesa/database"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func SendTransactionalSMS(msisdn string, uuid string) {
	var sms string = "Hongera! Ume SPIN na umepata TOKEN " + string(uuid) + " Tazama ITV 05/JAN saa 12:10 jioni kujua kama umeshinda GARI au MILIONI 1. Cheza zaidi kuongeza nafasi ya ushindi"
	//var sms string = "Hongera! Ume SPIN na umepata " + string(uuid) + " Tazama ITV leo saa 12:10 jioni kujua kama umeshinda GARI au MILIONI 1. Cheza zaidi kuongeza nafasi ya ushindi"

	if len(msisdn) == 9 {
		var editedMSISDN = Append255String(msisdn)
		SaveNotificationSMS(editedMSISDN, sms)
		SendSMS(sms, editedMSISDN, string(uuid))
	} else {
		SaveNotificationSMS(msisdn, sms)
		SendSMS(sms, msisdn, string(uuid))
	}

}

func CompareFloats(num1, num2, tolerance float64) int {
	//fmt.Println("Num1 is s%", num1)
	//fmt.Println("Num2 is s%", num2)
	if num1 == num2 {
		return 0 // num1 is equal to num2 within the tolerance
	} else if num1 > num2 {
		return 1 // num1 is greater than num2
	} else {
		return -1 // num1 is less than num2
	}
}

func Append255String(msisdn string) string {
	if len(msisdn) == 9 {
		return "255" + msisdn
	}
	return msisdn
}

func SaveTransactions(txnid string, customerreference string, msisdn string, error_code string, amount int) {
	query := "INSERT INTO `transactions` (`created_at`, `txn_id` , `msisdn`, `amount`, `custref_id`, `status`, `mno`) VALUES (?, ?, ?, ?, ?, ?,?)"
	insertResult, err := database.DB.ExecContext(context.Background(), query, time.Now(), txnid, msisdn, amount, customerreference, "SUCCESS", "M-Pesa")
	if err != nil {
		log.Fatalf("Transaction saving failed: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("Error retrieving new ID : %s", err)
	}

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf(" C2B response sent is :%s", string(rune(id)))
}

func GeTUUID() []byte {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	var token = strings.ReplaceAll(string(newUUID), "-", "")
	return []byte(token[0:10])
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("Health Check :%s", "Entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func SearchDuplicates(transaction_id string) bool {
	rows, err := database.DB.Query("SELECT txn_id, msisdn FROM `transactions` WHERE txn_id =  (?)", transaction_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func MsisdnNumericCheck(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func SaveNotificationSMS(msisdn string, content string) {
	query := "INSERT INTO `sms_outs` (`created_at`, `msisdn`, `content`) VALUES (?, ?, ?)"
	insertResult, err := database.DB.ExecContext(context.Background(), query, time.Now(), msisdn, content)
	if err != nil {
		log.Fatalf("SMS_Out saved successfuly: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("SMS-Out saving failed: %s", err)
	}
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("Data saved successful - SMS_Out Transactional %d", id)
}
