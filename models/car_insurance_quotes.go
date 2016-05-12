package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CarInsuranceQuotes struct {
	Id                int       `orm:"column(id);auto"`
	RequestId         int       `orm:"column(request_id);null"`
	UserId            int       `orm:"column(user_id);null"`
	ProviderId        int       `orm:"column(provider_id);null"`
	CreatedAt         time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt         time.Time `orm:"column(updated_at);type(timestamp);null"`
	LastUpdatedById   int       `orm:"column(last_updated_by_id);null"`
	LastUpdatedByName string    `orm:"column(last_updated_by_name);size(90);null"`
	Status            int       `orm:"column(status);null"`
	Open              uint64    `orm:"column(open);size(1);null"`
	MinPrice          int       `orm:"column(min_price);null"`
	MaxPrice          int       `orm:"column(max_price);null"`
	BatchId           int       `orm:"column(batch_id);null"`
}

func (t *CarInsuranceQuotes) TableName() string {
	return "car_insurance_quotes"
}

func init() {
	orm.RegisterModel(new(CarInsuranceQuotes))
}

// AddCarInsuranceQuotes insert a new CarInsuranceQuotes into database and returns
// last inserted Id on success.
func AddCarInsuranceQuotes(m *CarInsuranceQuotes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCarInsuranceQuotesById retrieves CarInsuranceQuotes by Id. Returns error if
// Id doesn't exist
func GetCarInsuranceQuotesById(id int) (v *CarInsuranceQuotes, err error) {
	o := orm.NewOrm()
	v = &CarInsuranceQuotes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCarInsuranceQuotes retrieves all CarInsuranceQuotes matches certain condition. Returns empty list if
// no records exist
func GetAllCarInsuranceQuotes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CarInsuranceQuotes))
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

	var l []CarInsuranceQuotes
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

// UpdateCarInsuranceQuotes updates CarInsuranceQuotes by Id and returns error if
// the record to be updated doesn't exist
func UpdateCarInsuranceQuotesById(m *CarInsuranceQuotes) (err error) {
	o := orm.NewOrm()
	v := CarInsuranceQuotes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCarInsuranceQuotes deletes CarInsuranceQuotes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCarInsuranceQuotes(id int) (err error) {
	o := orm.NewOrm()
	v := CarInsuranceQuotes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CarInsuranceQuotes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
