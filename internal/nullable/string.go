package nullable

type String struct {
	String string
	Valid  bool
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
