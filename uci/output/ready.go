package output

type ReadyOutput struct{}

func (o ReadyOutput) GetOutput() string {
	return "readyok\n"
}
