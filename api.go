package ebay

//GeteBayOfficialTime Gets the official time of the ebay site
func (u *User) GeteBayOfficialTime() (*RequestError, *GeteBayOfficialTimeResponse) {
	addHeaderParam(u.Headers, "GetEbayOfficialTime")

	params := GeteBayOfficialTime{}

	success, anything := xmlRequest(u, params, new(GeteBayOfficialTimeResponse))

	if success {
		return nil, anything.(*GeteBayOfficialTimeResponse)
	}
	return anything.(*RequestError), nil

}
