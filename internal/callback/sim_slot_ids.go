package callback

type SimSlotIds struct {
	Imeis    []string
	Meids    []string
	DeviceId string
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (ids SimSlotIds) IsDefined() bool {
	return ids.DeviceId != "" || len(ids.Imeis) > 0 || len(ids.Meids) > 0
}
