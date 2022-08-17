package types

type Msg struct {
	Error string
}

func ErrorMessage(msg string) Msg {
	return Msg{
		Error: msg,
	}
}
