package types

type Msg struct {
	Message string
}

func CommonResponse(msg string) Msg {
	return Msg{
		Message: msg,
	}
}
