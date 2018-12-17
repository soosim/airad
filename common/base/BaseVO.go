package base

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/context"
)

type VO interface {
}

type BaseVO struct {
}

type BaseListRequestVO struct {
	BaseVO
	Page int `validate:"required,max=1000,min=1"`
	Size int `validate:"required,max=1000,min=1"`
}

type BaseListResponseVO struct {
	BaseVO
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
}

func ParseJsonRequestToVO(ctx *context.Context, vo interface{}) error {
	if err := json.Unmarshal(ctx.Input.RequestBody, vo); err != nil {
		ctx.Output.JSON(ErrInputData, false, false)
		return errors.New("json参数解析异常")
	}
	return nil
}
