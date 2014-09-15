package dfa

const (
	invalidID = -1
)

type StateLabel int

const (
	notFinal          StateLabel = 0
	defaultFinal      StateLabel = 1
	labeledFinalStart StateLabel = 2
)

func (l StateLabel) final() bool {
	return l >= defaultFinal
}

func (l StateLabel) labeled() bool {
	return l >= labeledFinalStart
}

func (l StateLabel) toInternal() StateLabel {
	return l + labeledFinalStart
}

func (l StateLabel) toExternal() int {
	return int(l - labeledFinalStart)
}
