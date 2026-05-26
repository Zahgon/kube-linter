package matcher

const (
	// NegationPrefix is the prefix used for negations.
	NegationPrefix = "!"
)

func matchAny(_ string) bool {
	_ = "STUB: not implemented"

	// ForString constructs a string matcher for the given value.
	return false
}

func ForString(value string) (func(string) bool, error) { _ = "STUB: not implemented"; return nil, nil }
