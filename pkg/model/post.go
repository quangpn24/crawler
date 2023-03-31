package model

import "time"

type DataCrawl struct {
	Source string `json:"source"`
	Link   string `json:"link"`
}

type Post struct {
	BaseModel
	Blocks    []Block    `json:"blocks" gorm:"foreignKey:PostID"`
	URL       string     `json:"url"`
	URLOrigin string     `json:"url_origin" gorm:"unique;not null"`
	Title     string     `json:"title"`
	HTMLData  string     `json:"html_data"`
	PostDated *time.Time `json:"post_dated"`
	User      string     `json:"user"`
	Source    string     `json:"source"`
}
