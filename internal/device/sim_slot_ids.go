package device

//easyjson:json
type SimSlotIds struct {
	Imeis    []string `thrift:",1"`
	Meids    []string `thrift:",2"`
	DeviceId string   `thrift:",3"`
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (ids SimSlotIds) IsDefined() bool {
	return ids.DeviceId != "" || len(ids.Imeis) > 0 || len(ids.Meids) > 0
}
