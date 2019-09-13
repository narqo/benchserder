package nullable

import "database/sql"

type Bool struct {
	Bool  bool `thrift:",1"`
	Valid bool `thrift:",2"`
}

var (
	True  = NewBool(true)
	False = NewBool(false)
)

func NewBool(value bool) Bool {
	return Bool{
		Bool:  value,
		Valid: true,
	}
}

func (v Bool) BoolPtr() *bool {
	if v.Valid {
		return &v.Bool
	}

	return nil
}

func (v Bool) SqlNullBool() sql.NullBool {
	return sql.NullBool{
		Bool:  v.Bool,
		Valid: v.Valid,
	}
}

func (v Bool) IsDefined() bool {
	return v.Valid
}

func (v Bool) IsTrue() bool {
	return v.Valid && v.Bool
}

func (v Bool) String() string {
	if !v.Valid {
		return "invalid"
	}
	if v.Bool {
		return "true"
	}
	return "false"
}
