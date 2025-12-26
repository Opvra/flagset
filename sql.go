package flagset

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strconv"
)

// Scan implements the sql.Scanner interface.
func (f *Flag) Scan(value any) error {
	if value == nil {
		*f = 0
		return nil
	}

	switch v := value.(type) {
	case int64:
		if v < 0 {
			return fmt.Errorf("flagset: negative value %d", v)
		}
		*f = Flag(v)
		return nil
	case int32:
		if v < 0 {
			return fmt.Errorf("flagset: negative value %d", v)
		}
		*f = Flag(v)
		return nil
	case int:
		if v < 0 {
			return fmt.Errorf("flagset: negative value %d", v)
		}
		*f = Flag(v)
		return nil
	case int16:
		if v < 0 {
			return fmt.Errorf("flagset: negative value %d", v)
		}
		*f = Flag(v)
		return nil
	case int8:
		if v < 0 {
			return fmt.Errorf("flagset: negative value %d", v)
		}
		*f = Flag(v)
		return nil
	case uint64:
		*f = Flag(v)
		return nil
	case uint32:
		*f = Flag(v)
		return nil
	case uint:
		*f = Flag(v)
		return nil
	case uint16:
		*f = Flag(v)
		return nil
	case uint8:
		*f = Flag(v)
		return nil
	case []byte:
		return f.scanString(string(v))
	case string:
		return f.scanString(v)
	default:
		return fmt.Errorf("flagset: unsupported Scan type %T", value)
	}
}

func (f *Flag) scanString(value string) error {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	*f = Flag(parsed)
	return nil
}

// Value implements the driver.Valuer interface.
func (f Flag) Value() (driver.Value, error) {
	if f > Flag(math.MaxInt64) {
		return nil, fmt.Errorf("flagset: value %d overflows int64", f)
	}
	return int64(f), nil
}
