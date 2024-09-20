package oneid_go_sdk

// ScrmId 社群账号id
type ScrmId struct {
}

// GameOpenId 游戏openId
type GameOpenId struct {
	ChannelId   int32
	ChannelCode string
	OpenId      string
}

// GameRoleId 游戏角色id
type GameRoleId struct {
}

// MiniProgramId 小程序id
type MiniProgramId struct {
}

// WxOpenId 微信开放平台id
type WxOpenId struct {
}

// OtherId 其他渠道下的id
type OtherId struct {
	ChannelId   int32
	ChannelCode string
	OpenId      string
}

type OneIdResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// PassMessage 传递给下游pulsar队列的信息结构体
type PassMessage struct {
	ActionType   int //操作类型
	MatchType    int //两个账号对应关系
	SourceIdType int //A账号类型
	DestIdType2  int //B账号类型
}
