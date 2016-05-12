package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ServiceProviders struct {
	Id           int    `orm:"column(id);auto"`
	Name         string `orm:"column(name);size(60)"`
	BusinessName string `orm:"column(business_name);size(60)"`
	Address1     string `orm:"column(address1);size(90)"`
	Address2     string `orm:"column(address2);size(90)"`
	Suburb       string `orm:"column(suburb);size(60)"`
	State        string `orm:"column(state);size(20)"`
	Postcode     string `orm:"column(postcode);size(4)"`
	Active       uint64 `orm:"column(active);size(1)"`
	SubId        int    `orm:"column(sub_id)"`
	Logo         string `orm:"column(logo);size(255)"`
}

func (t *ServiceProviders) TableName() string {
	return "service_providers"
}

func init() {
	orm.RegisterModel(new(ServiceProviders))
}

// AddServiceProviders insert a new ServiceProviders into database and returns
// last inserted Id on success.
func AddServiceProviders(m *ServiceProviders) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetServiceProvidersById retrieves ServiceProviders by Id. Returns error if
// Id doesn't exist
func GetServiceProvidersById(id int) (v *ServiceProviders, err error) {
	o := orm.NewOrm()
	v = &ServiceProviders{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllServiceProviders retrieves all ServiceProviders matches certain condition. Returns empty list if
// no records exist
func GetAllServiceProviders(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ServiceProviders))
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

	var l []ServiceProviders
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

// UpdateServiceProviders updates ServiceProviders by Id and returns error if
// the record to be updated doesn't exist
func UpdateServiceProvidersById(m *ServiceProviders) (err error) {
	o := orm.NewOrm()
	v := ServiceProviders{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteServiceProviders deletes ServiceProviders by Id and returns error if
// the record to be deleted doesn't exist
func DeleteServiceProviders(id int) (err error) {
	o := orm.NewOrm()
	v := ServiceProviders{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ServiceProviders{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
