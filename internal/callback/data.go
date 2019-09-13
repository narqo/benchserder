package callback

import (
	"time"

	"github.com/narqo/benchserder/internal/device"
	"github.com/narqo/benchserder/internal/engagementtype"
	"github.com/narqo/benchserder/internal/fraud"
	"github.com/narqo/benchserder/internal/money"
	"github.com/narqo/benchserder/internal/nullable"
)

//easyjson:json
type Data struct {
	TermsSigned        bool                `json:",omitempty" thrift:",1"`
	ReferenceTag       string              `json:"ReferenceTag,omitempty" thrift:",2"`
	EngagementType     engagementtype.Type `json:",omitempty" thrift:",3"`
	Tracker            string              `json:",omitempty" thrift:",4"` // required
	InstallTracker     string              `json:",omitempty" thrift:",5"`
	LastTracker        string              `json:",omitempty" thrift:",6"`
	OutdatedTracker    string              `json:",omitempty" thrift:",7"`
	AppToken           string              `json:",omitempty" thrift:",8"`
	Adid               string              `json:",omitempty" thrift:",9"`
	OsName             string              `json:",omitempty" thrift:",10"`
	OsVersion          string              `json:",omitempty" thrift:",11"`
	OsBuild            string              `json:",omitempty" thrift:",12"`
	ApiLevel           string              `json:",omitempty" thrift:",13"`
	Nonce              string              `json:",omitempty" thrift:",14"`
	Random             string              `json:",omitempty" thrift:",15"`
	AppName            string              `json:",omitempty" thrift:",16"`
	AppVersion         string              `json:",omitempty" thrift:",17"`
	AppVersionShort    string              `json:",omitempty" thrift:",18"`
	Idfa               string              `json:",omitempty" thrift:",19"`
	Idfv               string              `json:",omitempty" thrift:",20"`
	AndroidId          string              `json:",omitempty" thrift:",21"`
	GoogleAdid         string              `json:"GoogleAdid,omitempty" thrift:",22"`
	FireAdid           string              `json:",omitempty" thrift:",23"`
	WebUuid            string              `json:",omitempty" thrift:",24"`
	WinUdid            string              `json:",omitempty" thrift:",25"`
	WinHwid            string              `json:",omitempty" thrift:",26"`
	WinNaid            string              `json:",omitempty" thrift:",27"`
	WinAdid            string              `json:",omitempty" thrift:",28"`
	SimSlotIds         device.SimSlotIds   `json:",omitempty" thrift:",29"`
	EventToken         string              `json:",omitempty" thrift:",30"`
	RevenueData        money.Amount        `json:"RevenueData,omitempty" thrift:",31"`
	Cost               money.Cost          `json:",omitempty" thrift:",32"`
	CostIdMd5          string              `json:",omitempty" thrift:",33"`
	IpAddress          string              `json:",omitempty" thrift:",34"`
	ProxyIpAddress     string              `json:",omitempty" thrift:",35"`
	ServerIp           string              `json:",omitempty" thrift:",36"` // IP address of the remote in case of s2s click
	MacSha1            string              `json:",omitempty" thrift:",37"`
	MacMd5             string              `json:",omitempty" thrift:",38"`
	Environment        string              `json:",omitempty" thrift:",39"`
	Referrer           string              `json:",omitempty" thrift:",40"`
	Country            string              `json:",omitempty" thrift:",41"`
	Language           string              `json:",omitempty" thrift:",42"`
	Region             string              `json:",omitempty" thrift:",43"`
	Mcc                string              `json:",omitempty" thrift:",44"`
	Mnc                string              `json:",omitempty" thrift:",45"`
	Manufacturer       string              `json:",omitempty" thrift:",46"`
	ClientSdk          string              `json:"ClientSdk,omitempty" thrift:",47"`
	StoreName          string              `json:",omitempty" thrift:",48"`
	StoreAppId         string              `json:",omitempty" thrift:",49"`
	SessionCount       string              `json:",omitempty" thrift:",50"`
	GlobalSessionCount int                 `json:",omitempty" thrift:",51"`
	PushToken          string              `json:",omitempty" thrift:",52"`
	TimeSpent          string              `json:",omitempty" thrift:",53"`
	LastTimeSpent      string              `json:",omitempty" thrift:",54"`
	DeviceName         string              `json:",omitempty" thrift:",55"`
	DeviceType         string              `json:",omitempty" thrift:",56"`
	TrackingEnabled    string              `json:",omitempty" thrift:",57"`
	ZoneOffset         nullable.Int        `json:",omitempty" thrift:",58"`
	ClickLabel         string              `json:",omitempty" thrift:",59"`
	UserAgent          string              `json:",omitempty" thrift:",60"`
	//ActivityKind       replacer.ActivityKind `json:",omitempty"` // protobuf can't do string enums
	ActivityKind    string       `json:",omitempty" thrift:",61"`
	Deeplink        string       `json:",omitempty" thrift:",62"`
	SearchTerm      string       `json:",omitempty" thrift:",63"`
	ClickReferer    string       `json:",omitempty" thrift:",64"`
	ImpressionBased bool         `json:",omitempty" thrift:",65"`
	FraudKind       fraud.Kind   `json:",omitempty" thrift:",66"`
	CpuType         string       `json:",omitempty" thrift:",67"`
	HardwareName    string       `json:",omitempty" thrift:",68"`
	NetworkType     string       `json:",omitempty" thrift:",69"`
	SecretId        nullable.Int `json:",omitempty" thrift:",70"`
	IsImported      bool         `json:",omitempty" thrift:",71"`

	IsServerToServer      bool `json:",omitempty" thrift:",72"`
	IsServerToServerBased bool `json:",omitempty" thrift:",73"`

	// TODO: more consistent naming?
	ClickTime                      time.Time         `json:"ClickTime,omitempty" thrift:",74"`
	InstalledAt                    time.Time         `json:"InstalledAt,omitempty" thrift:",75"`
	FirstSessionTime               time.Time         `json:",omitempty" thrift:",76"`
	FirstCountry                   string            `json:",omitempty" thrift:",77"`
	FirstDeviceType                string            `json:",omitempty" thrift:",78"`
	FirstOsName                    string            `json:",omitempty" thrift:",79"`
	ReceivedAt                     time.Time         `json:",omitempty" thrift:",80"`
	CreatedAtTime                  time.Time         `json:",omitempty" thrift:",81"`
	LastSessionTime                time.Time         `json:",omitempty" thrift:",82"`
	UninstalledAt                  time.Time         `json:",omitempty" thrift:",83"` // current or last uninstall time
	ReinstalledAt                  time.Time         `json:",omitempty" thrift:",84"` // current or last reinstall time
	AttributionUpdatedAt           time.Time         `json:",omitempty" thrift:",85"`
	SessionLength                  int               `json:",omitempty" thrift:",86"`
	PingbackTtl                    nullable.Duration `json:",omitempty" thrift:",87"`
	ClickAttributionWindow         nullable.Duration `json:",omitempty" thrift:",88"`
	ImpressionAttributionWindow    nullable.Duration `json:",omitempty" thrift:",89"`
	FingerprintAttributionWindow   nullable.Duration `json:",omitempty" thrift:",90"`
	ReattributionAttributionWindow nullable.Duration `json:",omitempty" thrift:",91"`
	InactiveUserDefinition         nullable.Duration `json:",omitempty" thrift:",92"`

	// Make sure to update Copy() when adding new maps!
	PartnerParameters         map[string]string `json:"PartnerParameters,omitempty" thrift:",93"`
	PartnerSdkParams          map[string]string `json:"PartnerSdkParams,omitempty" thrift:",94"`
	PublisherParams           map[string]string `json:",omitempty" thrift:",95"`
	DynamicCallbackParameters map[string]string `json:"DynamicCallbackParameters,omitempty" thrift:",95"`
	ApiPartnerParams          map[string]string `json:",omitempty" thrift:",96"`

	// for ad revenue update
	AdImpressionsCount  int    `json:",omitempty" thrift:",97"`
	AdRevenueNetwork    string `json:",omitempty" thrift:",98"`
	AdRevenueUnit       string `json:",omitempty" thrift:",99"`
	AdRevenuePlacement  string `json:",omitempty" thrift:",100"`
	AdMediationPlatform string `json:",omitempty" thrift:",101"`
}
