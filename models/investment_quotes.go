package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type InvestmentQuotes struct {
	Id                int       `orm:"column(id);auto"`
	RequestId         int       `orm:"column(request_id);null"`
	UserId            int       `orm:"column(user_id);null"`
	ProviderId        int       `orm:"column(provider_id);null"`
	CreatedAt         time.Time `orm:"column(created_at);type(datetime);null"`
	UpdatedAt         time.Time `orm:"column(updated_at);type(datetime);null"`
	LastUpdatedById   int       `orm:"column(last_updated_by_id);null"`
	LastUpdatedByName string    `orm:"column(last_updated_by_name);size(90);null"`
	Status            int       `orm:"column(status);null"`
	Open              uint64    `orm:"column(open);size(1);null"`
	MinPrice          int       `orm:"column(min_price);null"`
	MaxPrice          int       `orm:"column(max_price);null"`
	BatchId           int       `orm:"column(batch_id);null"`
}

func (t *InvestmentQuotes) TableName() string {
	return "investment_quotes"
}

func init() {
	orm.RegisterModel(new(InvestmentQuotes))
}

// AddInvestmentQuotes insert a new InvestmentQuotes into database and returns
// last inserted Id on success.
func AddInvestmentQuotes(m *InvestmentQuotes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInvestmentQuotesById retrieves InvestmentQuotes by Id. Returns error if
// Id doesn't exist
func GetInvestmentQuotesById(id int) (v *InvestmentQuotes, err error) {
	o := orm.NewOrm()
	v = &InvestmentQuotes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInvestmentQuotes retrieves all InvestmentQuotes matches certain condition. Returns empty list if
// no records exist
func GetAllInvestmentQuotes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InvestmentQuotes))
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

	var l []InvestmentQuotes
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

// UpdateInvestmentQuotes updates InvestmentQuotes by Id and returns error if
// the record to be updated doesn't exist
func UpdateInvestmentQuotesById(m *InvestmentQuotes) (err error) {
	o := orm.NewOrm()
	v := InvestmentQuotes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInvestmentQuotes deletes InvestmentQuotes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInvestmentQuotes(id int) (err error) {
	o := orm.NewOrm()
	v := InvestmentQuotes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InvestmentQuotes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
