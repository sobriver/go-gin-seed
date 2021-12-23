package e

const (
	SUCCESS        = 10000
	FAILURE        = 10001
	INVALID_PARAMS = 10002
	PWD_NOT_MATCH  = 10003
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	FAILURE:        "fail",
	INVALID_PARAMS: "param error",
	PWD_NOT_MATCH:  "loginName or password wrong",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[FAILURE]
}
