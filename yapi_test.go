package goyapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/mcuadros/go-jsonschema-generator"
)

type Request struct {
	Name string  `json:"name"`
	Ids  []int64 `json:"ids"`
}

var api = "http://yapi.smart-xwork.cn"
var token = "e731411e9869750ec42e0025adcf031fae3ff33667eaf5096b80ca801fc513a6"

func Test_yapi_AddInterface(t *testing.T) {
	catId, err := getBaseCatId()
	if err != nil {
		t.Error(err)
		return
	}

	req := &InterfaceAddReq{
		ReqQuery: []ReqKvItem{},
		ReqHeaders: []ReqKvItem{
			{Name: "Content-Type", Value: "application/json"},
			{Name: "X-Access-Token", Value: ""},
		},
		ReqBodyForm:         []ReqKvItem{},
		ReqParams:           []ReqKvItem{},
		ReqBodyType:         "json",
		Title:               "用户列表",
		Catid:               catId,
		Path:                "/api/v1/home/user/list",
		ResBodyType:         "json",
		Message:             "auto",
		Method:              "POST",
		ReqBodyIsJsonSchema: true,
		ResBodyIsJsonSchema: true,
		// ID:           "",
	}

	rBody := &Request{
		Name: "zhangsan",
		Ids:  []int64{1, 3, 4},
	}

	s := &jsonschema.Document{}
	s.Read(rBody)

	// fmt.Println(s)
	req.ReqBodyOther = s.String()

	s2 := &jsonschema.Document{}
	s2.Read(Resp{})
	req.ResBody = s2.String()

	ctx := context.Background()
	yapi := NewYapi(api, token, nil)
	exists, id, err := yapi.AddInterface(ctx, req)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	t.Logf("%v,%v", exists, id)
}

func Test_yapi_SaveInterface(t *testing.T) {
	catId, err := getBaseCatId()
	if err != nil {
		t.Error(err)
		return
	}

	req := &InterfaceSaveReq{
		ReqQuery: []ReqKvItem{},
		ReqHeaders: []ReqKvItem{
			{Name: "Content-Type", Value: "application/json"},
			{Name: "X-Access-Token", Value: ""},
		},
		ReqBodyForm:         []ReqKvItem{},
		ReqParams:           []ReqKvItem{},
		ReqBodyType:         "json",
		Title:               "用户列表",
		Catid:               catId,
		Path:                "/api/v1/home/user/list",
		ResBodyType:         "json",
		Message:             "auto",
		Method:              "POST",
		ReqBodyIsJsonSchema: true,
		ResBodyIsJsonSchema: true,
	}

	rBody := &Request{
		Name: "zhangsan",
		Ids:  []int64{1, 3, 4},
	}

	// rv := reflect.ValueOf(rBody).Interface()

	s := &jsonschema.Document{}
	s.Read(rBody)

	// fmt.Println(s)
	req.ReqBodyOther = s.String()

	s2 := &jsonschema.Document{}
	s2.Read(Resp{})
	req.ResBody = s2.String()

	ctx := context.Background()
	yapi := NewYapi(api, token, nil)
	resp, err := yapi.SaveInterface(ctx, req)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	t.Logf("%+v", resp)

}

func Test_yapi_GetProject(t *testing.T) {
	ctx := context.Background()
	yapi := NewYapi(api, token, nil)
	project, err := yapi.GetProject(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", project)
}

func Test_yapi_GetCatMenu(t *testing.T) {
	ctx := context.Background()
	yapi := NewYapi(api, token, nil)

	project, err := yapi.GetProject(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	catList, err := yapi.GetCatMenu(ctx, project.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", catList)
}

func Test_yapi_AddCat(t *testing.T) {
	ctx := context.Background()

	yapi := NewYapi(api, token, nil)

	project, err := yapi.GetProject(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	catList, err := yapi.GetCatMenu(ctx, project.ID)
	if err != nil {
		t.Error(err)
		return
	}

	cateNameMap := make(map[string]bool)
	for _, v := range catList {
		cateNameMap[v.Name] = true
	}

	if _, ok := cateNameMap["test"]; ok {
		return
	}

	req := &CatAddReq{
		Desc:      "",
		Name:      "test",
		ProjectId: project.ID,
	}

	cat, err := yapi.AddCat(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", cat)
}

// 获取公共分类id
func getBaseCatId() (catId int, err error) {
	ctx := context.Background()
	yapi := NewYapi(api, token, nil)

	project, err := yapi.GetProject(ctx)
	if err != nil {
		return
	}

	catList, err := yapi.GetCatMenu(ctx, project.ID)
	if err != nil {
		return
	}

	for _, v := range catList {
		if v.Name == "公共分类" {
			return v.ID, nil
		}
	}

	return 0, fmt.Errorf("base cat is not found")
}
