package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RealEstate struct {
	Id              int       `orm:"column(id);auto"`
	Title           string    `orm:"column(title);size(255)"`
	Postcode        string    `orm:"column(postcode);size(12)"`
	Newestablished  string    `orm:"column(newestablished);size(32)"`
	Propertytype    string    `orm:"column(propertytype)"`
	Minbeds         int       `orm:"column(minbeds)"`
	Maxbeds         int       `orm:"column(maxbeds)"`
	Minprice        int       `orm:"column(minprice)"`
	Maxprice        int       `orm:"column(maxprice)"`
	Indoorfeatures  string    `orm:"column(indoorfeatures)"`
	Outdoorfeatures string    `orm:"column(outdoorfeatures)"`
	Bathrooms       int       `orm:"column(bathrooms)"`
	Parking         int       `orm:"column(parking)"`
	Landsize        int       `orm:"column(landsize)"`
	Livecloseto     string    `orm:"column(livecloseto)"`
	Liveawayfrom    string    `orm:"column(liveawayfrom)"`
	Comments        string    `orm:"column(comments)"`
	Status          int       `orm:"column(status)"`
	Open            int       `orm:"column(open)"`
	UserId          int       `orm:"column(user_id)"`
	CreatedAt       time.Time `orm:"column(created_at);type(timestamp)"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(timestamp)"`
	BatchId         int       `orm:"column(batch_id);null"`
}

func (t *RealEstate) TableName() string {
	return "real_estate"
}

func init() {
	orm.RegisterModel(new(RealEstate))
}

// AddRealEstate insert a new RealEstate into database and returns
// last inserted Id on success.
func AddRealEstate(m *RealEstate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRealEstateById retrieves RealEstate by Id. Returns error if
// Id doesn't exist
func GetRealEstateById(id int) (v *RealEstate, err error) {
	o := orm.NewOrm()
	v = &RealEstate{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRealEstate retrieves all RealEstate matches certain condition. Returns empty list if
// no records exist
func GetAllRealEstate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RealEstate))
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

	var l []RealEstate
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

// UpdateRealEstate updates RealEstate by Id and returns error if
// the record to be updated doesn't exist
func UpdateRealEstateById(m *RealEstate) (err error) {
	o := orm.NewOrm()
	v := RealEstate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRealEstate deletes RealEstate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRealEstate(id int) (err error) {
	o := orm.NewOrm()
	v := RealEstate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RealEstate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
