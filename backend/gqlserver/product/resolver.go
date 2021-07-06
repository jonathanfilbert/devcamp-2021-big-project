package product

import (
	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
)

type Resolver struct {
	p *productmodule.Module
}

func NewResolver(p *productmodule.Module) *Resolver {
	return &Resolver{p:p}
}


