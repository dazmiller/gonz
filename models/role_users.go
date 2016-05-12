package models

import "time"

type RoleUsers struct {
	UserId    uint      `orm:"column(user_id)"`
	RoleId    uint      `orm:"column(role_id)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
}
