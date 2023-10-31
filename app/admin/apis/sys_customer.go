package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/google/uuid"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"strings"
)

type SysCustomer struct {
	api.Api
}

// GetPage 获取配置管理列表
// @Summary 获取配置管理列表
// @Description 获取配置管理列表
// @Tags 配置管理
// @Param customerName query string false "名称"
// @Param customerKey query string false "key"
// @Param customerType query string false "类型"
// @Param isFrontend query int false "是否前端"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysApi}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-customer [get]
// @Security Bearer
func (e SysCustomer) GetPage(c *gin.Context) {
	s := service.SysCustomer{}
	req := dto.SysCustomerGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysCustomer, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取配置管理
// @Summary 获取配置管理
// @Description 获取配置管理
// @Tags 配置管理
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SysCustomer} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-customer/{id} [get]
// @Security Bearer
func (e SysCustomer) Get(c *gin.Context) {
	req := dto.SysCustomerGetReq{}
	s := service.SysCustomer{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysCustomer

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建配置管理
// @Summary 创建配置管理
// @Description 创建配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCustomerControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "创建成功"}"
// @Router /api/v1/sys-customer [post]
// @Security Bearer
func (e SysCustomer) Insert(c *gin.Context) {
	s := service.SysCustomer{}
	req := dto.SysCustomerControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改配置管理
// @Summary 修改配置管理
// @Description 修改配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCustomerControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-customer/{id} [put]
// @Security Bearer
func (e SysCustomer) Update(c *gin.Context) {
	s := service.SysCustomer{}
	req := dto.SysCustomerControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除配置管理
// @Summary 删除配置管理
// @Description 删除配置管理
// @Tags 配置管理
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-customer [delete]
// @Security Bearer
func (e SysCustomer) Delete(c *gin.Context) {
	s := service.SysCustomer{}
	req := dto.SysCustomerDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Get2SysApp 获取系统配置信息
// @Summary 获取系统前台配置信息，主要注意这里不在验证权限
// @Description 获取系统配置信息，主要注意这里不在验证权限
// @Tags 配置管理
// @Success 200 {object} response.Response{data=map[string]string} "{"code": 200, "data": [...]}"
// @Router /api/v1/app-customer [get]
//func (e SysCustomer) Get2SysApp(c *gin.Context) {
//	req := dto.SysCustomerGetToSysAppReq{}
//	s := service.SysCustomer{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.Form).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		return
//	}
//	// 控制只读前台的数据
//	req.IsFrontend = "1"
//	list := make([]models.SysCustomer, 0)
//	err = s.GetWithKeyList(&req, &list)
//	if err != nil {
//		e.Error(500, err, "查询失败")
//		return
//	}
//	mp := make(map[string]string)
//	for i := 0; i < len(list); i++ {
//		key := list[i].PhoneNumber
//		if key != "" {
//			mp[key] = list[i].PhoneNumber
//		}
//	}
//	e.OK(mp, "查询成功")
//}

// Get2Set 获取配置
// @Summary 获取配置
// @Description 界面操作设置配置值的获取
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Success 200 {object} response.Response{data=map[string]interface{}}	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/set-customer [get]
// @Security Bearer
//func (e SysCustomer) Get2Set(c *gin.Context) {
//	s := service.SysCustomer{}
//	req := make([]dto.GetSetSysCustomerReq, 0)
//	err := e.MakeContext(c).
//		MakeOrm().
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	err = s.GetForSet(&req)
//	if err != nil {
//		e.Error(500, err, "查询失败")
//		return
//	}
//	m := make(map[string]interface{}, 0)
//	for _, v := range req {
//		m[v.PhoneNumber] = v.PhoneNumber
//	}
//	e.OK(m, "查询成功")
//}

// Update2Set 设置配置
// @Summary 设置配置
// @Description 界面操作设置配置值
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body []dto.GetSetSysCustomerReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/set-customer [put]
// @Security Bearer
//func (e SysCustomer) Update2Set(c *gin.Context) {
//	s := service.SysCustomer{}
//	req := make([]dto.GetSetSysCustomerReq, 0)
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.JSON).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	err = s.UpdateForSet(&req)
//	if err != nil {
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	e.OK("", "更新成功")
//}

// GetSysCustomerByKEYForService 根据Key获取SysCustomer的Service
// @Summary 根据Key获取SysCustomer的Service
// @Description 根据Key获取SysCustomer的Service
// @Tags 配置管理
// @Param customerKey path string false "customerKey"
// @Success 200 {object} response.Response{data=dto.SysCustomerByKeyReq} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-customer/{id} [get]
// @Security Bearer
//func (e SysCustomer) GetSysCustomerByKEYForService(c *gin.Context) {
//	var s = new(service.SysCustomer)
//	var req = new(dto.SysCustomerByKeyReq)
//	var resp = new(dto.GetSysCustomerByKEYForServiceResp)
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(req, nil).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	err = s.GetWithKey(req, resp)
//	if err != nil {
//		e.Error(500, err, err.Error())
//		return
//	}
//	e.OK(resp, s.Msg)
//}

// UploadCustomerFile
// @Summary 上传客户信息文件
// @Description 获取JSON
// @Tags 个人中心
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/avatar [post]
// @Security Bearer
func (e SysCustomer) UploadCustomerFile(c *gin.Context) {
	s := service.SysCustomer{}
	//req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	form, _ := c.MultipartForm()
	files := form.File["attachFile[]"]
	filenames := form.Value["filename"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid
	for i, file := range files {
		e.Logger.Debugf("upload avatar file: %s - filename : %s", file.Filename, filenames[i])
		// 上传文件至指定目录
		filePath := filPath + filenames[i]
		if !(strings.HasSuffix(filePath, ".xlsx") || strings.HasSuffix(filePath, ".xls")) {
			filePath += ".xlsx"
		}
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			e.Logger.Errorf("save file error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
		err := s.ParseCustomerFile(file, p, user.GetUserId(c))
		if err != nil {
			e.Logger.Errorf("parse file error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
	}
	//req.UserId = p.UserId
	//req.Avatar = "/" + filPath
	//
	//err = s.UpdateAvatar(&req, p)
	//if err != nil {
	//	e.Logger.Error(err)
	//	return
	//}
	e.OK(filPath, "保存成功")
}
