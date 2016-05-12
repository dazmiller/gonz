package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HomeInsuranceQuotes struct {
	Id                     int       `orm:"column(id);auto"`
	RequestId              int       `orm:"column(request_id);null"`
	UserId                 int       `orm:"column(user_id);null"`
	ProviderId             int       `orm:"column(provider_id);null"`
	CreatedAt              time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(timestamp);null"`
	Status                 int       `orm:"column(status);null"`
	Open                   int       `orm:"column(open);null"`
	MinPrice               int       `orm:"column(min_price);null"`
	MaxPrice               int       `orm:"column(max_price);null"`
	HomeInsuranceProfileId int       `orm:"column(home_insurance_profile_id);null"`
}

func (t *HomeInsuranceQuotes) TableName() string {
	return "home_insurance_quotes"
}

func init() {
	orm.RegisterModel(new(HomeInsuranceQuotes))
}

// AddHomeInsuranceQuotes insert a new HomeInsuranceQuotes into database and returns
// last inserted Id on success.
func AddHomeInsuranceQuotes(m *HomeInsuranceQuotes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	return
}

// GetHomeInsuranceQuotesById retrieves HomeInsuranceQuotes by Id. Returns error if
// Id doesn't exist
func GetHomeInsuranceQuotesById(id int) (v *HomeInsuranceQuotes, err error) {
	o := orm.NewOrm()
	v = &HomeInsuranceQuotes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHomeInsuranceQuotes retrieves all HomeInsuranceQuotes matches certain condition. Returns empty list if
// no records exist
func GetAllHomeInsuranceQuotes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HomeInsuranceQuotes))
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

	var l []HomeInsuranceQuotes
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

// UpdateHomeInsuranceQuotes updates HomeInsuranceQuotes by Id and returns error if
// the record to be updated doesn't exist
func UpdateHomeInsuranceQuotesById(m *HomeInsuranceQuotes) (err error) {
	o := orm.NewOrm()
	v := HomeInsuranceQuotes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHomeInsuranceQuotes deletes HomeInsuranceQuotes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHomeInsuranceQuotes(id int) (err error) {
	o := orm.NewOrm()
	v := HomeInsuranceQuotes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HomeInsuranceQuotes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
