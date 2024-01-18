package interfaces


import "encoding/xml"

type Transaction struct {
	XMLName         xml.Name `xml:"transaction"`
	Text            string   `xml:",chardata"`
	Timestamp       string   `xml:"timestamp"`
	Service         string   `xml:"service"`
	Reference       string   `xml:"reference"`
	Receipt         string   `xml:"receipt"`
	Amount          string   `xml:"amount"`
	Msisdn          string   `xml:"msisdn"`
	TransactionType string   `xml:"transactionType"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Receipt string   `xml:"receipt"`
	Result  string   `xml:"result"`
}
