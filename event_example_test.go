package benchserder

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleEvent_MarshalUnmarshal() {
	for _, tf := range testEvents {
		jsonData, _ := readFile(tf)
		var event Event
		err := json.Unmarshal(jsonData, &event)
		if err != nil {
			panic(err)
		}

		m1Data, _ := marshalMsgpack(event)
		m2Data, _ := marshalCodecMsgpack(event)
		c1Data, _ := marshalCodecCBOR(event)
		c2Data, _ := marshalFxamackerCBOR(event)
		bData, _ := marshalBSON(event)
		//pbData, _ := marshalEventPB(event)

		fmt.Fprintf(os.Stdout, "source %s\n", tf)
		fmt.Fprintf(os.Stdout, "--- json:\t\t %d\n", len(jsonData))
		fmt.Fprintf(os.Stdout, "--- vmihailenco/msgpack: %d\n", len(m1Data))
		fmt.Fprintf(os.Stdout, "--- codec/msgpack:\t %d\n", len(m2Data))
		fmt.Fprintf(os.Stdout, "--- codec/cbor:\t\t %d\n", len(c1Data))
		fmt.Fprintf(os.Stdout, "--- fxamacker/cbor:\t %d\n", len(c2Data))
		fmt.Fprintf(os.Stdout, "--- bson:\t\t %d\n", len(bData))
		//fmt.Fprintf(os.Stdout, "--- proto:\t\t %d\n", len(pbData))

		// Output:
	}
}

/*
func marshalEventPB(event Event) ([]byte, error) {
	e := &pb.Event{
		AppToken:              event.AppToken,
		Tracker:               event.Tracker,
		Adid:                  event.Adid,
		FacebookAttributionId: event.FacebookAttributionId,
		FacebookAnonId:        event.FacebookAnonId,
		TrackingEnabled:       event.TrackingEnabled,
		ClickTime:             event.ClickTime,
		FirstSessionTime:      event.FirstSessionTime,
		LastSessionTime:       event.LastSessionTime,
		LastEventTime:         event.LastEventTime,
		LastRevenueTime:       event.LastRevenueTime,
		CreatedAt:             event.CreatedAt,
		ReceivedAt:            event.ReceivedAt,
		InstallTime:           event.InstallTime,
		InstallTracker:        event.InstallTracker,
		InstallCountry:        event.InstallCountry,
		InstallImpressionBased: pb.Bool{
			Bool:  event.InstallImpressionBased.Bool,
			Valid: event.InstallImpressionBased.Valid,
		},
		EventToken: event.EventToken,
		Revenue: pb.Amount{
			OriginalAmount:    event.Revenue.OriginalAmount,
			OriginalCurrency:  event.Revenue.OriginalCurrency,
			BaseAmount:        event.Revenue.BaseAmount,
			ReportingAmount:   event.Revenue.ReportingAmount,
			ReportingCurrency: event.Revenue.ReportingCurrency,
		},
		DeviceType:  event.DeviceType,
		Environment: event.Environment,
		NullSdkLevel: pb.Int{
			Int:   event.NullSdkLevel.Int,
			Valid: event.NullSdkLevel.Valid,
		},
		ZoneOffset:  event.ZoneOffset,
		FraudKind:   event.FraudKind,
		PingbackUrl: event.PingbackUrl,
		CallbackData: &pb.Data{
			TermsSigned:     event.CallbackData.TermsSigned,
			ReferenceTag:    event.CallbackData.ReferenceTag,
			EngagementType:  event.CallbackData.EngagementType,
			Tracker:         event.CallbackData.Tracker,
			InstallTracker:  event.CallbackData.InstallTracker,
			LastTracker:     event.CallbackData.LastTracker,
			OutdatedTracker: event.CallbackData.OutdatedTracker,
			AppToken:        event.CallbackData.AppToken,
			Adid:            event.CallbackData.Adid,
			OsName:          event.CallbackData.OsName,
			OsVersion:       event.CallbackData.OsVersion,
			OsBuild:         event.CallbackData.OsBuild,
			ApiLevel:        event.CallbackData.ApiLevel,
			Nonce:           event.CallbackData.Nonce,
			Random:          event.CallbackData.Random,
			AppName:         event.CallbackData.AppName,
			AppVersion:      event.CallbackData.AppVersion,
			AppVersionShort: event.CallbackData.AppVersionShort,
			Idfa:            event.CallbackData.Idfa,
			Idfv:            event.CallbackData.Idfv,
			AndroidId:       event.CallbackData.AndroidId,
			GoogleAdid:      event.CallbackData.GoogleAdid,
			FireAdid:        event.CallbackData.FireAdid,
			WebUuid:         event.CallbackData.WebUuid,
			WinUdid:         event.CallbackData.WinUdid,
			WinHwid:         event.CallbackData.WinHwid,
			WinNaid:         event.CallbackData.WinNaid,
			WinAdid:         event.CallbackData.WinAdid,
			SimSlotIds: pb.SimSlotIds{
				Imeis:    nil,
				Meids:    nil,
				DeviceId: "",
			},
			EventToken: event.CallbackData.EventToken,
			RevenueData: pb.Amount{
				OriginalAmount:    event.CallbackData.RevenueData.OriginalAmount,
				OriginalCurrency:  event.CallbackData.RevenueData.OriginalCurrency,
				BaseAmount:        event.CallbackData.RevenueData.BaseAmount,
				ReportingAmount:   event.CallbackData.RevenueData.ReportingAmount,
				ReportingCurrency: event.CallbackData.RevenueData.ReportingCurrency,
			},
			Cost: pb.Cost{
				Amount: pb.Amount{
					OriginalAmount:    event.CallbackData.Cost.Amount.OriginalAmount,
					OriginalCurrency:  event.CallbackData.Cost.Amount.OriginalCurrency,
					BaseAmount:        event.CallbackData.Cost.Amount.BaseAmount,
					ReportingAmount:   event.CallbackData.Cost.Amount.ReportingAmount,
					ReportingCurrency: event.CallbackData.Cost.Amount.ReportingCurrency,
				},
				CostType: event.CallbackData.Cost.CostType,
			},
			CostIdMd5:          event.CallbackData.CostIdMd5,
			IpAddress:          event.CallbackData.IpAddress,
			ProxyIpAddress:     event.CallbackData.ProxyIpAddress,
			ServerIp:           event.CallbackData.ServerIp,
			MacSha1:            event.CallbackData.MacSha1,
			MacMd5:             event.CallbackData.MacMd5,
			Environment:        event.CallbackData.Environment,
			Referrer:           event.CallbackData.Referrer,
			Country:            event.CallbackData.Country,
			Language:           event.CallbackData.Language,
			Region:             event.CallbackData.Region,
			Mcc:                event.CallbackData.Mcc,
			Mnc:                event.CallbackData.Mnc,
			Manufacturer:       event.CallbackData.Manufacturer,
			ClientSdk:          event.CallbackData.ClientSdk,
			StoreName:          event.CallbackData.StoreName,
			StoreAppId:         event.CallbackData.StoreAppId,
			SessionCount:       event.CallbackData.SessionCount,
			GlobalSessionCount: event.CallbackData.GlobalSessionCount,
			PushToken:          event.CallbackData.PushToken,
			TimeSpent:          event.CallbackData.TimeSpent,
			LastTimeSpent:      event.CallbackData.LastTimeSpent,
			DeviceName:         event.CallbackData.DeviceName,
			DeviceType:         event.CallbackData.DeviceType,
			TrackingEnabled:    event.CallbackData.TrackingEnabled,
			ZoneOffset: pb.Int{
				Int:   event.CallbackData.ZoneOffset.Int,
				Valid: event.CallbackData.ZoneOffset.Valid,
			},
			ClickLabel:      event.CallbackData.ClickLabel,
			UserAgent:       event.CallbackData.UserAgent,
			ActivityKind:    event.CallbackData.ActivityKind,
			Deeplink:        event.CallbackData.Deeplink,
			SearchTerm:      event.CallbackData.SearchTerm,
			ClickReferer:    event.CallbackData.ClickReferer,
			ImpressionBased: event.CallbackData.ImpressionBased,
			FraudKind:       event.CallbackData.FraudKind,
			CpuType:         event.CallbackData.CpuType,
			HardwareName:    event.CallbackData.HardwareName,
			NetworkType:     event.CallbackData.NetworkType,
			SecretId: pb.Int{
				Int:   event.CallbackData.SecretId.Int,
				Valid: event.CallbackData.SecretId.Valid,
			},
			IsImported:            event.CallbackData.IsImported,
			IsServerToServer:      event.CallbackData.IsServerToServer,
			IsServerToServerBased: event.CallbackData.IsServerToServerBased,
			ClickTime:             event.CallbackData.ClickTime,
			InstalledAt:           event.CallbackData.InstalledAt,
			FirstSessionTime:      event.CallbackData.FirstSessionTime,
			FirstCountry:          event.CallbackData.FirstCountry,
			FirstDeviceType:       event.CallbackData.FirstDeviceType,
			FirstOsName:           event.CallbackData.FirstOsName,
			ReceivedAt:            event.CallbackData.ReceivedAt,
			CreatedAtTime:         event.CallbackData.CreatedAtTime,
			LastSessionTime:       event.CallbackData.LastSessionTime,
			UninstalledAt:         event.CallbackData.UninstalledAt,
			ReinstalledAt:         event.CallbackData.ReinstalledAt,
			AttributionUpdatedAt:  event.CallbackData.AttributionUpdatedAt,
			SessionLength:         event.CallbackData.SessionLength,
			PingbackTtl: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			ClickAttributionWindow: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			ImpressionAttributionWindow: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			FingerprintAttributionWindow: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			ReattributionAttributionWindow: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			InactiveUserDefinition: pb.Duration{
				Duration: 0,
				Valid:    false,
			},
			PartnerParameters:         event.CallbackData.PartnerParameters,
			PartnerSdkParams:          event.CallbackData.PartnerSdkParams,
			PublisherParams:           event.CallbackData.PublisherParams,
			DynamicCallbackParameters: event.CallbackData.DynamicCallbackParameters,
			ApiPartnerParams:          event.CallbackData.ApiPartnerParams,
			AdImpressionsCount:        event.CallbackData.AdImpressionsCount,
			AdRevenueNetwork:          event.CallbackData.AdRevenueNetwork,
			AdRevenueUnit:             event.CallbackData.AdRevenueUnit,
			AdRevenuePlacement:        event.CallbackData.AdRevenuePlacement,
			AdMediationPlatform:       event.CallbackData.AdMediationPlatform,
		},
		FirstOsName:        event.FirstOsName,
		FirstCountry:       event.FirstCountry,
		FirstDeviceType:    event.FirstDeviceType,
		ImpressionBased:    event.ImpressionBased,
		DeviceReattributed: event.DeviceReattributed,
	}
	return proto.Marshal(e)
}
*/
