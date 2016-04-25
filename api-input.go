package ebay

//GeteBayOfficialTime input struct
type GeteBayOfficialTime struct {
}

//AddItem list a product on ebay
type AddItem struct {
	Item Item
}

//Item of ebayItem
type Item struct {
	ConditionID          string
	ConditionDescription string
	Country              string
	CrossBorderTrade     string
	Currency             string
	Description          string
	DiscountPriceInfo    string
	EBayPlus             string
	ItemSpecifis         ItemSpecifis `xml:"ItemSpecifis>NameValueList"`
	PayPalEmailAddress   string
	PostalCode           string
}

//ItemSpecifis for item
type ItemSpecifis struct {
	Name  string
	Value string
}
