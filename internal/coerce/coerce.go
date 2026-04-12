package coerce

import (
	"fmt"
	"strconv"
)

// Int converts an any-typed API field (float64, int, string, or nil) to int.
func Int(v any) (int, error) {
	switch x := v.(type) {
	case float64:
		return int(x), nil
	case int:
		return x, nil
	case string:
		return strconv.Atoi(x)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("coerce: cannot convert %T to int", v)
	}
}

// Float64 converts an any-typed API field to float64.
func Float64(v any) (float64, error) {
	switch x := v.(type) {
	case float64:
		return x, nil
	case int:
		return float64(x), nil
	case string:
		return strconv.ParseFloat(x, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("coerce: cannot convert %T to float64", v)
	}
}
