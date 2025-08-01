package instrumenttypes

import "fmt"

type Instrument string

const (
	InstrumentEquity Instrument = "equity"
	InstrumentFutStk Instrument = "futstk"
	InstrumentOptStk Instrument = "optstk"
	InstrumentFutIdx Instrument = "futidx"
	InstrumentOptIdx Instrument = "optidx"
)

func (s Instrument) String() string {
	return string(s)
}

func ParseInstrument(s string) (Instrument, error) {
	switch s {
	case string(InstrumentEquity):
		return InstrumentEquity, nil
	case string(InstrumentFutStk):
		return InstrumentFutStk, nil
	case string(InstrumentOptStk):
		return InstrumentOptStk, nil
	case string(InstrumentFutIdx):
		return InstrumentFutIdx, nil
	case string(InstrumentOptIdx):
		return InstrumentOptIdx, nil
	default:
		return "", fmt.Errorf("invalid instrumentType: %s", s)
	}
}
