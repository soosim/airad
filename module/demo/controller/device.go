package controller

import (
	"airad/common/base"
	"airad/common/util"
	"airad/module/demo/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//  DeviceController operations for Device
type DeviceController struct {
	base.BaseController
}

// URLMapping ...
func (c *DeviceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Device
// @Param	body		body 	models.Device	true		"body for Device content"
// @Success 201 {int} models.Device
// @Failure 403 body is empty
// @router / [post]
func (c *DeviceController) Post() {
	var v models.Device
	token := c.Ctx.Input.Header("token")
	//id := c.Ctx.Input.Header("id")
	et := util.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = base.BaseResponse{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := util.CheckNewDevicePost(v.UserId, v.DeviceName,
			v.Address, v.Status, v.Latitude, v.Longitude); errorMessage != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, errorMessage, ""}
			c.ServeJSON()
			return
		}
		if models.CheckDeviceName(v.DeviceName) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "设备名已经注册了", ""}
			c.ServeJSON()
			return
		}

		if !models.CheckUserId(v.UserId) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "用户ID不存在", ""}
			c.ServeJSON()
			return
		}

		if !models.CheckUserIdAndToken(v.UserId, token) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "用户ID和Token不匹配", ""}
			c.ServeJSON()
			return
		}

		if deviceId, err := models.AddDevice(&v); err == nil {
			if user, err := models.GetUserById(v.UserId); err == nil {
				models.UpdateUserDeviceCount(user)
				c.Ctx.Output.SetStatus(201)
				var returnData = &CreateObjectData{int(deviceId)}
				c.Data["json"] = &base.BaseResponse{0, 0, "ok", returnData}
			} else {
				c.Ctx.ResponseWriter.WriteHeader(403)
				c.Data["json"] = base.BaseResponse{403, 403, "用户Id不存在", ""}
				c.ServeJSON()
				return
			}

			c.Ctx.Output.SetStatus(201)
			var returnData = &CreateObjectData{int(deviceId)}
			c.Data["json"] = &base.BaseResponse{0, 0, "ok", returnData}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = &base.BaseResponse{1, 1, "设备名注册失败", err.Error()}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Device by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Device
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DeviceController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDeviceById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Device
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Device
// @Failure 403
// @router / [get]
func (c *DeviceController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 10
	var offset int
	var userId int

	token := c.Ctx.Input.Header("token")
	//id := c.Ctx.Input.Header("id")
	et := util.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(token)
	if !valido {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = base.BaseResponse{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	if found, user := models.GetUserByToken(token); !found {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = &base.BaseResponse{401, 401, "未找到相关的用户", ""}
		c.ServeJSON()
		return
	} else {
		userId = user.Id
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, totalCount, err := models.GetAllDevices(query, fields, sortby, order, offset, limit, userId)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var returnData = &GetDeviceData{totalCount, l}
		c.Data["json"] = &base.BaseResponse{0, 0, "ok", returnData}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Device
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Device	true		"body for Device content"
// @Success 200 {object} models.Device
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DeviceController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Device{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateDeviceById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Device
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DeviceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDevice(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Get All Devices by User Id
// @Param	userId		path 	int	true
// @Param	limit		path 	int	true
// @Param	offset		path 	int	true
// @Param	fields		path 	String	true
// @Param	token		path 	String	true
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [post]
func (c *DeviceController) GetDevicesByUserId() {
	var limit int = 10
	var offset int
	var fields []string
	var v models.DeviceRequestStruct

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := util.CheckUserDevicePost(v.UserId, v.Limit, v.Offset); errorMessage != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, errorMessage, ""}
			c.ServeJSON()
			return
		}
		if v := v.Fields; v != "" {
			fields = strings.Split(v, ",")
		} else {
			fields = strings.Split("DeviceName,Address,Status,Latitude,Longitude", ",")
		}
		// limit: 10 (default is 10)
		if para := v.Limit; para != 0 {
			limit = para
		}
		// offset: 0 (default is 0)
		if para := v.Offset; para != 0 {
			offset = para
		}
		if !models.CheckUserId(v.UserId) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "用户ID不存在", ""}
			c.ServeJSON()
			return
		}
		if devices, err := models.GetDevicesByUserId(v.UserId, fields, limit, offset); err == nil {
			c.Data["json"] = devices
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = &base.BaseResponse{1, 1, "设备结构解析失败", err.Error()}
	}
	c.ServeJSON()
}
