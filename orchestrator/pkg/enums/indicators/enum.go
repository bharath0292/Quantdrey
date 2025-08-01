package indicators

type Indicator string

const (
	IndicatorRsi        Indicator = "rsi"
	IndicatorCci        Indicator = "cci"
	IndicatorSma        Indicator = "sma"
	IndicatorBb         Indicator = "bb"
	IndicatorSupertrend Indicator = "supertrend"
	IndicatorPivotPoint Indicator = "pp"
)

func (s Indicator) String() string {
	return string(s)
}
