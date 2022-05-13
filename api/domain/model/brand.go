package model

import (
	"fmt"
)

type Brand struct {
	ID    int64
	Name  string
	empty bool
}

func EmptyBrand() *Brand {
	return &Brand{empty: true}
}

func (p *Brand) Empty() (bool, error) {
	if p.empty {
		return true, fmt.Errorf("brand is empty")
	}
	return false, nil
}
