package goyapi

import "fmt"

type Logger interface {
	Printf(format string, v ...interface{})
}

type Resp struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Resp) Error() string {
	if r.ErrCode == 0 {
		return ""
	}
	return fmt.Sprintf("%d:%s", r.ErrCode, r.ErrMsg)
}

type ReqKvItem struct {
	Required string `json:"required,omitempty"`
	Name     string `json:"name"`
	Value    string `json:"value,omitempty"`
	Example  string `json:"example,omitempty"`
	Desc     string `json:"desc,omitempty"`
	Type     string `json:"type,omitempty"`
}

type InterfaceAddReq struct {
	Token               string      `json:"token"`
	ReqQuery            []ReqKvItem `json:"req_query,omitempty"`
	ReqHeaders          []ReqKvItem `json:"req_headers,omitempty"`
	ReqBodyForm         []ReqKvItem `json:"req_body_form,omitempty"`
	ReqParams           []ReqKvItem `json:"req_params,omitempty"`
	ReqBodyType         string      `json:"req_body_type" `
	ReqBodyOther        string      `json:"req_body_other,omitempty"`
	ReqBodyIsJsonSchema bool        `json:"req_body_is_json_schema"`
	ResBodyIsJsonSchema bool        `json:"res_body_is_json_schema"`
	Title               string      `json:"title"`
	Catid               int         `json:"catid"`
	Path                string      `json:"path"`
	ResBodyType         string      `json:"res_body_type"`
	ResBody             string      `json:"res_body"`
	Message             string      `json:"message"`
	Desc                string      `json:"desc,omitempty"`
	Method              string      `json:"method"`
	ID                  string      `json:"id,omitempty"`
}

type InterfaceSaveResp []InterfaceDetail

type InterfaceSaveReq = InterfaceAddReq

type InterfaceDetail struct {
	ID int `json:"_id"`
}

type Project struct {
	IsMockOpen  bool   `json:"is_mock_open"`
	Strice      bool   `json:"strice"`
	IsJSON5     bool   `json:"is_json5"`
	ID          int    `json:"_id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Basepath    string `json:"basepath"`
	ProjectType string `json:"project_type"`
	UID         int    `json:"uid"`
	GroupID     int    `json:"group_id"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	AddTime     int    `json:"add_time"`
	UpTime      int    `json:"up_time"`
}

type Cat struct {
	Index     int    `json:"index"`
	ID        int    `json:"_id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	Desc      string `json:"desc"`
	UID       int    `json:"uid"`
	AddTime   int    `json:"add_time"`
	UpTime    int    `json:"up_time"`
	V         int    `json:"__v"`
}

type CatList []Cat

type CatAddReq struct {
	Desc      string `json:"desc"`
	Name      string `json:"name"`
	ProjectId int    `json:"project_id"`
	Token     string `json:"token"`
}
