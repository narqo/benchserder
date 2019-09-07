package engagementtype

//proteus:generate
type Source int

const (
	FromUnknown                 Source = iota // unknown source
	FromIpList                                // found in IP list
	FromMappedDtag                            // found by mapping from device tag
	FromReferrer                              // found by reftag in session body (legacy)
	FromSdkReftag                             // found from sdk click reference tag
	FromSdkDeeplink                           // found from sdk click deeplink
	FromSdkIad                                // found from sdk click iad
	FromFacebook                              // found from facebook response
	FromGoogle                                // found from old google api response (legacy)
	FromAdwords                               // found from google adwords response
	FromTwitter                               // found from Twitter response
	FromDoubleClick                           // found from DoubleClick response
	FromYahooJapan                            // found from YahooJapan response
	FromSnapchat                              // found from Snapchat response
	FromExternalClickId                       // found from referrer sdk click with an external_click_id attached
	FromYahooGemini                           // found from YahooGemini response
	FromKing                                  // found from King response
	FromTencent                               // found from Tencent response
	FromGoogleMarketingPlatform               // found from google marketing platform response
	FromSdkIad3                               // found from sdk click iad3 (legacy)
	FromSdkAppleSearchAds                     // found from sdk click apple search ads
	FromSdkGoogleOrganicSearch                // found from sdk click - google organic search
	FromSdkAdjustTracker                      // found from sdk click - adjust tracker
)

func (source Source) Score() int {
	// for now we only prioritize SdkReftag, we may consider the order later
	switch source {
	case FromSdkReftag:
		return 3
	default:
		return 0
	}
}

func (source Source) String() string {
	switch source {
	case FromIpList:
		return "fromIpList"
	case FromMappedDtag:
		return "fromMappedDtag"
	case FromReferrer:
		return "fromReferrer"
	case FromSdkReftag:
		return "fromSdkReftag"
	case FromSdkGoogleOrganicSearch:
		return "fromSdkGoogleOrganicSearch"
	case FromSdkAdjustTracker:
		return "fromSdkAdjustTracker"
	case FromSdkDeeplink:
		return "fromSdkDeeplink"
	case FromSdkIad:
		return "fromSdkIad"
	case FromSdkIad3:
		return "fromSdkIad3"
	case FromSdkAppleSearchAds:
		return "fromSdkIad3SearchAds"
	case FromFacebook:
		return "fromFacebook"
	case FromGoogle:
		return "fromGoogle"
	case FromAdwords:
		return "fromAdwords"
	case FromTwitter:
		return "fromTwitter"
	case FromDoubleClick:
		return "fromDoubleClick"
	case FromYahooJapan:
		return "fromYahooJapan"
	case FromYahooGemini:
		return "fromYahooGemini"
	case FromSnapchat:
		return "fromSnapchat"
	case FromKing:
		return "fromKing"
	case FromTencent:
		return "fromTencent"
	case FromGoogleMarketingPlatform:
		return "fromGoogleMarketingPlatform"
	case FromExternalClickId:
		return "fromExternalClickId"
	default:
		return "unknown"
	}
}
