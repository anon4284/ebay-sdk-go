package ebay

import "net/http"

//Creds credentials for your ebay app
type Creds struct {
	RuName   string
	AppName  string
	DevName  string
	CertName string
}

//User who makes the request
type User struct {
	Token   string
	Headers http.Header
	Site    string
}

//ConnectInput input used to connect a user
type ConnectInput struct {
	Token   string `required:"true"`
	SiteID  string
	Sandbox bool
}

var credsProduction Creds

var credsSandbox Creds

//InitProduction the app with your app info
func InitProduction(input Creds) {
	credsProduction = input
}

//InitSandBox adds sandbox functionality to the app
func InitSandBox(input Creds) {
	credsSandbox = input
}

//Connect to ebay
func Connect(input *ConnectInput) User {
	if len(input.Token) < 8 {
		panic("Please Provide a valid user token")
	}
	return User{Token: input.Token, Headers: baseHeader(input), Site: site(input.Sandbox)}
}
