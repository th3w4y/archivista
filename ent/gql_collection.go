// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/in-toto/archivista/ent/attestation"
	"github.com/in-toto/archivista/ent/attestationcollection"
	"github.com/in-toto/archivista/ent/attributeassertion"
	"github.com/in-toto/archivista/ent/dsse"
	"github.com/in-toto/archivista/ent/payloaddigest"
	"github.com/in-toto/archivista/ent/signature"
	"github.com/in-toto/archivista/ent/statement"
	"github.com/in-toto/archivista/ent/subject"
	"github.com/in-toto/archivista/ent/subjectdigest"
	"github.com/in-toto/archivista/ent/timestamp"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AttestationQuery) CollectFields(ctx context.Context, satisfies ...string) (*AttestationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AttestationQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(attestation.Columns))
		selectedFields = []string{attestation.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "attestationCollection":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttestationCollectionClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.withAttestationCollection = query
		case "type":
			if _, ok := fieldSeen[attestation.FieldType]; !ok {
				selectedFields = append(selectedFields, attestation.FieldType)
				fieldSeen[attestation.FieldType] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type attestationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AttestationPaginateOption
}

func newAttestationPaginateArgs(rv map[string]any) *attestationPaginateArgs {
	args := &attestationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AttestationWhereInput); ok {
		args.opts = append(args.opts, WithAttestationFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ac *AttestationCollectionQuery) CollectFields(ctx context.Context, satisfies ...string) (*AttestationCollectionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ac, nil
	}
	if err := ac.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ac, nil
}

func (ac *AttestationCollectionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(attestationcollection.Columns))
		selectedFields = []string{attestationcollection.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "attestations":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttestationClient{config: ac.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ac.WithNamedAttestations(alias, func(wq *AttestationQuery) {
				*wq = *query
			})
		case "statement":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatementClient{config: ac.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ac.withStatement = query
		case "name":
			if _, ok := fieldSeen[attestationcollection.FieldName]; !ok {
				selectedFields = append(selectedFields, attestationcollection.FieldName)
				fieldSeen[attestationcollection.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ac.Select(selectedFields...)
	}
	return nil
}

type attestationcollectionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AttestationCollectionPaginateOption
}

func newAttestationCollectionPaginateArgs(rv map[string]any) *attestationcollectionPaginateArgs {
	args := &attestationcollectionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AttestationCollectionWhereInput); ok {
		args.opts = append(args.opts, WithAttestationCollectionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (aa *AttributeAssertionQuery) CollectFields(ctx context.Context, satisfies ...string) (*AttributeAssertionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return aa, nil
	}
	if err := aa.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return aa, nil
}

func (aa *AttributeAssertionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(attributeassertion.Columns))
		selectedFields = []string{attributeassertion.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "attributesReport":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttributeReportClient{config: aa.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			aa.withAttributesReport = query
		case "attribute":
			if _, ok := fieldSeen[attributeassertion.FieldAttribute]; !ok {
				selectedFields = append(selectedFields, attributeassertion.FieldAttribute)
				fieldSeen[attributeassertion.FieldAttribute] = struct{}{}
			}
		case "target":
			if _, ok := fieldSeen[attributeassertion.FieldTarget]; !ok {
				selectedFields = append(selectedFields, attributeassertion.FieldTarget)
				fieldSeen[attributeassertion.FieldTarget] = struct{}{}
			}
		case "conditions":
			if _, ok := fieldSeen[attributeassertion.FieldConditions]; !ok {
				selectedFields = append(selectedFields, attributeassertion.FieldConditions)
				fieldSeen[attributeassertion.FieldConditions] = struct{}{}
			}
		case "evidence":
			if _, ok := fieldSeen[attributeassertion.FieldEvidence]; !ok {
				selectedFields = append(selectedFields, attributeassertion.FieldEvidence)
				fieldSeen[attributeassertion.FieldEvidence] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		aa.Select(selectedFields...)
	}
	return nil
}

type attributeassertionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AttributeAssertionPaginateOption
}

func newAttributeAssertionPaginateArgs(rv map[string]any) *attributeassertionPaginateArgs {
	args := &attributeassertionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AttributeAssertionWhereInput); ok {
		args.opts = append(args.opts, WithAttributeAssertionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ar *AttributeReportQuery) CollectFields(ctx context.Context, satisfies ...string) (*AttributeReportQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ar, nil
	}
	if err := ar.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ar, nil
}

func (ar *AttributeReportQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "attributes":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttributeAssertionClient{config: ar.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ar.WithNamedAttributes(alias, func(wq *AttributeAssertionQuery) {
				*wq = *query
			})
		case "statement":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatementClient{config: ar.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ar.withStatement = query
		}
	}
	return nil
}

type attributereportPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AttributeReportPaginateOption
}

func newAttributeReportPaginateArgs(rv map[string]any) *attributereportPaginateArgs {
	args := &attributereportPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AttributeReportWhereInput); ok {
		args.opts = append(args.opts, WithAttributeReportFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DsseQuery) CollectFields(ctx context.Context, satisfies ...string) (*DsseQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return d, nil
	}
	if err := d.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DsseQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(dsse.Columns))
		selectedFields = []string{dsse.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "statement":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatementClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			d.withStatement = query
		case "signatures":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SignatureClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			d.WithNamedSignatures(alias, func(wq *SignatureQuery) {
				*wq = *query
			})
		case "payloadDigests":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PayloadDigestClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			d.WithNamedPayloadDigests(alias, func(wq *PayloadDigestQuery) {
				*wq = *query
			})
		case "gitoidSha256":
			if _, ok := fieldSeen[dsse.FieldGitoidSha256]; !ok {
				selectedFields = append(selectedFields, dsse.FieldGitoidSha256)
				fieldSeen[dsse.FieldGitoidSha256] = struct{}{}
			}
		case "payloadType":
			if _, ok := fieldSeen[dsse.FieldPayloadType]; !ok {
				selectedFields = append(selectedFields, dsse.FieldPayloadType)
				fieldSeen[dsse.FieldPayloadType] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		d.Select(selectedFields...)
	}
	return nil
}

type dssePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DssePaginateOption
}

func newDssePaginateArgs(rv map[string]any) *dssePaginateArgs {
	args := &dssePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*DsseWhereInput); ok {
		args.opts = append(args.opts, WithDsseFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pd *PayloadDigestQuery) CollectFields(ctx context.Context, satisfies ...string) (*PayloadDigestQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pd, nil
	}
	if err := pd.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pd, nil
}

func (pd *PayloadDigestQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(payloaddigest.Columns))
		selectedFields = []string{payloaddigest.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "dsse":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DsseClient{config: pd.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			pd.withDsse = query
		case "algorithm":
			if _, ok := fieldSeen[payloaddigest.FieldAlgorithm]; !ok {
				selectedFields = append(selectedFields, payloaddigest.FieldAlgorithm)
				fieldSeen[payloaddigest.FieldAlgorithm] = struct{}{}
			}
		case "value":
			if _, ok := fieldSeen[payloaddigest.FieldValue]; !ok {
				selectedFields = append(selectedFields, payloaddigest.FieldValue)
				fieldSeen[payloaddigest.FieldValue] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pd.Select(selectedFields...)
	}
	return nil
}

type payloaddigestPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PayloadDigestPaginateOption
}

func newPayloadDigestPaginateArgs(rv map[string]any) *payloaddigestPaginateArgs {
	args := &payloaddigestPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PayloadDigestWhereInput); ok {
		args.opts = append(args.opts, WithPayloadDigestFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SignatureQuery) CollectFields(ctx context.Context, satisfies ...string) (*SignatureQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SignatureQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(signature.Columns))
		selectedFields = []string{signature.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "dsse":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DsseClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withDsse = query
		case "timestamps":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TimestampClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.WithNamedTimestamps(alias, func(wq *TimestampQuery) {
				*wq = *query
			})
		case "keyID":
			if _, ok := fieldSeen[signature.FieldKeyID]; !ok {
				selectedFields = append(selectedFields, signature.FieldKeyID)
				fieldSeen[signature.FieldKeyID] = struct{}{}
			}
		case "signature":
			if _, ok := fieldSeen[signature.FieldSignature]; !ok {
				selectedFields = append(selectedFields, signature.FieldSignature)
				fieldSeen[signature.FieldSignature] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type signaturePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SignaturePaginateOption
}

func newSignaturePaginateArgs(rv map[string]any) *signaturePaginateArgs {
	args := &signaturePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*SignatureWhereInput); ok {
		args.opts = append(args.opts, WithSignatureFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *StatementQuery) CollectFields(ctx context.Context, satisfies ...string) (*StatementQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *StatementQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(statement.Columns))
		selectedFields = []string{statement.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "subjects":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SubjectClient{config: s.config}).Query()
			)
			args := newSubjectPaginateArgs(fieldArgs(ctx, new(SubjectWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newSubjectPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					s.loadTotal = append(s.loadTotal, func(ctx context.Context, nodes []*Statement) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"statement_subjects"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(statement.SubjectsColumn), ids...))
						})
						if err := query.GroupBy(statement.SubjectsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					s.loadTotal = append(s.loadTotal, func(_ context.Context, nodes []*Statement) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Subjects)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Subject")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(statement.SubjectsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			s.WithNamedSubjects(alias, func(wq *SubjectQuery) {
				*wq = *query
			})
		case "attestationCollections":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttestationCollectionClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withAttestationCollections = query
		case "attributesReport":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AttributeReportClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withAttributesReport = query
		case "dsse":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DsseClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.WithNamedDsse(alias, func(wq *DsseQuery) {
				*wq = *query
			})
		case "predicate":
			if _, ok := fieldSeen[statement.FieldPredicate]; !ok {
				selectedFields = append(selectedFields, statement.FieldPredicate)
				fieldSeen[statement.FieldPredicate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type statementPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StatementPaginateOption
}

func newStatementPaginateArgs(rv map[string]any) *statementPaginateArgs {
	args := &statementPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*StatementWhereInput); ok {
		args.opts = append(args.opts, WithStatementFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SubjectQuery) CollectFields(ctx context.Context, satisfies ...string) (*SubjectQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SubjectQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(subject.Columns))
		selectedFields = []string{subject.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "subjectDigests":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SubjectDigestClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.WithNamedSubjectDigests(alias, func(wq *SubjectDigestQuery) {
				*wq = *query
			})
		case "statement":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatementClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withStatement = query
		case "name":
			if _, ok := fieldSeen[subject.FieldName]; !ok {
				selectedFields = append(selectedFields, subject.FieldName)
				fieldSeen[subject.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type subjectPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SubjectPaginateOption
}

func newSubjectPaginateArgs(rv map[string]any) *subjectPaginateArgs {
	args := &subjectPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*SubjectWhereInput); ok {
		args.opts = append(args.opts, WithSubjectFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sd *SubjectDigestQuery) CollectFields(ctx context.Context, satisfies ...string) (*SubjectDigestQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sd, nil
	}
	if err := sd.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *SubjectDigestQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(subjectdigest.Columns))
		selectedFields = []string{subjectdigest.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "subject":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SubjectClient{config: sd.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			sd.withSubject = query
		case "algorithm":
			if _, ok := fieldSeen[subjectdigest.FieldAlgorithm]; !ok {
				selectedFields = append(selectedFields, subjectdigest.FieldAlgorithm)
				fieldSeen[subjectdigest.FieldAlgorithm] = struct{}{}
			}
		case "value":
			if _, ok := fieldSeen[subjectdigest.FieldValue]; !ok {
				selectedFields = append(selectedFields, subjectdigest.FieldValue)
				fieldSeen[subjectdigest.FieldValue] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		sd.Select(selectedFields...)
	}
	return nil
}

type subjectdigestPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SubjectDigestPaginateOption
}

func newSubjectDigestPaginateArgs(rv map[string]any) *subjectdigestPaginateArgs {
	args := &subjectdigestPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*SubjectDigestWhereInput); ok {
		args.opts = append(args.opts, WithSubjectDigestFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TimestampQuery) CollectFields(ctx context.Context, satisfies ...string) (*TimestampQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TimestampQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(timestamp.Columns))
		selectedFields = []string{timestamp.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "signature":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SignatureClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withSignature = query
		case "type":
			if _, ok := fieldSeen[timestamp.FieldType]; !ok {
				selectedFields = append(selectedFields, timestamp.FieldType)
				fieldSeen[timestamp.FieldType] = struct{}{}
			}
		case "timestamp":
			if _, ok := fieldSeen[timestamp.FieldTimestamp]; !ok {
				selectedFields = append(selectedFields, timestamp.FieldTimestamp)
				fieldSeen[timestamp.FieldTimestamp] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type timestampPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TimestampPaginateOption
}

func newTimestampPaginateArgs(rv map[string]any) *timestampPaginateArgs {
	args := &timestampPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TimestampWhereInput); ok {
		args.opts = append(args.opts, WithTimestampFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
