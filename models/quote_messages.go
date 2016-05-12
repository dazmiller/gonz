package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type QuoteMessages struct {
	Id              int       `orm:"column(id);auto"`
	QuoteId         int       `orm:"column(quote_id);null"`
	ProviderId      int       `orm:"column(provider_id);null"`
	ProviderName    string    `orm:"column(provider_name);size(90);null"`
	Title           string    `orm:"column(title);size(255);null"`
	Body            string    `orm:"column(body);null"`
	Isread          uint64    `orm:"column(isread);size(1);null"`
	CreatedDatetime time.Time `orm:"column(created_datetime);type(datetime);null"`
	ReadDatetime    time.Time `orm:"column(read_datetime);type(datetime);null"`
	ParentId        int       `orm:"column(parent_id);null"`
	UserId          int       `orm:"column(user_id);null"`
}

func (t *QuoteMessages) TableName() string {
	return "quote_messages"
}

func init() {
	orm.RegisterModel(new(QuoteMessages))
}

// AddQuoteMessages insert a new QuoteMessages into database and returns
// last inserted Id on success.
func AddQuoteMessages(m *QuoteMessages) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuoteMessagesById retrieves QuoteMessages by Id. Returns error if
// Id doesn't exist
func GetQuoteMessagesById(id int) (v *QuoteMessages, err error) {
	o := orm.NewOrm()
	v = &QuoteMessages{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQuoteMessages retrieves all QuoteMessages matches certain condition. Returns empty list if
// no records exist
func GetAllQuoteMessages(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QuoteMessages))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []QuoteMessages
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateQuoteMessages updates QuoteMessages by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuoteMessagesById(m *QuoteMessages) (err error) {
	o := orm.NewOrm()
	v := QuoteMessages{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQuoteMessages deletes QuoteMessages by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuoteMessages(id int) (err error) {
	o := orm.NewOrm()
	v := QuoteMessages{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QuoteMessages{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
