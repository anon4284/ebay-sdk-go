package ebay

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

func baseHeader(input *ConnectInput) http.Header {
	headers := http.Header{}
	headers.Add("Content-Type", "text/xml")
	headers.Add("X-EBAY-API-SITEID", input.SiteID)
	headers.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "951")
	headers.Add("X-EBAY-API-DETAIL-LEVEL", "0")

	if input.Sandbox {
		headers.Add("X-EBAY-API-DEV-NAME", credsSandbox.DevName)
		headers.Add("X-EBAY-API-APP-NAME", credsSandbox.AppName)
		headers.Add("X-EBAY-API-CERT-NAME", credsSandbox.CertName)
		return headers
	}
	headers.Add("X-EBAY-API-DEV-NAME", credsProduction.DevName)
	headers.Add("X-EBAY-API-APP-NAME", credsProduction.AppName)
	headers.Add("X-EBAY-API-CERT-NAME", credsProduction.CertName)
	return headers
}

func xmlHeadCreds(str []byte, token string) []byte {

	temp := string(str)
	rqc := "<RequesterCredentials><eBayAuthToken>" + token + "</eBayAuthToken></RequesterCredentials>"

	index := strings.IndexAny(temp, ">") + 1
	start := temp[1 : index-1]
	end := temp[index:len(temp)]

	startx := "<" + start + ` xmlns="urn:ebay:apis:eBLBaseComponents"` + ">"
	full := xml.Header + startx + rqc + end

	return []byte(full)
}

func addHeaderParam(header http.Header, request string) http.Header {
	header.Add("X-EBAY-API-CALL-NAME", request)
	return header
}

func site(sandbox bool) string {
	if sandbox {
		return "https://api.sandbox.ebay.com/ws/api.dll"
	}
	return "https://api.ebay.com/ws/api.dll"
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
