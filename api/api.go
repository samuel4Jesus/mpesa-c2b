package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"mpesa/interfaces"
	"mpesa/utils"
	"net/http"
	"strconv"
	"strings"
)

func MPesaC2BStagging(writer http.ResponseWriter, request *http.Request) {

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("M-Pesa Collection Request call")

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/xml")
	writer.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	var transaction interfaces.Transaction
	xml.Unmarshal(body, &transaction)

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("M-Pesa Collection Single-stage Request :\n %s \n", string(body))

	stagging := interfaces.Response{}

	check := utils.SearchDuplicates(transaction.Receipt)

	if !check {
		if strings.EqualFold(transaction.Service, "VUNADEILI") && strings.EqualFold(transaction.Reference, "664455") {

			stagging.Receipt = "30ZX987D"
			stagging.Result = "SUCCESS"

			amount, err := strconv.Atoi(transaction.Amount)
			if err != nil {
				// handle error
				panic(err)
			}
			utils.SaveTransactions(transaction.Receipt, string(utils.GeTUUID()), transaction.Msisdn, "200", amount)
			go utils.SendTransactionalSMS(transaction.Msisdn, string(utils.GeTUUID()))
		}
	} else {
		stagging.Receipt = "30ZX987D"
		stagging.Result = "DUPLICATE"
	}

	d, err := xml.MarshalIndent(&stagging, "", "  ")
	if err != nil {
		log.Fatalf("xml.MarshalIndent failed with '%s'\n", err)
	}

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("M-Pesa Collection, Single-stage response XML :\n %s \n", string(d))

	fmt.Fprintf(writer, string(d))
}
