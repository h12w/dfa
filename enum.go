package dfa

const (
	invalidID      = -1
	trivialFinalID = -2
	excludingID    = -3
)

type stateLabel int

const (
	notFinal          stateLabel = 0
	defaultFinal      stateLabel = 1
	labeledFinalStart stateLabel = 2
)

func (l stateLabel) final() bool {
	return l >= defaultFinal
}

func (l stateLabel) labeled() bool {
	return l >= labeledFinalStart
}

func (l stateLabel) toInternal() stateLabel {
	return l + labeledFinalStart
}

func (l stateLabel) toExternal() int {
	if l >= labeledFinalStart {
		return int(l - labeledFinalStart)
	}
	panic("machine is not labeled")
}
