package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		APIToken string `json:"api_token"`
		Status   string `json:"status"`
	} `json:"data"`
}

type SMSRequest struct {
	SenderID    string `json:"senderId"`
	Message     string `json:"message"`
	CountryCode string `json:"countryCode"`
	Msisdn      string `json:"msisdn"`
	ReferenceID string `json:"referenceId"`
}

type SMSRequestResponse struct {
	Response string `json:"response"`
	Data     struct {
		SenderID    string `json:"senderId"`
		Message     string `json:"message"`
		CountryCode string `json:"countryCode"`
		Msisdn      string `json:"msisdn"`
		ReferenceID string `json:"referenceId"`
		UserID      int    `json:"user_id"`
		Mno         string `json:"mno"`
	} `json:"data"`
}

var apiToken string
var loginResponse LoginResponse

func SendSMS(sms string, msisdn string, token string) {
	// Extract and check if the API Token is empty

	if len(apiToken) == 0 {
		// The API Token is either not present or empty
		//fmt.Println("API Token is empty or not present")
		apiToken = Login()
		SendNotifictaion(sms, msisdn, token, apiToken)
	} else {
		// Sending a message
		SendNotifictaion(sms, msisdn, token, apiToken)
	}

}

func Login() string {
	// Replace the URL with your actual API endpoint
	apiURL := "https://wrapper.emalify.com/api/user/login"

	// Create a sample login request
	loginRequest := LoginRequest{
		Name:     "RT_Tanzania",
		Email:    "samwel.lawrence@roamtech.com",
		Password: "password",
	}

	// Convert the login request to JSON
	requestBody, err := json.MarshalIndent(loginRequest, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	//fmt.Println("The requests is \n\n" + string(requestBody))

	// Send the POST request
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
	}

	// Extract and print the API Token
	apiToken := loginResponse.Data.APIToken
	//fmt.Printf("API Token: %s\n\n", apiToken)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("Wrapper api login, successful")
	return apiToken
}

func SendNotifictaion(sms string, msisdn string, SMStoken string, token string) {

	apiURL := "https://wrapper.emalify.com/api/wrap/mt"
	//bearerToken := apiToken // Replace with your actual Bearer token

	// Create a sample SMS request
	smsRequest := SMSRequest{
		SenderID:    "VUNADEILI",
		Message:     sms,
		CountryCode: "255",
		Msisdn:      msisdn,
		ReferenceID: SMStoken,
	}

	// Convert the SMS request to JSON
	requestBody, err := json.Marshal(smsRequest)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header with Bearer token
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	//fmt.Println(token)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	// Decode the response
	var smsResponse SMSRequestResponse
	err = json.NewDecoder(resp.Body).Decode(&smsResponse)
	//fmt.Println("The requests is \n\n" + string(smsResponse.Response))
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	//fmt.Println("The response is " + resp.Status)

	// Extract and compare ReferenceID
	requestReferenceID := smsRequest.ReferenceID
	responseReferenceID := smsResponse.Data.ReferenceID

	// fmt.Printf("Request ReferenceID: %s\n", requestReferenceID)
	// fmt.Printf("Response ReferenceID: %s\n", responseReferenceID)

	if requestReferenceID == responseReferenceID {
		log.SetPrefix("LOG: ")
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
		log.Printf("Message sent to msisdn: %s", string(smsResponse.Data.Msisdn))
	} else {
		fmt.Println("ReferenceIDs do not match.")
	}
}
