package models

const (
	TimeFormatLayout = "2006-01-02 15:04:05"

	StatusOK           = 200
	InvalidReqIsNil    = 400
	InvalidReqIsNilMsg = "请求参数错误: req is nil"

	ServerErr = 500

	AESSalt = "321423u9y8d2fwfl" // TODO
)
