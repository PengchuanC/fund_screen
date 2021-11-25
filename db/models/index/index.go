package index

import "github.com/golang-module/carbon"

type Index struct {
	Secucode string `json:"secucode,omitempty" gorm:"primaryKey"`
	Secuabbr string `json:"secuabbr,omitempty"`
	Category int    `json:"category,omitempty"`
	Chiname  string `json:"chiname,omitempty" gorm:"char(20)"`
}

func (i Index) TableName() string {
	return "sc_index_secumain"
}

type IQuote struct {
	Id             int64               `json:"id,omitempty"`
	Secucode       string              `json:"secucode,omitempty"`
	Category       string              `json:"category,omitempty"`
	Closeprice     float64             `json:"closeprice,omitempty"`
	Prevcloseprice float64             `json:"prevcloseprice,omitempty"`
	Changepct      float64             `json:"changepct,omitempty"`
	Date           carbon.ToDateString `json:"date"`
	Jsid           int64               `json:"jsid,omitempty"`
}

func (q IQuote) TableName() string {
	return "sc_index_quote"
}