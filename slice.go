package expect

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// SliceContaining expects got to be a slice of element type T contain all values given as wants in any order.
// Duplicates in wants are not considered to be contained multiple times in the given slice.
func SliceContaining[T comparable](wants ...T) Matcher {
	return MatcherFunc(func(ctx Context, got any) {
		s, ok := got.([]T)
		if !ok {
			ctx.Failf("expected value of type %T but got %T", s, got)
			return
		}

		if len(wants) == 0 {
			return
		}

		wantsMissing := mapset.NewSet(wants...)

		for _, g := range s {
			wantsMissing.Remove(g)
			if wantsMissing.Cardinality() == 0 {
				return
			}
		}

		ctx.Failf("%T does not contain %v", got, wantsMissing.ToSlice())
	})
}

// SliceContainingInOrder expects go to be a slice with element type T containing all values given as wants
// in the same order they are given as wants.
func SliceContainingInOrder[T comparable](wants ...T) Matcher {
	return MatcherFunc(func(ctx Context, got any) {
		s, ok := got.([]T)
		if !ok {
			ctx.Failf("expected value of type %T but got %T", s, got)
			return
		}

		if len(wants) == 0 {
			return
		}

		for _, g := range s {
			if g == wants[0] {
				wants = wants[1:]
				if len(wants) == 0 {
					return
				}
			}
		}

		ctx.Failf("%T does not contain %v in order", got, wants[0])
	})
}