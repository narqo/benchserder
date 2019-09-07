package nullable

import "strconv"

type Int struct {
	Int   int
	Valid bool

	XXX_unrecognized []byte `json:"-"`
}

func NewInt(value int) Int {
	return Int{
		Int:   value,
		Valid: true,
	}
}

// for legacy int which uses -1 as invalid
func LegacyInt(value int) Int {
	return Int{
		Int:   value,
		Valid: value != -1,
	}
}

func (v Int) IsDefined() bool {
	return v.Valid
}

func (v Int) String() string {
	if !v.Valid {
		return ""
	}
	return strconv.Itoa(v.Int)
}
