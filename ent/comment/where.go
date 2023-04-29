// Code generated by ent, DO NOT EDIT.

package comment

import (
	"test/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldComment, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUserID, vs...))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldComment, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		p(s.Not())
	})
}
