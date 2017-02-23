package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type Regions struct {
	id int
	name string
	suzhou string
	scope string
  sort string
  maplat string
  maplng string
}

func (r *Regions) String() string {
	return fmt.Sprintf("Regions(%s)", r.name)
}

func (region *Regions) Validate(v *revel.Validation) {
}
