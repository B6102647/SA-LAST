// Code generated by entc, DO NOT EDIT.

package bookborrow

import (
	"time"

	"github.com/B6102647/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BOOKBORROWID applies equality check predicate on the "BOOKBORROW_ID" field. It's identical to BOOKBORROWIDEQ.
func BOOKBORROWID(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKBORROWID), v))
	})
}

// ADDEDTIME applies equality check predicate on the "ADDED_TIME" field. It's identical to ADDEDTIMEEQ.
func ADDEDTIME(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDEDTIME), v))
	})
}

// BOOKBORROWIDEQ applies the EQ predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDEQ(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKBORROWID), v))
	})
}

// BOOKBORROWIDNEQ applies the NEQ predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDNEQ(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBOOKBORROWID), v))
	})
}

// BOOKBORROWIDIn applies the In predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDIn(vs ...int) predicate.BookBorrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBOOKBORROWID), v...))
	})
}

// BOOKBORROWIDNotIn applies the NotIn predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDNotIn(vs ...int) predicate.BookBorrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBOOKBORROWID), v...))
	})
}

// BOOKBORROWIDGT applies the GT predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDGT(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBOOKBORROWID), v))
	})
}

// BOOKBORROWIDGTE applies the GTE predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDGTE(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBOOKBORROWID), v))
	})
}

// BOOKBORROWIDLT applies the LT predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDLT(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBOOKBORROWID), v))
	})
}

// BOOKBORROWIDLTE applies the LTE predicate on the "BOOKBORROW_ID" field.
func BOOKBORROWIDLTE(v int) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBOOKBORROWID), v))
	})
}

// ADDEDTIMEEQ applies the EQ predicate on the "ADDED_TIME" field.
func ADDEDTIMEEQ(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMENEQ applies the NEQ predicate on the "ADDED_TIME" field.
func ADDEDTIMENEQ(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMEIn applies the In predicate on the "ADDED_TIME" field.
func ADDEDTIMEIn(vs ...time.Time) predicate.BookBorrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldADDEDTIME), v...))
	})
}

// ADDEDTIMENotIn applies the NotIn predicate on the "ADDED_TIME" field.
func ADDEDTIMENotIn(vs ...time.Time) predicate.BookBorrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BookBorrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldADDEDTIME), v...))
	})
}

// ADDEDTIMEGT applies the GT predicate on the "ADDED_TIME" field.
func ADDEDTIMEGT(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMEGTE applies the GTE predicate on the "ADDED_TIME" field.
func ADDEDTIMEGTE(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMELT applies the LT predicate on the "ADDED_TIME" field.
func ADDEDTIMELT(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMELTE applies the LTE predicate on the "ADDED_TIME" field.
func ADDEDTIMELTE(v time.Time) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldADDEDTIME), v))
	})
}

// HasOwner applies the HasEdge predicate on the "Owner" edge.
func HasOwner() predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "Owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBOOK applies the HasEdge predicate on the "BOOK" edge.
func HasBOOK() predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BOOKTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BOOKTable, BOOKColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBOOKWith applies the HasEdge predicate on the "BOOK" edge with a given conditions (other predicates).
func HasBOOKWith(preds ...predicate.Book) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BOOKInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BOOKTable, BOOKColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPURPOSE applies the HasEdge predicate on the "PURPOSE" edge.
func HasPURPOSE() predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PURPOSETable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PURPOSETable, PURPOSEColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPURPOSEWith applies the HasEdge predicate on the "PURPOSE" edge with a given conditions (other predicates).
func HasPURPOSEWith(preds ...predicate.Purpose) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PURPOSEInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PURPOSETable, PURPOSEColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.BookBorrow) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.BookBorrow) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
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
func Not(p predicate.BookBorrow) predicate.BookBorrow {
	return predicate.BookBorrow(func(s *sql.Selector) {
		p(s.Not())
	})
}
