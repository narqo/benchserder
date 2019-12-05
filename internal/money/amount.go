package money

import (
	"fmt"
	"strconv"
)

type Amount struct {
	// original values from the SDK
	OriginalAmount   float64
	OriginalCurrency string

	// converted to base currency USD
	BaseAmount float64

	// converted to reporting currency as per app setting
	ReportingAmount   float64
	ReportingCurrency string
}

func (amount Amount) String() string {
	if amount.IsEmpty() {
		return ""
	}

	if amount.ReportingCurrency != "" {
		return fmt.Sprintf("[R %.6g %s; %.6g USD; %.6g %s]",
			amount.OriginalAmount,
			amount.OriginalCurrency,
			amount.BaseAmount,
			amount.ReportingAmount,
			amount.ReportingCurrency,
		)
	}

	if amount.BaseAmount != 0 {
		return fmt.Sprintf("[R %.6g %s; %.6g USD]",
			amount.OriginalAmount,
			amount.OriginalCurrency,
			amount.BaseAmount,
		)
	}

	return fmt.Sprintf("[R %.6g %s]",
		amount.OriginalAmount,
		amount.OriginalCurrency,
	)
}

func (amount Amount) GetOriginalAmount() string {
	if amount.IsEmpty() {
		return ""
	}
	return formatFloat(amount.OriginalAmount)
}

func (amount Amount) GetOriginalAmountCents() string {
	if amount.IsEmpty() {
		return ""
	}
	return formatFloat(100 * amount.OriginalAmount)
}

func (amount Amount) GetBaseAmount() string {
	if amount.IsEmpty() {
		return ""
	}
	return formatFloat(amount.BaseAmount)
}

func (amount Amount) GetBaseAmountCents() string {
	if amount.IsEmpty() {
		return ""
	}
	return formatFloat(100 * amount.BaseAmount)
}

func (amount Amount) GetBaseAmountMicro() string {
	if amount.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%.0f", 1000000*amount.BaseAmount)
}

func (amount Amount) GetReportingAmount() string {
	if amount.ReportingCurrency == "" {
		return ""
	}
	return formatFloat(amount.ReportingAmount)
}

func (amount Amount) GetOriginalCurrency() string {
	return amount.OriginalCurrency
}

func (amount Amount) GetReportingCurrency() string {
	return amount.ReportingCurrency
}

func (amount Amount) IsEmpty() bool {
	return amount.OriginalCurrency == ""
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (amount Amount) IsDefined() bool {
	return !amount.IsEmpty()
}

func (amount Amount) Scale(factor float64) Amount {
	return Amount{
		OriginalAmount:    amount.OriginalAmount * factor,
		OriginalCurrency:  amount.OriginalCurrency,
		BaseAmount:        amount.BaseAmount * factor,
		ReportingAmount:   amount.ReportingAmount * factor,
		ReportingCurrency: amount.ReportingCurrency,
	}
}

func formatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
