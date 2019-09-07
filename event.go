package benchserder

import (
	"time"

	"github.com/narqo/benchserder/internal/callback"
	"github.com/narqo/benchserder/internal/fraud"
	"github.com/narqo/benchserder/internal/money"
	"github.com/narqo/benchserder/internal/nullable"
	"github.com/narqo/benchserder/internal/tokens"
)

//easyjson:json
//proteus:generate
type Event struct {
	AppToken tokens.AppToken
	Tracker  tokens.TrackerToken
	Adid     tokens.Adid

	FacebookAttributionId string
	FacebookAnonId        string
	TrackingEnabled       string

	ClickTime        time.Time
	FirstSessionTime time.Time
	LastSessionTime  time.Time
	LastEventTime    time.Time
	LastRevenueTime  time.Time
	CreatedAt        time.Time
	ReceivedAt       time.Time

	InstallTime            time.Time           `json:",omitempty"`
	InstallTracker         tokens.TrackerToken `json:",omitempty"`
	InstallCountry         string              `json:",omitempty"`
	InstallImpressionBased nullable.Bool       `json:",omitempty"`

	EventToken string

	Revenue money.Amount `json:"RevenueData"`

	DeviceType  string
	Environment string

	NullSdkLevel nullable.Int `json:"NullSdkLevel"`
	ZoneOffset   int

	// for attribution
	FraudKind fraud.Kind `json:"FraudKind"`

	// for callbacks
	PingbackUrl  string
	CallbackData *callback.Data

	// for insert
	FirstOsName        string
	FirstCountry       string
	FirstDeviceType    string
	ImpressionBased    bool
	DeviceReattributed bool

	//headerData      *header.Data
	//eventCount      int
	//lastEventCount  int
	//sessionCount    int
	//subsessionCount int
	//sessionLength   int
	//timeSpent       int

	XXX_unrecognized []byte `json:"-"`
}
