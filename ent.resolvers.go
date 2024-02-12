package archivista

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/in-toto/archivista/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// AttributeAssertions is the resolver for the attributeAssertions field.
func (r *queryResolver) AttributeAssertions(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.AttributeAssertionWhereInput) (*ent.AttributeAssertionConnection, error) {
	panic(fmt.Errorf("not implemented: AttributeAssertions - attributeAssertions"))
}

// Dsses is the resolver for the dsses field.
func (r *queryResolver) Dsses(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DsseWhereInput) (*ent.DsseConnection, error) {
	return r.client.Dsse.Query().Paginate(ctx, after, first, before, last, ent.WithDsseFilter(where.Filter))
}

// Subjects is the resolver for the subjects field.
func (r *queryResolver) Subjects(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.SubjectWhereInput) (*ent.SubjectConnection, error) {
	return r.client.Subject.Query().Paginate(ctx, after, first, before, last, ent.WithSubjectFilter(where.Filter))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
