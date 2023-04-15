package entity

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"errors"
	"strconv"
)

var (
	ErrOverflow     = errors.New("Overflow")
	ErrInvalidValue = errors.New("Invalid value")
)

// Flags binary hold the base2(binary) format in raw byte array [0,1].
type Flags []byte

// FlagIndex to differentiate with another int like type.
type FlagIndex int

// String implement a Stringer.
func (f Flags) String() string { return f.Format("0", "_", 8) }

// Format will pretty print a base2(binary) format.
func (f Flags) Format(pad, sep string, n int) (o string) {
	s := bytes.TrimLeft(f, pad)
	m := len(s) % n
	k, r := true, m > 0

	for len(s) > 0 {
		if r {
			o += sep
			if k {
				o = string(bytes.Repeat([]byte(pad), n-m))
			}
		}

		if !k {
			m = n
		}

		o += string(s[:m])
		s = s[m:]
		k, r = false, true
	}

	return o
}

// Valid will check for 0 or 1 byte in each of it's element.
func (f Flags) Valid() bool {
	for i := range f {
		switch f[i] {
		case '0', '1':
			continue
		default:
			return false
		}
	}

	return len(f) > 0
}

// Is compare with other Flags with exact equal.
func (f Flags) Is(g Flags) bool {
	if len(f) != len(g) {
		f = Flags(bytes.TrimLeft(f, "0"))
		g = Flags(bytes.TrimLeft(g, "0"))
	}

	return (len(f) < 1 && len(g) < 1) ||
		(bytes.Equal(f, g) && g.Valid())
}

func (f Flags) Has(g Flags) bool {
	if len(g) > 0 && g[0] == '0' {
		g = Flags(bytes.TrimLeft(g, "0"))
	}

	for i := len(g) - 1; i >= 0; i-- {
		if g[i] != '1' {
			continue
		} else if v, _ := f.Get(FlagIndex(len(g) - i)); !v {
			return false
		}
	}

	return len(g) > 0
}

// Get bit value on the given index, start from 1 instead of 0 (right to left).
func (f Flags) Get(i FlagIndex) (v bool, err error) {
	if i < 1 || int(i) > len(f) {
		return false, ErrOverflow
	} else if !f.Valid() {
		return false, ErrInvalidValue
	}

	bitIndex := len(f) - int(i)

	switch f[bitIndex] {
	case '1':
		return true, err
	default:
		return false, err
	}
}

// Set bit value on the given index, start from 1 instead of 0 (right to left).
func (f Flags) Set(i FlagIndex, v bool) (g Flags, err error) {
	if i < 1 {
		return nil, ErrOverflow
	} else if !f.Valid() {
		return nil, ErrInvalidValue
	}

	s := 0
	if g = make(Flags, len(f)); len(f) < int(i) {
		g, s = make(Flags, i), int(i)-len(f)
		for i := 0; i < s; i++ {
			g[i] = '0'
		}
	}

	copy(g[s:], f)

	bitIndex := len(g) - int(i)

	switch {
	case v:
		g[bitIndex] = '1'
	default:
		g[bitIndex] = '0'
	}

	return g, err
}

// Scan is sql.Scanner implementation.
func (f *Flags) Scan(src interface{}) (err error) {
	_ = sql.Scanner(f)

	*f = Flags(nil)

	switch src := src.(type) {
	case int64:
		*f = Flags(strconv.FormatUint(uint64(src), 2))
	case float64:
		*f = Flags(strconv.FormatUint(uint64(src), 2))
	case bool:
		if src {
			*f = Flags("1")
		} else {
			*f = Flags("0")
		}
	case []byte:
		*f = Flags(string(src))
	case string:
		*f = Flags(src)
	}

	if !f.Valid() {
		return ErrInvalidValue
	}

	return err
}

// Value is driver.Valuer implementation.
func (f Flags) Value() (v driver.Value, err error) {
	_ = driver.Valuer(f)

	if !f.Valid() {
		return nil, ErrInvalidValue
	}

	return string(f), err
}
