package header

import (
	"github.com/narqo/benchserder/internal/nullable"
)

//proteus:generate
type Data struct {
	// header data might be invalid if we parsed some headers and recognized
	// bots or other known malformed user agents. this flag is used to redirect
	// all clicks
	valid bool `json:"-"` // invalid for bots and unknowns, used for clicks

	Country         string
	ClientSdk       string `json:"ClientSdk"`
	AppName         string
	AppVersion      string
	AppVersionShort string
	Region          string
	DeviceName      string
	DeviceType      string
	Referer         string
	OsName          string
	OsVersion       string
	ApiLevel        string
	Language        string
	DisplaySize     string
	DisplayFormat   string
	DisplayDensity  string
	DisplayWidth    string
	DisplayHeight   string
	Author          string
	Publisher       string
	Manufacturer    string
	Architecture    string
	NetworkType     string // WIFI, MOBILE
	NetworkSubtype  string // EDGE, HSDPA, etc.
	Mcc             string // Mobile Country Code
	Mnc             string // Mobile Network Code

	IpAddress       string // for country lookup
	RemoteIpAddress string // an actual IP address the request came from, ignores ip_address= param

	NullSdkLevel nullable.Int `json:"NullSdkLevel"` // derived by SdkVersion
	UserAgentErr error
}

func (data *Data) IsValid() bool {
	return data.valid
}
