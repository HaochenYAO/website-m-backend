package models

import (
	"fmt"

	"github.com/revel/revel"
)

type Regions struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Pinyin string `json:"pinyin"`
	Maplat string `json:"maplat"`
	Maplng string `json:"maplng"`
	Scope  string `json:"scope"`
	Sort   int64  `json:"sort"`
}

func (r *Regions) String() string {
	return fmt.Sprintf("Regions(%s)", r.Name)
}

func (region *Regions) Validate(v *revel.Validation) {
}
