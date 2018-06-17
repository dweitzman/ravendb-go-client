package ravendb

type GroupByCountToken struct {
	*QueryToken

	_fieldName string
}

func NewGroupByCountToken(fieldName string) *GroupByCountToken {
	return &GroupByCountToken{
		QueryToken: NewQueryToken(),

		_fieldName: fieldName,
	}
}

func GroupByCountToken_create(fieldName string) *GroupByCountToken {
	return NewGroupByCountToken(fieldName)
}

func (t *GroupByCountToken) writeTo(writer *StringBuilder) {
	writer.append("count()")

	if t._fieldName == "" {
		return
	}

	writer.append(" as ")
	writer.append(t._fieldName)
}
