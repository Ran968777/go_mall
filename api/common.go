package api

// Response 通用响应结构
type Response struct {
	Code      string      `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      interface{} `json:"data"`
	Version   string      `json:"version"`
	Timestamp interface{} `json:"timestamp"`
	Sign      interface{} `json:"sign"`
	Fail      bool        `json:"fail"`
	Success   bool        `json:"success"`
}

// SuccessResponse 成功响应
func SuccessResponse(data interface{}) Response {
	return Response{
		Code:      "00000",
		Msg:       nil,
		Data:      data,
		Version:   "mall4j.v230424",
		Timestamp: nil,
		Sign:      nil,
		Fail:      false,
		Success:   true,
	}
}
