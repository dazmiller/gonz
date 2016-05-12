package models

type PostsCategories struct {
	PostUuid int    `orm:"column(post_uuid);null"`
	Category string `orm:"column(category);null"`
}
