package custom_errors

import "net/http"

type Error struct {
	status  int
	Success bool              `json:"success"`
	Msg     string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e Error) Error() string {
	return e.Msg
}

func (e Error) Status() int {
	return e.status
}

func (e Error) MsgAndErr() map[string]any {
	return map[string]any{
		"msg": e.Msg,
		"err": e.Errors,
	}
}

func Internal(msg string) Error {
	if msg == "" {
		msg = "some thing went wrong please try again"
	}

	return Error{
		status: http.StatusInternalServerError,
		Msg:    msg,
	}
}

func BadReq() Error {
	return Error{
		status: http.StatusInternalServerError,
		Msg:    "bad request",
	}
}
