package evmmock

const (
	ParamEqualToJson ParamMatchingStrategy = "equalToJson"
)

type ParamMatchingStrategy string

type ParamMatcher struct {
	strategy ParamMatchingStrategy
	value    string
}

func (m ParamMatcher) Strategy() ParamMatchingStrategy {
	return m.strategy
}

func (m ParamMatcher) Value() string {
	return m.value
}

func EqualToJson(param string) ParamMatcher {
	return ParamMatcher{
		strategy: ParamEqualToJson,
		value:    param,
	}
}
