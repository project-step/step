package session

import "strings"

type CmdType int

const (
	CmdType_Select   CmdType = 0
	CmdType_MakeView CmdType = 1
	CmdType_Other    CmdType = 256
)

func parseCmdType(query string) CmdType {
	s := strings.Split(query, " ")
	if len(s) > 0 && strings.ToUpper(s[0]) == "SELECT" {
		return CmdType_Select
	} else if len(s) > 1 && strings.ToUpper(s[0]) == "MAKE" && strings.ToUpper(s[1]) == "VIEW" {
		return CmdType_MakeView
	} else {
		return CmdType_Other
	}
}
