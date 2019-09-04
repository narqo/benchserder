package engagementtype

import "log"

//proteus:generate
type Type int

// types who's index is lower than that are for tests, or legacy types
const knownTypesOffset = 10

const (
	// NOTE: Don't change any values here! just add new ones
	// we are not using iota here to avoid changing values when adding more types in the future
	// all click types should be even numbers and all impression types odd numbers

	None    Type = -1 // no engagement type (currently also used for organic)
	Unknown Type = 0  // unknown engagement type

	// legacy type represents types 1..10 that we used to have; it must be used in testing only
	TestLegacyType Type = 1

	// engagement types from adjust server
	AdjustDtagClick         Type = 10  // click from adjust having device tags
	AdjustDtagImpression    Type = 11  // impression from adjust having device tags
	AdjustDtagFallbackClick Type = 100 // click from adjust having device tags and fallback flag
	AdjustIpClick           Type = 12  // click from adjust having only IP address
	AdjustIpImpression      Type = 13  // impression from adjust having only IP address
	AdjustIpFallbackClick   Type = 102 // click from adjust having only IP address and fallback flag

	// engagement types from attribution worker apis
	SdkDeeplinkClick                  Type = 14  // click from sdk based on deeplink
	SdkReftagClick                    Type = 16  // click from sdk click reference tag
	SdkReftagFallbackClick            Type = 104 // click from sdk click reference tag with fallback flag
	SdkIadClick                       Type = 18  // click from sdk click iad
	SdkIad3Click                      Type = 20  // click from sdk click iad3
	SdkAppleSearchAdsClick            Type = 22  // click from sdk click - apple search Ads
	SdkGoogleOrganicSearchClick       Type = 24  // click from sdk click - google organic search
	SdkPreinstallFallbackClick        Type = 106 // click from sdk click - adjust tracker
	FacebookClick                     Type = 26  // click from facebook attribution api
	FacebookImpression                Type = 27  // impression from facebook attribution api
	GoogleClick                       Type = 28  // click from old google attribution api (legacy)
	AdwordsClick                      Type = 30  // click from adwords attribution api
	AdwordsImpression                 Type = 31  // impression from google adwords attribution api
	TwitterClick                      Type = 32  // click from twitter attribution api
	TwitterImpression                 Type = 33  // impression from twitter attribution api
	DoubleclickClick                  Type = 34  // click from doubleclick attribution api
	DoubleclickImpression             Type = 35  // impression from doubleclick attribution api
	YahooJapanClick                   Type = 36  // click from yahoo japan attribution api
	YahooGeminiClick                  Type = 38  // click from yahoo gemini attribution api
	YahooGeminiImpression             Type = 39  // impression from yahoo gemini attribution api
	SnapchatClick                     Type = 40  // click from snapchat attribution api
	SnapchatImpression                Type = 41  // impression from snapchat attribution api
	ExternalClickIdClick              Type = 42  // click from referrer sdk click with an external_click_id attached
	KingClick                         Type = 44  // click from king attribution api
	KingImpression                    Type = 45  // impression from king attribution api
	ReferrerClick                     Type = 46  // click from session referrer (old sdk when we didn't have sdk click)
	TencentClick                      Type = 48  // click from tencent attribution api
	TencentImpression                 Type = 49  // impression from tencent attribution api
	GoogleMarketingPlatformClick      Type = 50  // click from google marketing platform api
	GoogleMarketingPlatformImpression Type = 51  // impression from google marketing platform api
)

type Kind int

const (
	KindUnknown Kind = iota
	KindFallbackClick
	KindImpression
	KindClick
)

func (kind Kind) String() string {
	switch kind {
	case KindClick:
		return "click"
	case KindImpression:
		return "impression"
	case KindFallbackClick:
		return "fallbackClick"
	default:
		return "unknown"
	}
}

// Score returns the score for an engagement type.
//
// higher score means better engagement
// used early to prefer clicks over impressions even if the click is older then
// the impression
func (engagementType Type) Score() int {
	var score int

	switch engagementType.Kind() {
	case KindClick:
		score = 1 << 3
	case KindImpression:
		score = 1 << 2
	case KindFallbackClick:
		score = 1 << 1
	default:
		return 0 // unknown
	}

	// prefer non-fingerprint engagement of the same kind
	if engagementType.Source() != FromIpList {
		score |= 1
	}

	return score
}

// higher score means better engagement
// used late, only when both engagements have the same engagement type. in that
// case prefer attribution worker engagements over adjust server engagements
// for legacy reasons. we might remove this later
func (engagementType Type) ServiceScore() int {
	switch engagementType.Kind() {
	case KindClick:
		if engagementType.isAttributionApi() {
			return 4
		}
		return 3
	case KindImpression:
		return 2
	case KindFallbackClick:
		return 1
	}
	return 0
}

// Score after Score(), ServiceScore(), ClickTime, are all the same
func (engagementType Type) SourceScore() int {
	// for now we only prioritize SdkReftag, we may consider the order later
	switch engagementType {
	case SdkReftagClick:
		return 3
	default:
		return 0
	}
}

func (engagementType Type) IsImpression() bool {
	return engagementType.Kind() == KindImpression
}

func (engagementType Type) Kind() Kind {
	// handle fallback types
	switch engagementType {
	case Unknown,
		None,
		TestLegacyType:
		return KindUnknown
	case AdjustDtagFallbackClick,
		AdjustIpFallbackClick,
		SdkReftagFallbackClick,
		SdkPreinstallFallbackClick:
		return KindFallbackClick
	}

	if engagementType < knownTypesOffset {
		// fallback for legacy types
		return KindUnknown
	}

	if engagementType%2 == 0 {
		return KindClick // clicks are even
	} else {
		return KindImpression // impressions are odd
	}
}

func (engagementType Type) IsAdjust() bool {
	switch engagementType {
	case AdjustDtagClick,
		AdjustDtagImpression,
		AdjustDtagFallbackClick,
		AdjustIpClick,
		AdjustIpImpression,
		AdjustIpFallbackClick:
		return true
	default:
		return false
	}
}

func (engagementType Type) IsSan() bool {
	switch engagementType {
	case SdkAppleSearchAdsClick,
		FacebookClick,
		FacebookImpression,
		GoogleClick,
		GoogleMarketingPlatformClick,
		GoogleMarketingPlatformImpression,
		AdwordsClick,
		AdwordsImpression,
		TwitterClick,
		TwitterImpression,
		DoubleclickClick,
		DoubleclickImpression,
		YahooJapanClick,
		YahooGeminiClick,
		YahooGeminiImpression,
		SnapchatClick,
		SnapchatImpression,
		KingClick,
		KingImpression,
		TencentClick,
		TencentImpression:
		return true
	default:
		return false
	}
}

func (engagementType Type) isAttributionApi() bool {
	if engagementType.Kind() == KindUnknown {
		return false
	}
	return !engagementType.IsAdjust()
}

func (engagementType Type) Source() Source {
	switch engagementType {
	case AdjustDtagClick:
		return FromMappedDtag
	case AdjustDtagImpression:
		return FromMappedDtag
	case AdjustDtagFallbackClick:
		return FromMappedDtag
	case AdjustIpClick:
		return FromIpList
	case AdjustIpImpression:
		return FromIpList
	case AdjustIpFallbackClick:
		return FromIpList
	case SdkDeeplinkClick:
		return FromSdkDeeplink
	case SdkReftagClick:
		return FromSdkReftag
	case SdkReftagFallbackClick:
		return FromSdkReftag
	case SdkIadClick:
		return FromSdkIad
	case SdkIad3Click:
		return FromSdkIad3
	case SdkAppleSearchAdsClick:
		return FromSdkAppleSearchAds
	case SdkGoogleOrganicSearchClick:
		return FromSdkGoogleOrganicSearch
	case SdkPreinstallFallbackClick:
		return FromSdkAdjustTracker
	case FacebookClick:
		return FromFacebook
	case FacebookImpression:
		return FromFacebook
	case GoogleClick:
		return FromGoogle
	case AdwordsClick:
		return FromAdwords
	case AdwordsImpression:
		return FromAdwords
	case TwitterClick:
		return FromTwitter
	case TwitterImpression:
		return FromTwitter
	case DoubleclickClick:
		return FromDoubleClick
	case DoubleclickImpression:
		return FromDoubleClick
	case YahooJapanClick:
		return FromYahooJapan
	case YahooGeminiClick:
		return FromYahooGemini
	case YahooGeminiImpression:
		return FromYahooGemini
	case SnapchatClick:
		return FromSnapchat
	case SnapchatImpression:
		return FromSnapchat
	case ExternalClickIdClick:
		return FromExternalClickId
	case KingClick:
		return FromKing
	case KingImpression:
		return FromKing
	case TencentClick:
		return FromTencent
	case TencentImpression:
		return FromTencent
	case GoogleMarketingPlatformClick:
		return FromGoogleMarketingPlatform
	case GoogleMarketingPlatformImpression:
		return FromGoogleMarketingPlatform
	case ReferrerClick:
		return FromReferrer
	default:
		return FromUnknown
	}
}

func (engagementType Type) MatchType() string {
	switch engagementType {
	case AdjustIpClick,
		AdjustIpImpression,
		AdjustIpFallbackClick:
		return MatchTypeIpList
	case AdjustDtagClick,
		AdjustDtagImpression,
		AdjustDtagFallbackClick:
		return MatchTypeMappedDtag
	case ReferrerClick:
		return MatchTypeReferrer
	case SdkReftagClick,
		SdkReftagFallbackClick:
		return MatchTypeSdkReftag
	case SdkGoogleOrganicSearchClick:
		return MatchTypeSdkGoogleOrganicSearch
	case SdkPreinstallFallbackClick:
		return MatchTypeSdkAdjustTracker
	case SdkDeeplinkClick:
		return MatchTypeSdkDeeplink
	case SdkIadClick:
		return MatchTypeSdkIad
	case SdkIad3Click:
		return MatchTypeSdkIad3
	case SdkAppleSearchAdsClick:
		return MatchTypeSdkAppleSearchAds
	case FacebookClick,
		FacebookImpression:
		return MatchTypeFacebook
	case GoogleClick:
		return MatchTypeGoogle
	case AdwordsClick,
		AdwordsImpression:
		return MatchTypeAdwords
	case TwitterClick,
		TwitterImpression:
		return MatchTypeTwitter
	case DoubleclickClick,
		DoubleclickImpression:
		return MatchTypeDoubleClick
	case YahooJapanClick:
		return MatchTypeYahooJapan
	case YahooGeminiClick,
		YahooGeminiImpression:
		return MatchTypeYahooGemini
	case SnapchatClick,
		SnapchatImpression:
		return MatchTypeSnapchat
	case KingClick,
		KingImpression:
		return MatchTypeKing
	case TencentClick,
		TencentImpression:
		return MatchTypeTencent
	case GoogleMarketingPlatformClick,
		GoogleMarketingPlatformImpression:
		return MatchTypeGoogleMarketingPlatform
	case ExternalClickIdClick:
		return MatchTypeExternalClickId
	case TestLegacyType:
		return MatchTypeLegacy
	default:
		return MatchTypeNone
	}
}

func FromMatchType(matchType string, isImpression bool) Type {
	switch matchType {
	case MatchTypeNone:
		return None
	case MatchTypeIpList:
		if isImpression {
			return AdjustIpImpression
		}
		return AdjustIpClick
	case MatchTypeMappedDtag:
		if isImpression {
			return AdjustDtagImpression
		}
		return AdjustDtagClick
	case MatchTypeReferrer:
		return ReferrerClick
	case MatchTypeSdkReftag:
		return SdkReftagClick
	case MatchTypeSdkGoogleOrganicSearch:
		return SdkGoogleOrganicSearchClick
	case MatchTypeSdkAdjustTracker:
		return SdkPreinstallFallbackClick
	case MatchTypeSdkDeeplink:
		return SdkDeeplinkClick
	case MatchTypeSdkIad:
		return SdkIadClick
	case MatchTypeSdkIad3:
		return SdkIad3Click
	case MatchTypeSdkAppleSearchAds:
		return SdkAppleSearchAdsClick
	case MatchTypeFacebook:
		if isImpression {
			return FacebookImpression
		}
		return FacebookClick
	case MatchTypeGoogle:
		return GoogleClick
	case MatchTypeAdwords:
		if isImpression {
			return AdwordsImpression
		}
		return AdwordsClick
	case MatchTypeTwitter:
		if isImpression {
			return TwitterImpression
		}
		return TwitterClick
	case MatchTypeDoubleClick:
		if isImpression {
			return DoubleclickImpression
		}
		return DoubleclickClick
	case MatchTypeYahooJapan:
		return YahooJapanClick
	case MatchTypeYahooGemini:
		if isImpression {
			return YahooGeminiImpression
		}
		return YahooGeminiClick
	case MatchTypeSnapchat:
		if isImpression {
			return SnapchatImpression
		}
		return SnapchatClick
	case MatchTypeKing:
		if isImpression {
			return KingImpression
		}
		return KingClick
	case MatchTypeTencent:
		if isImpression {
			return TencentImpression
		}
		return TencentClick
	case MatchTypeGoogleMarketingPlatform:
		if isImpression {
			return GoogleMarketingPlatformImpression
		}
		return GoogleMarketingPlatformClick
	case MatchTypeExternalClickId:
		return ExternalClickIdClick
	default:
		log.Printf("DEBUG#5102: found unknown match type %s", matchType)
		return Unknown
	}
}

// FromInt filters out legacy types.
func FromInt(typeInt int) Type {
	if typeInt > 0 && typeInt < knownTypesOffset {
		log.Printf("DEBUG#5102 unknown engagement type from int %d", typeInt)
		return Unknown
	}
	return Type(typeInt)
}

func (engagementType Type) String() string {
	switch engagementType {
	case None:
		return "none"
	case Unknown:
		return "unknown"
	case AdjustDtagClick:
		return "clickAdjustDtag"
	case AdjustDtagImpression:
		return "impressionAdjustDtag"
	case AdjustIpClick:
		return "clickAdjustIp"
	case AdjustIpImpression:
		return "impressionAdjustIp"
	case AdjustDtagFallbackClick:
		return "fallbackClickAdjustDtag"
	case AdjustIpFallbackClick:
		return "fallbackClickAdjustIp"
	case SdkDeeplinkClick:
		return "clickSdkDeeplink"
	case SdkReftagClick:
		return "clickSdkReftag"
	case SdkReftagFallbackClick:
		return "fallbackClickSdkReftag"
	case SdkIadClick:
		return "clickSdkIad"
	case SdkIad3Click:
		return "clickSdkIad3"
	case SdkAppleSearchAdsClick:
		return "clickSdkAppleSearchAds"
	case SdkGoogleOrganicSearchClick:
		return "clickSdkGoogleOrganicSearch"
	case SdkPreinstallFallbackClick:
		return "fallbackClickSdkAdjustTracker"
	case FacebookClick:
		return "clickFacebook"
	case FacebookImpression:
		return "impressionFacebook"
	case GoogleClick:
		return "clickGoogle"
	case AdwordsClick:
		return "clickAdwords"
	case AdwordsImpression:
		return "impressionAdwords"
	case TwitterClick:
		return "clickTwitter"
	case TwitterImpression:
		return "impressionTwitter"
	case DoubleclickClick:
		return "clickDoubleClick"
	case DoubleclickImpression:
		return "impressionDoubleClick"
	case YahooJapanClick:
		return "clickYahooJapan"
	case YahooGeminiClick:
		return "clickYahooGemini"
	case YahooGeminiImpression:
		return "impressionYahooGemini"
	case SnapchatClick:
		return "clickSnapchat"
	case SnapchatImpression:
		return "impressionSnapchat"
	case ExternalClickIdClick:
		return "clickExternalId"
	case KingClick:
		return "clickKing"
	case KingImpression:
		return "impressionKing"
	case TencentClick:
		return "clickTencent"
	case TencentImpression:
		return "impressionTencent"
	case GoogleMarketingPlatformClick:
		return "clickGoogleMarketingPlatform"
	case GoogleMarketingPlatformImpression:
		return "impressionGoogleMarketingPlatform"
	case ReferrerClick:
		return "clickReferrer"
	default:
		return "unknown"
	}
}
