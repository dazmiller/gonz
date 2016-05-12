package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UsersOld struct {
	Id            int       `orm:"column(id);auto"`
	Email         string    `orm:"column(email);size(191);null"`
	Username      uint8     `orm:"column(username);null"`
	Salt          string    `orm:"column(salt);size(191);null"`
	Password      string    `orm:"column(password);size(191);null"`
	Created       time.Time `orm:"column(created);type(datetime);null"`
	LastLogin     string    `orm:"column(last_login);size(191);null"`
	LoginAttempts uint      `orm:"column(login_attempts);null"`
	Firstname     string    `orm:"column(firstname);size(191);null"`
	Lastname      string    `orm:"column(lastname);size(191);null"`
	Bean          uint8     `orm:"column(bean);null"`
	State         string    `orm:"column(state);size(191);null"`
	Postcode      float64   `orm:"column(postcode);null"`
	City          string    `orm:"column(city);size(191);null"`
	Housenum      uint8     `orm:"column(housenum);null"`
	Housename     uint8     `orm:"column(housename);null"`
	Street        uint8     `orm:"column(street);null"`
	Street2       uint8     `orm:"column(street2);null"`
	Phone1        uint8     `orm:"column(phone1);null"`
	Phone2        uint8     `orm:"column(phone2);null"`
	Mob           uint8     `orm:"column(mob);null"`
	Contactemail  uint8     `orm:"column(contactemail);null"`
	Gender        uint8     `orm:"column(gender);null"`
	Dob           uint8     `orm:"column(dob);null"`
}

func (t *UsersOld) TableName() string {
	return "users_old"
}

func init() {
	orm.RegisterModel(new(UsersOld))
}

// AddUsersOld insert a new UsersOld into database and returns
// last inserted Id on success.
func AddUsersOld(m *UsersOld) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersOldById retrieves UsersOld by Id. Returns error if
// Id doesn't exist
func GetUsersOldById(id int) (v *UsersOld, err error) {
	o := orm.NewOrm()
	v = &UsersOld{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsersOld retrieves all UsersOld matches certain condition. Returns empty list if
// no records exist
func GetAllUsersOld(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UsersOld))
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

	var l []UsersOld
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

// UpdateUsersOld updates UsersOld by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersOldById(m *UsersOld) (err error) {
	o := orm.NewOrm()
	v := UsersOld{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsersOld deletes UsersOld by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsersOld(id int) (err error) {
	o := orm.NewOrm()
	v := UsersOld{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UsersOld{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
