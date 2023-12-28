package util

func GetStatusStr(status bool) string {
	if status {
		return "ACTIVE"
	} else {
		return "NON ACTIVE"
	}
}
