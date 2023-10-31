package models

import "go-admin/common/models"

type SysCustomer struct {
	models.Model
	ExpressDate    string `json:"expressDate"    gorm:"size:64;comment:日期"`
	CustomerName   string `json:"customerName"   gorm:"size:128;comment:客户姓名"`  //
	PhoneNumber    string `json:"phoneNumber"    gorm:"size:128;comment:联系电话"`  //
	Address        string `json:"address"        gorm:"size:255;comment:地址"`    //
	OperatorName   string `json:"operatorName"   gorm:"size:128;comment:业务员姓名"` //
	MachineNumber  string `json:"machineNumber"  gorm:"size:64;comment:机具号"`
	Issue          string `json:"issue"          gorm:"size:255;comment:处理问题"` //
	ExpressNumber  string `json:"expressNumber"  gorm:"size:64;comment:运单号"`   //
	ExpressResult  string `json:"expressResult"  gorm:"size:255;comment:签收情况"` //
	RegisterResult string `json:"registerResult" gorm:"size:255;comment:注册情况"` //
	UseResult      string `json:"useResult"      gorm:"size:255;comment:使用情况"` //
	ModelName      string `json:"modelName"      gorm:"size:64;comment:品牌型号"`  //
	SheetName      string `json:"sheetName"      gorm:"size:64;comment:分表名"`   //
	FileName       string `json:"fileName"       gorm:"size:64;comment:文件名"`   //
	Remark         string `json:"remark"         gorm:"size:255;comment:备注"`   //
	models.ControlBy
	models.ModelTime
}

func (*SysCustomer) TableName() string {
	return "sys_customer"
}
