package fraud

import "errors"

//proteus:generate
type Kind uint8

const (
	// NOTE: Don't change any values here! just add new ones
	// we are not using iota here to avoid changing values when adding more types in the future

	None Kind = 0

	// device level fraud, usually detected based on the install request and remembered forever
	AnonymousTraffic       Kind = 1
	DeprecatedSdkVersion   Kind = 4 // not created anymore, legacy only
	InvalidSignature       Kind = 5
	UnverifiableOsVersion  Kind = 7
	MalformedAdvertisingId Kind = 8

	// engagement fraud, the device is fine, but we don't trust the engagement we found
	TooManyEngagements  Kind = 2
	DistributionOutlier Kind = 3
	EngagementInjection Kind = 6

	fraudLabel = "Untrusted Devices"
)

var errorBadFraudKind = errors.New("invalid fraud kind")

func (kind Kind) IsFraud() bool {
	return kind != None
}

func (kind Kind) String() string {
	switch kind {
	case None:
		return ""
	case AnonymousTraffic:
		return "Anonymous traffic"
	case DeprecatedSdkVersion:
		return "Deprecated SDK Version"
	case InvalidSignature:
		return "Invalid signature"
	case UnverifiableOsVersion:
		return "Unverifiable OS Version"
	case TooManyEngagements:
		return "Too many engagements"
	case DistributionOutlier:
		return "Distribution outlier"
	case EngagementInjection:
		return "Engagement injection"
	case MalformedAdvertisingId:
		return "Malformed advertising id"
	default:
		return "Unknown kind"
	}
}

func (kind Kind) MetricKey() string {
	switch kind {
	case None:
		return "none"
	case AnonymousTraffic:
		return "anonymous-traffic"
	case DeprecatedSdkVersion:
		return "deprecated-sdk-version"
	case InvalidSignature:
		return "invalid-signature"
	case UnverifiableOsVersion:
		return "unverifiable-os-version"
	case TooManyEngagements:
		return "too-many-engagements"
	case DistributionOutlier:
		return "distribution-outlier"
	case EngagementInjection:
		return "engagement-injection"
	case MalformedAdvertisingId:
		return "malformed-advertising-id"
	default:
		return "unknown"
	}
}

// Don't modify this without first updating the rejecttype enum in production dbs
func (kind Kind) DbValue() string {
	switch kind {
	case None:
		return "none"
	case AnonymousTraffic:
		return "anon_ip"
	case DeprecatedSdkVersion:
		return "deprecated_sdk_version"
	case InvalidSignature:
		return "invalid_signature"
	case UnverifiableOsVersion:
		return "unverifiable_os_version"
	case TooManyEngagements:
		return "too_many_engagements"
	case DistributionOutlier:
		return "distribution_outlier"
	case EngagementInjection:
		return "engagement_injection"
	case MalformedAdvertisingId:
		return "malformed_advertising_id"
	default:
		return "unknown"
	}
}

// only add those which we want to create "Untrusted Devices" subtrackers for
func (kind Kind) TrackerLabels() (labels []string, err error) {
	switch kind {
	case AnonymousTraffic:
		return []string{fraudLabel, "Anonymous IPs"}, nil
	case DeprecatedSdkVersion:
		return []string{fraudLabel, "Deprecated SDK Version"}, nil
	case InvalidSignature:
		return []string{fraudLabel, "Invalid Signature"}, nil
	case UnverifiableOsVersion:
		return []string{fraudLabel, "Unverifiable OS Version"}, nil
	case MalformedAdvertisingId:
		return []string{fraudLabel, "Malformed Advertising ID"}, nil
	default: // includes None and engagement fraud kinds
		return nil, errorBadFraudKind
	}
}
