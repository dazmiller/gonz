package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BroadBand struct {
	Id                  int       `orm:"column(id);auto"`
	Postcode            string    `orm:"column(postcode);size(6)"`
	MovingOrFirst       string    `orm:"column(moving_or_first);size(30)"`
	ExistingConnection  int8      `orm:"column(existing_connection)"`
	ActivePhoneLine     int8      `orm:"column(active_phone_line)"`
	MainReason          string    `orm:"column(main_reason);size(100)"`
	HouseholdUse        string    `orm:"column(household_use);size(100)"`
	PeopleNetUse        int       `orm:"column(people_net_use)"`
	NumDevices          string    `orm:"column(num_devices);size(100)"`
	StreamingTv         int8      `orm:"column(streaming_tv)"`
	TvExtrasBundle      int8      `orm:"column(tv_extras_bundle)"`
	PhoneCalls          int8      `orm:"column(phone_calls)"`
	WhatImportant       string    `orm:"column(what_important);size(100)"`
	CurrentProvider     string    `orm:"column(current_provider);size(100)"`
	KnowCost            int8      `orm:"column(know_cost)"`
	DataLimit           int       `orm:"column(data_limit)"`
	CurrentlyContracted int8      `orm:"column(currently_contracted)"`
	NetSpeed            string    `orm:"column(net_speed);size(100)"`
	BrowseEmail         string    `orm:"column(browse_email);size(100)"`
	TimeStreaming       string    `orm:"column(time_streaming);size(100)"`
	MoviesMusic         string    `orm:"column(movies_music);size(100)"`
	Gaming              string    `orm:"column(gaming);size(100)"`
	CreatedAt           time.Time `orm:"column(created_at);type(timestamp)"`
	UpdatedAt           time.Time `orm:"column(updated_at);type(timestamp)"`
	Status              int       `orm:"column(status)"`
	UserId              int       `orm:"column(user_id);null"`
	BatchId             int       `orm:"column(batch_id);null"`
}

func (t *BroadBand) TableName() string {
	return "broad_band"
}

func init() {
	orm.RegisterModel(new(BroadBand))
}

// AddBroadBand insert a new BroadBand into database and returns
// last inserted Id on success.
func AddBroadBand(m *BroadBand) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBroadBandById retrieves BroadBand by Id. Returns error if
// Id doesn't exist
func GetBroadBandById(id int) (v *BroadBand, err error) {
	o := orm.NewOrm()
	v = &BroadBand{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBroadBand retrieves all BroadBand matches certain condition. Returns empty list if
// no records exist
func GetAllBroadBand(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BroadBand))
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

	var l []BroadBand
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

// UpdateBroadBand updates BroadBand by Id and returns error if
// the record to be updated doesn't exist
func UpdateBroadBandById(m *BroadBand) (err error) {
	o := orm.NewOrm()
	v := BroadBand{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBroadBand deletes BroadBand by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBroadBand(id int) (err error) {
	o := orm.NewOrm()
	v := BroadBand{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BroadBand{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
