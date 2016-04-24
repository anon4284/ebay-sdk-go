package ebay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

type GeteBayOfficialTime struct {
}

//GeteBayOfficialTime Gets the official time of the ebay site
func (u *User) GeteBayOfficialTime() {
	addHeaderParam(u.Headers, "GetEbayOfficialTime")

	params := GeteBayOfficialTime{}
	myXML, err := xml.Marshal(params)
	checkErr(err)
	myXML = xmlHeadCreds(myXML, u.Token)
	fmt.Print(string(myXML))

	r := bytes.NewReader([]byte("hello"))

	client := &http.Client{}

	req, err := http.NewRequest("POST", u.Site, r)
	checkErr(err)
	req.Header = u.Headers
	resp, err := client.Do(req)
	checkErr(err)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//s := buf.String()

	//fmt.Print(s)

}
