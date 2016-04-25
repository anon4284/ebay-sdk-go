package ebay

//RequestError if ebay sends error it uses this struct
type RequestError struct {
	Ack    string
	Errors []Errors
}

type Errors struct {
	ErrorClassification string
	ShortMessage        string
	LongMessage         string
	SeverityCode        string
}

//GeteBayOfficialTimeResponse response
type GeteBayOfficialTimeResponse struct {
	Ack       string
	Timestamp string
}
