package go_vagrant

type outputHandler interface {
	handleOutput(target, key string, message []string)
}
