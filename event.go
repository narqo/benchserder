package benchserder

import (
	"time"

	"github.com/narqo/benchserder/internal/callback"
	"github.com/narqo/benchserder/internal/fraud"
	"github.com/narqo/benchserder/internal/money"
	"github.com/narqo/benchserder/internal/nullable"
)

//easyjson:json
//proteus:generate
type Event struct {
	AppToken string `thrift:",1"`
	Tracker  string `thrift:",2"`
	Adid     string `thrift:",3"`

	FacebookAttributionId string `thrift:",4"`
	FacebookAnonId        string `thrift:",5"`
	TrackingEnabled       string `thrift:",6"`

	ClickTime        time.Time `thrift:",7"`
	FirstSessionTime time.Time `thrift:",8"`
	LastSessionTime  time.Time `thrift:",9"`
	LastEventTime    time.Time `thrift:",10"`
	LastRevenueTime  time.Time `thrift:",11"`
	CreatedAt        time.Time `thrift:",12"`
	ReceivedAt       time.Time `thrift:",13"`

	InstallTime            time.Time     `json:",omitempty" thrift:",14"`
	InstallTracker         string        `json:",omitempty" thrift:",15"`
	InstallCountry         string        `json:",omitempty" thrift:",16"`
	InstallImpressionBased nullable.Bool `json:",omitempty" thrift:",17"`

	EventToken string `thrift:",18"`

	RevenueData money.Amount `json:"RevenueData" thrift:",19"`

	DeviceType  string `thrift:",20"`
	Environment string `thrift:",21"`

	NullSdkLevel nullable.Int `json:"NullSdkLevel" thrift:",22"`
	ZoneOffset   int          `thrift:",23"`

	// for attribution
	FraudKind fraud.Kind `json:"FraudKind" thrift:",24"`

	// for callbacks
	PingbackUrl  string         `thrift:",25"`
	CallbackData *callback.Data `thrift:",26"`

	// for insert
	FirstOsName        string `thrift:",27"`
	FirstCountry       string `thrift:",28"`
	FirstDeviceType    string `thrift:",29"`
	ImpressionBased    bool   `thrift:",30"`
	DeviceReattributed bool   `thrift:",31"`
}
