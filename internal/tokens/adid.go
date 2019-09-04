package tokens

//proteus:generate
type Adid string

type AdidMap map[AppToken]Adid // appToken -> adid

func (adid Adid) IsEmpty() bool {
	return adid == ""
}

func (adid Adid) String() string {
	return string(adid)
}

func (adidMap AdidMap) ToStringMap() map[string]string {
	if adidMap == nil {
		return nil
	}

	stringMap := make(map[string]string, len(adidMap))
	for appToken, adid := range adidMap {
		stringMap[appToken.String()] = adid.String()
	}

	return stringMap
}
