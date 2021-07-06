package product

import (
	"math"

	"github.com/graphql-go/graphql"
)

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"product_id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"product_name": &graphql.Field{
			}
		}
	}
)