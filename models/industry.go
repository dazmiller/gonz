package models

type Industry struct {
	Id_RENAME int    `orm:"column(id)"`
	Name      string `orm:"column(name);size(60);null"`
	Code      string `orm:"column(code);size(20);null"`
	Desc      string `orm:"column(desc);size(255);null"`
	ParentId  int    `orm:"column(parent_id);null"`
	Category  string `orm:"column(category);size(40);null"`
}
