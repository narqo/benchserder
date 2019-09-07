package money

import (
	"fmt"
	"strings"
)

type CostType int

const (
	CostNone = CostType(iota)
	CostPerMille
	CostPerClick
	CostPerInstall
)

func NewCostType(costTypeStr string) (CostType, error) {
	switch strings.ToUpper(costTypeStr) {
	case "":
		return CostNone, nil
	case "CPM", "IMPRESSION", "VIEW":
		return CostPerMille, nil
	case "CPC", "CLICK":
		return CostPerClick, nil
	case "CPI", "INSTALL":
		return CostPerInstall, nil
	default:
		return CostNone, fmt.Errorf("unknown cost type string: %s", costTypeStr)
	}
}

func (cost CostType) String() string {
	switch cost {
	case CostNone:
		return ""
	case CostPerMille:
		return "CPM"
	case CostPerClick:
		return "CPC"
	case CostPerInstall:
		return "CPI"
	default:
		return "unknown"
	}
}

// represents costs that could be associated with an incoming activity
//easyjson:json
type Cost struct {
	Amount   Amount
	CostType CostType

	XXX_unrecognized []byte `json:"-"`
}

func NewCost(amount Amount, costType CostType) Cost {
	if costType == CostNone {
		return Cost{}
	}

	if amount.IsEmpty() {
		return Cost{}
	}

	return Cost{Amount: amount, CostType: costType}
}

func (cost Cost) CostPerMille() Amount {
	return cost.getAmountIfSameType(CostPerMille)
}

func (cost Cost) CostPerImpression() Amount {
	return cost.getAmountIfSameType(CostPerMille).Scale(1e-3)
}

func (cost Cost) CostPerClick() Amount {
	return cost.getAmountIfSameType(CostPerClick)
}

func (cost Cost) CostPerInstall() Amount {
	return cost.getAmountIfSameType(CostPerInstall)
}

func (cost Cost) getAmountIfSameType(costType CostType) Amount {
	if cost.CostType == costType {
		return cost.Amount
	}
	return Amount{}
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (cost Cost) IsDefined() bool {
	return cost.Amount.IsDefined() && cost.CostType != CostNone
}
