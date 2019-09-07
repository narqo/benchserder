package device

//easyjson:json
type SimSlotIds struct {
	Imeis    []string
	Meids    []string
	DeviceId string

	XXX_unrecognized []byte `json:"-"`
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (ids SimSlotIds) IsDefined() bool {
	return ids.DeviceId != "" || len(ids.Imeis) > 0 || len(ids.Meids) > 0
}
