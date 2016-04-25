package example

import (
	"fmt"
	"packages/ebay-sdk-go"
)

func GetOfficialTime() {
	creds := ebay.Creds{AppName: "",
		DevName:  "",
		CertName: "",
		RuName:   ""}

	ebay.InitSandBox(creds)

	cInput := &ebay.ConnectInput{
		Token:   "Your Sandbox User Token",
		SiteID:  "0",
		Sandbox: true,
	}

	user := ebay.Connect(cInput)

	err, result := user.GeteBayOfficialTime()
	if err != nil {
		fmt.Println(err.Errors[0].LongMessage)
	} else {
		fmt.Println(result.Timestamp)
	}
}
