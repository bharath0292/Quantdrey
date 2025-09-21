package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Uint64 = uint64

// MarshalUint64 serializes a uint64 as a string
func MarshalUint64(u uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatUint(u, 10))
	})
}

// UnmarshalUint64 deserializes an input value to uint64, expects string or int input
func UnmarshalUint64(v any) (uint64, error) {
	switch v := v.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("Uint64 cannot be negative")
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("Uint64 cannot be negative")
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("Uint64 cannot be negative")
		}
		return uint64(v), nil
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse Uint64 from string: %v", err)
		}
		return u, nil
	default:
		return 0, fmt.Errorf("unexpected type for Uint64: %T", v)
	}
}
