package models

type Companies struct {
	Id_RENAME      int    `orm:"column(id)"`
	Name           string `orm:"column(name);size(60);null"`
	BusinessName   string `orm:"column(business_name);size(60);null"`
	Address1       string `orm:"column(address1);size(60);null"`
	Address2       string `orm:"column(address2);size(60);null"`
	Suburb         string `orm:"column(suburb);size(60);null"`
	State          string `orm:"column(state);size(20);null"`
	Postcode       string `orm:"column(postcode);size(4);null"`
	Active         uint64 `orm:"column(active);size(1);null"`
	SubscriptionId int    `orm:"column(subscription_id);null"`
	Tagline        string `orm:"column(tagline);size(60);null"`
	TradingName    string `orm:"column(trading_name);size(60);null"`
}
