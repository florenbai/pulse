package wework

type Event struct {
	MsgType string `xml:"MsgType"`
	Event   string `xml:"Event"`
}

type ButtonTemplateEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	TaskId       string `xml:"TaskId"`
	CardType     string `xml:"CardType"`
	ResponseCode string `xml:"ResponseCode"`
	AgentID      int64  `xml:"AgentID"`
	QuestionKey  string `xml:"QuestionKey"`
	OptionIds    string `xml:"OptionIds"`
}
