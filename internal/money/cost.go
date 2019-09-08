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
}

// IsDefined implements easyjson.Optional to support "omitempty" json tag.
func (cost Cost) IsDefined() bool {
	return cost.Amount.IsDefined() && cost.CostType != CostNone
}
