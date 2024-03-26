package gresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
	"log"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

// GetAllCountries is the resolver for the getAllCountries field.
func (r *queryResolver) GetAllCountries(ctx context.Context) ([]*gmodel.Country, error) {
	qb := table.Country.
		SELECT(
			table.Country.AllColumns,
			table.AdministrativeDivision.AdministrativeDivision,
			table.AdministrativeDivision.Cities,
			table.Currency.AllColumns,
		).
		FROM(
			table.Country.
				INNER_JOIN(table.AdministrativeDivision, table.AdministrativeDivision.CountryCode.EQ(table.Country.Code)).
				INNER_JOIN(table.Currency, table.Currency.CurrencyCode.EQ(table.Country.Currency)),
		)
	var countries []*gmodel.Country
	if err := qb.QueryContext(ctx, r.AppContext.DB, &countries); err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not query countries data")
	}
	return countries, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
