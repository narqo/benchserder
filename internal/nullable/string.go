package nullable

type String struct {
	String string  `thrift:",1"`
	Valid  bool    `thrift:",2"`
}

func NewString(s string) String {
	return String{
		String: s,
		Valid:  true,
	}
}

func (s String) IsDefined() bool {
	return s.Valid
}
