package callback

import (
	"time"

	"github.com/narqo/benchserder/internal/device"
	"github.com/narqo/benchserder/internal/engagementtype"
	"github.com/narqo/benchserder/internal/fraud"
	"github.com/narqo/benchserder/internal/money"
	"github.com/narqo/benchserder/internal/nullable"
	"github.com/narqo/benchserder/internal/replacer"
	"github.com/narqo/benchserder/internal/tokens"
)

//easyjson:json
type Data struct {
	TermsSigned        bool                  `json:",omitempty"`
	ReferenceTag       string                `json:"ReferenceTag,omitempty"`
	EngagementType     engagementtype.Type   `json:",omitempty"`
	Tracker            tokens.TrackerToken   `json:",omitempty"` // required
	InstallTracker     tokens.TrackerToken   `json:",omitempty"`
	LastTracker        tokens.TrackerToken   `json:",omitempty"`
	OutdatedTracker    tokens.TrackerToken   `json:",omitempty"`
	AppToken           tokens.AppToken       `json:",omitempty"`
	Adid               tokens.Adid           `json:",omitempty"`
	OsName             string                `json:",omitempty"`
	OsVersion          string                `json:",omitempty"`
	OsBuild            string                `json:",omitempty"`
	ApiLevel           string                `json:",omitempty"`
	Nonce              string                `json:",omitempty"`
	Random             string                `json:",omitempty"`
	AppName            string                `json:",omitempty"`
	AppVersion         string                `json:",omitempty"`
	AppVersionShort    string                `json:",omitempty"`
	Idfa               string                `json:",omitempty"`
	Idfv               string                `json:",omitempty"`
	AndroidId          string                `json:",omitempty"`
	GoogleAdid         string                `json:"GoogleAdid,omitempty"`
	FireAdid           string                `json:",omitempty"`
	WebUuid            string                `json:",omitempty"`
	WinUdid            string                `json:",omitempty"`
	WinHwid            string                `json:",omitempty"`
	WinNaid            string                `json:",omitempty"`
	WinAdid            string                `json:",omitempty"`
	SimSlotIds         device.SimSlotIds     `json:",omitempty"`
	EventToken         string                `json:",omitempty"`
	RevenueData        money.Amount          `json:"RevenueData,omitempty"`
	Cost               money.Cost            `json:",omitempty"`
	CostIdMd5          string                `json:",omitempty"`
	IpAddress          string                `json:",omitempty"`
	ProxyIpAddress     string                `json:",omitempty"`
	ServerIp           string                `json:",omitempty"` // IP address of the remote in case of s2s click
	MacSha1            string                `json:",omitempty"`
	MacMd5             string                `json:",omitempty"`
	Environment        string                `json:",omitempty"`
	Referrer           string                `json:",omitempty"`
	Country            string                `json:",omitempty"`
	Language           string                `json:",omitempty"`
	Region             string                `json:",omitempty"`
	Mcc                string                `json:",omitempty"`
	Mnc                string                `json:",omitempty"`
	Manufacturer       string                `json:",omitempty"`
	ClientSdk          string                `json:"ClientSdk,omitempty"`
	StoreName          string                `json:",omitempty"`
	StoreAppId         string                `json:",omitempty"`
	SessionCount       string                `json:",omitempty"`
	GlobalSessionCount int                   `json:",omitempty"`
	PushToken          string                `json:",omitempty"`
	TimeSpent          string                `json:",omitempty"`
	LastTimeSpent      string                `json:",omitempty"`
	DeviceName         string                `json:",omitempty"`
	DeviceType         string                `json:",omitempty"`
	TrackingEnabled    string                `json:",omitempty"`
	ZoneOffset         nullable.Int          `json:",omitempty"`
	ClickLabel         string                `json:",omitempty"`
	UserAgent          string                `json:",omitempty"`
	ActivityKind       replacer.ActivityKind `json:",omitempty"`
	Deeplink           string                `json:",omitempty"`
	SearchTerm         string                `json:",omitempty"`
	ClickReferer       string                `json:",omitempty"`
	ImpressionBased    bool                  `json:",omitempty"`
	FraudKind          fraud.Kind            `json:",omitempty"`
	CpuType            string                `json:",omitempty"`
	HardwareName       string                `json:",omitempty"`
	NetworkType        string                `json:",omitempty"`
	SecretId           nullable.Int          `json:",omitempty"`
	IsImported         bool                  `json:",omitempty"`

	IsServerToServer      bool `json:",omitempty"`
	IsServerToServerBased bool `json:",omitempty"`

	// TODO: more consistent naming?
	ClickTime                      time.Time         `json:"ClickTime,omitempty"`
	InstalledAt                    time.Time         `json:"InstalledAt,omitempty"`
	FirstSessionTime               time.Time         `json:",omitempty"`
	FirstCountry                   string            `json:",omitempty"`
	FirstDeviceType                string            `json:",omitempty"`
	FirstOsName                    string            `json:",omitempty"`
	ReceivedAt                     time.Time         `json:",omitempty"`
	CreatedAtTime                  time.Time         `json:",omitempty"`
	LastSessionTime                time.Time         `json:",omitempty"`
	UninstalledAt                  time.Time         `json:",omitempty"` // current or last uninstall time
	ReinstalledAt                  time.Time         `json:",omitempty"` // current or last reinstall time
	AttributionUpdatedAt           time.Time         `json:",omitempty"`
	SessionLength                  int               `json:",omitempty"`
	PingbackTtl                    nullable.Duration `json:",omitempty"`
	ClickAttributionWindow         nullable.Duration `json:",omitempty"`
	ImpressionAttributionWindow    nullable.Duration `json:",omitempty"`
	FingerprintAttributionWindow   nullable.Duration `json:",omitempty"`
	ReattributionAttributionWindow nullable.Duration `json:",omitempty"`
	InactiveUserDefinition         nullable.Duration `json:",omitempty"`

	// Make sure to update Copy() when adding new maps!
	PartnerParameters         Parameters `json:"PartnerParameters,omitempty"`
	PartnerSdkParams          Parameters `json:"PartnerSdkParams,omitempty"`
	PublisherParams           Parameters `json:",omitempty"`
	DynamicCallbackParameters Parameters `json:"DynamicCallbackParameters,omitempty"`
	ApiPartnerParams          Parameters `json:",omitempty"`

	// for ad revenue update
	AdImpressionsCount  int    `json:",omitempty"`
	AdRevenueNetwork    string `json:",omitempty"`
	AdRevenueUnit       string `json:",omitempty"`
	AdRevenuePlacement  string `json:",omitempty"`
	AdMediationPlatform string `json:",omitempty"`
}
