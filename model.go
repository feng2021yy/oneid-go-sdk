package oneid_go_sdk

type Game_id struct {
	ChannelId   int32
	ChannelCode string
	OpenId      string
}

type Wx_id struct {
}

// 其他渠道下的id】
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
