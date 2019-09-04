package engagementtype

const (
	MatchTypeNone                    = ""
	MatchTypeIpList                  = "fingerprint"
	MatchTypeMappedDtag              = "device_tag"
	MatchTypeReferrer                = "reference_tag"
	MatchTypeSdkReftag               = "reftag"
	MatchTypeSdkGoogleOrganicSearch  = "google_organic_search"
	MatchTypeSdkAdjustTracker        = "referrer_adjust_tracker"
	MatchTypeSdkDeeplink             = "deeplink"
	MatchTypeSdkIad                  = "iad"
	MatchTypeSdkIad3                 = "iad3"
	MatchTypeSdkAppleSearchAds       = "iad3_search_ads"
	MatchTypeFacebook                = "facebook"
	MatchTypeGoogle                  = "google" // from old google api (legacy)
	MatchTypeAdwords                 = "adwords"
	MatchTypeTwitter                 = "twitter"
	MatchTypeDoubleClick             = "doubleclick"
	MatchTypeYahooJapan              = "yahoo_japan"
	MatchTypeYahooGemini             = "yahoo_gemini"
	MatchTypeSnapchat                = "snapchat"
	MatchTypeKing                    = "king"
	MatchTypeTencent                 = "tencent"
	MatchTypeGoogleMarketingPlatform = "google_marketing_platform"
	MatchTypeExternalClickId         = "external_click_id"

	// MatchTypeLegacy is only used in testing
	MatchTypeLegacy = "legacy"
)

var AllMatchTypes = []string{
	MatchTypeNone,
	MatchTypeIpList,
	MatchTypeMappedDtag,
	MatchTypeReferrer,
	MatchTypeSdkReftag,
	MatchTypeSdkGoogleOrganicSearch,
	MatchTypeSdkAdjustTracker,
	MatchTypeSdkDeeplink,
	MatchTypeSdkIad,
	MatchTypeSdkIad3,
	MatchTypeSdkAppleSearchAds,
	MatchTypeFacebook,
	MatchTypeGoogle,
	MatchTypeAdwords,
	MatchTypeTwitter,
	MatchTypeDoubleClick,
	MatchTypeYahooJapan,
	MatchTypeYahooGemini,
	MatchTypeSnapchat,
	MatchTypeKing,
	MatchTypeTencent,
	MatchTypeGoogleMarketingPlatform,
	MatchTypeExternalClickId,
}
