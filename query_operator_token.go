package ravendb

type QueryOperatorToken struct {
	*QueryToken

	_queryOperator QueryOperator
}

var (
	QueryOperatorToken_AND = NewQueryOperatorToken(QueryOperator_AND)
	QueryOperatorToken_OR  = NewQueryOperatorToken(QueryOperator_OR)
)

func NewQueryOperatorToken(queryOperator QueryOperator) *QueryOperatorToken {
	return &QueryOperatorToken{
		QueryToken: NewQueryToken(),

		_queryOperator: queryOperator,
	}
}

func (t *QueryOperatorToken) writeTo(writer *StringBuilder) {
	if t._queryOperator == QueryOperator_AND {
		writer.append("and")
		return
	}

	writer.append("or")
}
