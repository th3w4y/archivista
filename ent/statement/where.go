// Code generated by ent, DO NOT EDIT.

package statement

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/in-toto/archivista/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldID, id))
}

// Predicate applies equality check predicate on the "predicate" field. It's identical to PredicateEQ.
func Predicate(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldPredicate, v))
}

// PredicateEQ applies the EQ predicate on the "predicate" field.
func PredicateEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldPredicate, v))
}

// PredicateNEQ applies the NEQ predicate on the "predicate" field.
func PredicateNEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldPredicate, v))
}

// PredicateIn applies the In predicate on the "predicate" field.
func PredicateIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldPredicate, vs...))
}

// PredicateNotIn applies the NotIn predicate on the "predicate" field.
func PredicateNotIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldPredicate, vs...))
}

// PredicateGT applies the GT predicate on the "predicate" field.
func PredicateGT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldPredicate, v))
}

// PredicateGTE applies the GTE predicate on the "predicate" field.
func PredicateGTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldPredicate, v))
}

// PredicateLT applies the LT predicate on the "predicate" field.
func PredicateLT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldPredicate, v))
}

// PredicateLTE applies the LTE predicate on the "predicate" field.
func PredicateLTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldPredicate, v))
}

// PredicateContains applies the Contains predicate on the "predicate" field.
func PredicateContains(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContains(FieldPredicate, v))
}

// PredicateHasPrefix applies the HasPrefix predicate on the "predicate" field.
func PredicateHasPrefix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasPrefix(FieldPredicate, v))
}

// PredicateHasSuffix applies the HasSuffix predicate on the "predicate" field.
func PredicateHasSuffix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasSuffix(FieldPredicate, v))
}

// PredicateEqualFold applies the EqualFold predicate on the "predicate" field.
func PredicateEqualFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEqualFold(FieldPredicate, v))
}

// PredicateContainsFold applies the ContainsFold predicate on the "predicate" field.
func PredicateContainsFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContainsFold(FieldPredicate, v))
}

// HasSubjects applies the HasEdge predicate on the "subjects" edge.
func HasSubjects() predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubjectsTable, SubjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectsWith applies the HasEdge predicate on the "subjects" edge with a given conditions (other predicates).
func HasSubjectsWith(preds ...predicate.Subject) predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := newSubjectsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttestationCollections applies the HasEdge predicate on the "attestation_collections" edge.
func HasAttestationCollections() predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AttestationCollectionsTable, AttestationCollectionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttestationCollectionsWith applies the HasEdge predicate on the "attestation_collections" edge with a given conditions (other predicates).
func HasAttestationCollectionsWith(preds ...predicate.AttestationCollection) predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := newAttestationCollectionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttributesReport applies the HasEdge predicate on the "attributes_report" edge.
func HasAttributesReport() predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AttributesReportTable, AttributesReportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttributesReportWith applies the HasEdge predicate on the "attributes_report" edge with a given conditions (other predicates).
func HasAttributesReportWith(preds ...predicate.AttributeReport) predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := newAttributesReportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDsse applies the HasEdge predicate on the "dsse" edge.
func HasDsse() predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DsseTable, DsseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDsseWith applies the HasEdge predicate on the "dsse" edge with a given conditions (other predicates).
func HasDsseWith(preds ...predicate.Dsse) predicate.Statement {
	return predicate.Statement(func(s *sql.Selector) {
		step := newDsseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.NotPredicates(p))
}
