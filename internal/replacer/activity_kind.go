package replacer

type ActivityKind string

const (
	AcRedirect               ActivityKind = "redirect"
	AcSdkClick               ActivityKind = "sdk_click"
	AcSdkInfo                ActivityKind = "sdk_info"
	AcClick                  ActivityKind = "click"
	AcSanClick               ActivityKind = "san_click"
	AcImpression             ActivityKind = "impression"
	AcSanImpression          ActivityKind = "san_impression"
	AcInstall                ActivityKind = "install"
	AcRejectedInstall        ActivityKind = "rejected_install"
	AcSession                ActivityKind = "session"
	AcEvent                  ActivityKind = "event"
	AcReattribution          ActivityKind = "reattribution"
	AcRejectedReattribution  ActivityKind = "rejected_reattribution"
	AcAttributionUpdate      ActivityKind = "update"
	AcUninstall              ActivityKind = "uninstall"
	AcReinstall              ActivityKind = "reinstall"
	AcReattributionReinstall ActivityKind = "reattribution_reinstall"
	AcOptOut                 ActivityKind = "optout"
	AcGdprForgetDevice       ActivityKind = "gdpr_forget_device"
	AcAdRevenue              ActivityKind = "ad_revenue"
	AcCostUpdate             ActivityKind = "cost_update"
)

var ActivityKinds = map[ActivityKind]struct{}{
	AcRedirect:               {},
	AcSdkInfo:                {},
	AcClick:                  {},
	AcSanClick:               {},
	AcImpression:             {},
	AcSanImpression:          {},
	AcInstall:                {},
	AcRejectedInstall:        {},
	AcSession:                {},
	AcEvent:                  {},
	AcReattribution:          {},
	AcRejectedReattribution:  {},
	AcAttributionUpdate:      {},
	AcUninstall:              {},
	AcReinstall:              {},
	AcReattributionReinstall: {},
	AcOptOut:                 {},
	AcGdprForgetDevice:       {},
	AcAdRevenue:              {},
	AcCostUpdate:             {},
}

func (activityKind ActivityKind) IsReinstall() bool {
	if activityKind == AcReinstall || activityKind == AcReattributionReinstall {
		return true
	}

	return false
}
