package tokens

//proteus:generate
type AppToken string

func (appToken AppToken) IsEmpty() bool {
	return appToken == ""
}

func (appToken AppToken) String() string {
	return string(appToken)
}

func (appToken AppToken) IsValid() bool {
	return len(appToken) == 12
}
