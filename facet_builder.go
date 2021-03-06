package ravendb

import "strings"

func isRqlKeyword(s string) bool {
	s = strings.ToLower(s)
	switch s {
	case "as", "select", "where", "load", "group", "order", "include", "update":
		return true
	}
	return false
}

type FacetBuilder struct {
	_range   *GenericRangeFacet
	_default *Facet
}

func NewFacetBuilder() *FacetBuilder {
	return &FacetBuilder{}
}

func (b *FacetBuilder) ByRanges(rng *RangeBuilder, ranges ...*RangeBuilder) *FacetBuilder {
	if rng == nil {
		//throw new IllegalArgumentError("Range cannot be null")
		panic("Range cannot be null")
	}

	if b._range == nil {
		b._range = NewGenericRangeFacet(nil)
	}

	b._range.addRange(rng)

	for _, rng := range ranges {
		b._range.addRange(rng)
	}

	return b
}

func (b *FacetBuilder) ByField(fieldName string) *FacetBuilder {
	if b._default == nil {
		b._default = NewFacet()
	}

	if isRqlKeyword(fieldName) {
		fieldName = "'" + fieldName + "'"
	}

	b._default.FieldName = fieldName

	return b
}

func (b *FacetBuilder) AllResults() *FacetBuilder {
	if b._default == nil {
		b._default = NewFacet()
	}

	b._default.FieldName = ""
	return b
}

func (b *FacetBuilder) WithOptions(options *FacetOptions) *FacetBuilder {
	b.GetFacet().SetOptions(options)
	return b
}

func (b *FacetBuilder) WithDisplayName(displayName string) *FacetBuilder {
	b.GetFacet().SetDisplayFieldName(displayName)
	return b
}

func (b *FacetBuilder) SumOn(path string) *FacetBuilder {
	b.GetFacet().GetAggregations()[FacetAggregationSum] = path
	return b
}

func (b *FacetBuilder) MinOn(path string) *FacetBuilder {
	b.GetFacet().GetAggregations()[FacetAggregationMin] = path
	return b
}

func (b *FacetBuilder) MaxOn(path string) *FacetBuilder {
	b.GetFacet().GetAggregations()[FacetAggregationMax] = path
	return b
}

func (b *FacetBuilder) AverageOn(path string) *FacetBuilder {
	b.GetFacet().GetAggregations()[FacetAggregationAverage] = path
	return b
}

func (b *FacetBuilder) GetFacet() FacetBase {
	if b._default != nil {
		return b._default
	}

	return b._range
}
