package goyapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/levigross/grequests"
	"moul.io/http2curl"
)

// @refer https://hellosean1025.github.io/yapi/openapi.html

type yapi struct {
	token   string
	address string
	log     Logger
}

func NewYapi(address, token string, logger Logger) (i *yapi) {
	api := &yapi{
		address: address,
		token:   token,
		log:     logger,
	}

	if api.log == nil {
		api.log = log.Default()
	}

	return api
}

func (y *yapi) AddCat(ctx context.Context, req *CatAddReq) (reply *Cat, err error) {
	req.Token = y.token

	resp, err := y.post(ctx, "/api/interface/add_cat", req)

	err = y.resolveResp(resp, err, &reply)
	return
}

func (y *yapi) GetCatMenu(ctx context.Context, projectId int) (reply CatList, err error) {
	resp, err := y.get(ctx, "/api/interface/getCatMenu", map[string]string{"token": y.token})

	err = y.resolveResp(resp, err, &reply)
	return

}

func (y *yapi) GetProject(ctx context.Context) (reply *Project, err error) {
	resp, err := y.get(ctx, "/api/project/get", map[string]string{"token": y.token})

	err = y.resolveResp(resp, err, &reply)
	return
}

// 新增或修改接口
func (y *yapi) SaveInterface(ctx context.Context, req *InterfaceSaveReq) (id int, err error) {
	req.Token = y.token

	resp, err := y.post(ctx, "/api/interface/save", req)

	var reply InterfaceSaveResp

	err = y.resolveResp(resp, err, &reply)

	if err != nil {
		return 0, err
	}

	if len(reply) > 0 {
		return reply[0].ID, nil
	}

	return 0, nil
}

func (y *yapi) AddInterface(ctx context.Context, req *InterfaceAddReq) (exists bool, id int, err error) {
	req.Token = y.token

	resp, err := y.post(ctx, "/api/interface/add", req)

	var reply InterfaceDetail

	err = y.resolveResp(resp, err, &reply)

	if err != nil {
		if v, ok := err.(Resp); ok {
			if v.ErrCode == 40022 {
				return true, 0, nil
			}
		}
		return false, 0, err
	}

	return false, reply.ID, nil
}

func (y *yapi) post(ctx context.Context, api string, req interface{}) (*grequests.Response, error) {
	url := fmt.Sprintf("%s/%s", y.address, strings.TrimLeft(api, "/"))

	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    req,
		Context: ctx,
		// Headers: map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		BeforeRequest: func(req *http.Request) error {
			cmd, _ := http2curl.GetCurlCommand(req)
			y.log.Printf("%s\n", cmd)
			return nil
		},
	})

	return resp, err
}

func (y *yapi) get(ctx context.Context, api string, params map[string]string) (*grequests.Response, error) {
	url := fmt.Sprintf("%s/%s", y.address, strings.TrimLeft(api, "/"))

	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Context: ctx,
		Params:  params,
		BeforeRequest: func(req *http.Request) error {
			cmd, _ := http2curl.GetCurlCommand(req)
			y.log.Printf("%s\n", cmd)
			return nil
		},
	})

	return resp, err
}

func (y *yapi) resolveResp(resp *grequests.Response, err error, reply interface{}) error {
	if err != nil {
		return err
	}
	defer resp.Close()

	var baseResp Resp
	if reply != nil {
		baseResp.Data = &reply
	}

	err = resp.JSON(&baseResp)
	if err != nil {
		return err
	}

	if baseResp.ErrCode == 0 {
		return nil
	}

	return baseResp
}
