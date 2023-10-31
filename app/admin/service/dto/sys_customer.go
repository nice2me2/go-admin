package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

// SysCustomerGetPageReq 列表或者搜索使用结构体
type SysCustomerGetPageReq struct {
	dto.Pagination `search:"-"`
	//CustomerName   string `form:"customerName" search:"type:contains;column:customer_name;table:sys_customer"`
	//CustomerKey    string `form:"customerKey" search:"type:contains;column:customer_key;table:sys_customer"`
	//CustomerType   string `form:"customerType" search:"type:exact;column:customer_type;table:sys_customer"`
	//IsFrontend     string `form:"isFrontend" search:"type:exact;column:is_frontend;table:sys_customer"`
	CustomerName  string `form:"customerName"  json:"customerName"  search:"type:contains;column:customer_name;table:sys_customer"`  //"客户姓名"`  //
	ExpressDate   string `form:"expressDate"   json:"expressDate"   search:"type:contains;column:express_date;table:sys_customer"`   //"日期"`  //
	PhoneNumber   string `form:"phoneNumber"   json:"phoneNumber"   search:"type:contains;column:phone_number;table:sys_customer"`   //"联系电话"`  //
	OperatorName  string `form:"operatorName"  json:"operatorName"  search:"type:contains;column:operator_name;table:sys_customer"`  //"业务员姓名"` //
	MachineNumber string `form:"machineNumber" json:"machineNumber" search:"type:contains;column:machine_number;table:sys_customer"` //"运单号"`   //
	SheetName     string `form:"sheetName"     json:"sheetName"     search:"type:contains;column:sheet_name;table:sys_customer"`     //"运单号"`   //
	ExpressNumber string `form:"expressNumber" json:"expressNumber" search:"type:contains;column:express_number;table:sys_customer"` //"运单号"`   //
	SysCustomerOrder
}

type SysCustomerOrder struct {
	IdOrder           string `search:"type:order;column:id;table:sys_customer" form:"idOrder"`
	CustomerNameOrder string `search:"type:order;column:customer_name;table:sys_customer" form:"customerNameOrder"`
	CustomerKeyOrder  string `search:"type:order;column:customer_key;table:sys_customer" form:"customerKeyOrder"`
	CustomerTypeOrder string `search:"type:order;column:customer_type;table:sys_customer" form:"customerTypeOrder"`
	CreatedAtOrder    string `search:"type:order;column:created_at;table:sys_customer" form:"createdAtOrder"`
}

func (m *SysCustomerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysCustomerGetToSysAppReq struct {
	IsFrontend string `form:"isFrontend" search:"type:exact;column:is_frontend;table:sys_customer"`
}

func (m *SysCustomerGetToSysAppReq) GetNeedSearch() interface{} {
	return *m
}

// SysCustomerControl 增、改使用的结构体
type SysCustomerControl struct {
	Id int `uri:"Id" comment:"编码"` // 编码
	//CustomerName  string `json:"customerName" comment:""`
	//CustomerKey   string `uri:"customerKey" json:"customerKey" comment:""`
	//CustomerValue string `json:"customerValue" comment:""`
	//CustomerType  string `json:"customerType" comment:""`
	//IsFrontend    string `json:"isFrontend"`
	//Remark        string `json:"remark" comment:""`
	ExpressDate    string `json:"expressDate"    comment:""` //"查询日期"`
	CustomerName   string `json:"customerName"   comment:""` //"客户姓名"`  //
	PhoneNumber    string `json:"phoneNumber"    comment:""` //"联系电话"`  //
	Address        string `json:"address"        comment:""` //"地址"`  //
	OperatorName   string `json:"operatorName"   comment:""` //"业务员姓名"` //
	MachineNumber  string `json:"machineNumber"  comment:""` //"机具号"`
	Issue          string `json:"issue"          comment:""` //"处理问题"` //
	ExpressNumber  string `json:"expressNumber"  comment:""` //"运单号"`   //
	ExpressResult  string `json:"expressResult"  comment:""` //"签收情况"` //
	RegisterResult string `json:"registerResult" comment:""` //"注册情况"` //
	UseResult      string `json:"useResult"      comment:""` //"使用情况"` //
	ModelName      string `json:"modelName"      comment:""` //"品牌型号"` //
	SheetName      string `json:"sheetName"      comment:""` //"分表名"` //
	Remark         string `json:"remark"         comment:""` //"备注"` //
	common.ControlBy
}

// Generate 结构体数据转化 从 SysCustomerControl 至 system.SysCustomer 对应的模型
func (s *SysCustomerControl) Generate(model *models.SysCustomer) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ExpressDate = s.ExpressDate
	model.CustomerName = s.CustomerName
	model.PhoneNumber = s.PhoneNumber
	model.OperatorName = s.OperatorName
	model.MachineNumber = s.MachineNumber
	model.Address = s.Address
	model.Issue = s.Issue
	model.ExpressNumber = s.ExpressNumber
	model.ExpressResult = s.ExpressResult
	model.RegisterResult = s.RegisterResult
	model.UseResult = s.UseResult
	model.ModelName = s.ModelName
	model.SheetName = s.SheetName
	model.Remark = s.Remark

}

// GetId 获取数据对应的ID
func (s *SysCustomerControl) GetId() interface{} {
	return s.Id
}

// GetSetSysCustomerReq 增、改使用的结构体
type GetSetSysCustomerReq struct {
	CustomerName string `json:"customerName" comment:""`
	PhoneNumber  string `json:"customerValue" comment:""`
}

// Generate 结构体数据转化 从 SysCustomerControl 至 system.SysCustomer 对应的模型
func (s *GetSetSysCustomerReq) Generate(model *models.SysCustomer) {
	model.CustomerName = s.CustomerName
	model.PhoneNumber = s.PhoneNumber
}

type UpdateSetSysCustomerReq map[string]string

// SysCustomerByKeyReq 根据Key获取配置
type SysCustomerByKeyReq struct {
	CustomerKey string `uri:"customerKey" search:"type:contains;column:customer_key;table:sys_customer"`
}

func (m *SysCustomerByKeyReq) GetNeedSearch() interface{} {
	return *m
}

type GetSysCustomerByKEYForServiceResp struct {
	//CustomerKey   string `json:"customerKey" comment:""`
	//CustomerValue string `json:"customerValue" comment:""`
	ExpressDate    time.Time `json:"createdAt"      comment:""` //"查询日期"`
	CustomerName   string    `json:"customerName"   comment:""` //"客户姓名"`  //
	PhoneNumber    string    `json:"phoneNumber"    comment:""` //"联系电话"`  //
	OperatorName   string    `json:"operatorName"   comment:""` //"业务员姓名"` //
	MachineNumber  string    `json:"machineNumber"  comment:""` //"机具号"`
	Issue          string    `json:"issue"          comment:""` //"处理问题"` //
	ExpressNumber  string    `json:"expressNumber"  comment:""` //"运单号"`   //
	ExpressResult  string    `json:"expressResult"  comment:""` //"签收情况"` //
	RegisterResult string    `json:"registerResult" comment:""` //"注册情况"` //
	UseResult      string    `json:"useResult"      comment:""` //"使用情况"` //
	Remark         string    `json:"remark"         comment:""` //"备注"` //
}

type SysCustomerGetReq struct {
	Id int `uri:"id"`
}

func (s *SysCustomerGetReq) GetId() interface{} {
	return s.Id
}

type SysCustomerDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysCustomerDeleteReq) GetId() interface{} {
	return s.Ids
}
