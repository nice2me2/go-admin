package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/xuri/excelize/v2"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"mime/multipart"
)

type SysCustomer struct {
	service.Service
}

// GetPage 获取SysCustomer列表
func (e *SysCustomer) GetPage(c *dto.SysCustomerGetPageReq, list *[]models.SysCustomer, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysCustomerPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysCustomer对象
func (e *SysCustomer) Get(d *dto.SysCustomerGetReq, model *models.SysCustomer) error {
	err := e.Orm.First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysCustomerPage error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("Service GetSysCustomer error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysCustomer对象
func (e *SysCustomer) Insert(c *dto.SysCustomerControl) error {
	var err error
	var data models.SysCustomer
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertSysCustomer error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysCustomer对象
func (e *SysCustomer) Update(c *dto.SysCustomerControl) error {
	var err error
	var model = models.SysCustomer{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(&model)
	err = db.Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysCustomer error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// SetSysCustomer 修改SysCustomer对象
//func (e *SysCustomer) SetSysCustomer(c *[]dto.GetSetSysCustomerReq) error {
//	var err error
//	for _, req := range *c {
//		var model = models.SysCustomer{}
//		e.Orm.Where("phone_number = ?", req.PhoneNumber).First(&model)
//		if model.Id != 0 {
//			req.Generate(&model)
//			db := e.Orm.Save(&model)
//			err = db.Error
//			if err != nil {
//				e.Log.Errorf("Service SetSysCustomer error:%s", err)
//				return err
//			}
//			if db.RowsAffected == 0 {
//				return errors.New("无权更新该数据")
//			}
//		}
//	}
//	return nil
//}

//func (e *SysCustomer) GetForSet(c *[]dto.GetSetSysCustomerReq) error {
//	var err error
//	var data models.SysCustomer
//
//	err = e.Orm.Model(&data).
//		Find(c).Error
//	if err != nil {
//		e.Log.Errorf("Service GetSysCustomerPage error:%s", err)
//		return err
//	}
//	return nil
//}

//func (e *SysCustomer) UpdateForSet(c *[]dto.GetSetSysCustomerReq) error {
//	m := *c
//	for _, req := range m {
//		var data models.SysCustomer
//		if err := e.Orm.Where("phone_number = ?", req.PhoneNumber).
//			First(&data).Error; err != nil {
//			e.Log.Errorf("Service GetSysCustomerPage error:%s", err)
//			return err
//		}
//		if data.PhoneNumber != req.PhoneNumber {
//			data.PhoneNumber = req.PhoneNumber
//
//			if err := e.Orm.Save(&data).Error; err != nil {
//				e.Log.Errorf("Service GetSysCustomerPage error:%s", err)
//				return err
//			}
//		}
//
//	}
//
//	return nil
//}

// Remove 删除SysCustomer
func (e *SysCustomer) Remove(d *dto.SysCustomerDeleteReq) error {
	var err error
	var data models.SysCustomer

	db := e.Orm.Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Service RemoveSysCustomer error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetWithKey 根据Key获取SysCustomer
//func (e *SysCustomer) GetWithKey(c *dto.SysCustomerByKeyReq, resp *dto.GetSysCustomerByKEYForServiceResp) error {
//	var err error
//	var data models.SysCustomer
//	err = e.Orm.Table(data.TableName()).Where("customer_key = ?", c.CustomerKey).First(resp).Error
//	if err != nil {
//		e.Log.Errorf("At Service GetSysCustomerByKEY Error:%s", err)
//		return err
//	}
//
//	return nil
//}
//
//func (e *SysCustomer) GetWithKeyList(c *dto.SysCustomerGetToSysAppReq, list *[]models.SysCustomer) error {
//	var err error
//	err = e.Orm.
//		Scopes(
//			cDto.MakeCondition(c.GetNeedSearch()),
//		).
//		Find(list).Error
//	if err != nil {
//		e.Log.Errorf("Service GetSysCustomerByKey error:%s", err)
//		return err
//	}
//	return nil
//}

// ParseCustomerFile 解析客户信息文件
func (e *SysCustomer) ParseCustomerFile(file *multipart.FileHeader, p *actions.DataPermission, createBy int) error {
	var (
		err error
		//data        models.SysCustomer
		sheetTitles = []string{"日期", "姓名", "电话", "联系地址", "业务员", "机具号", "品牌型号", "处理问题", "运单号", "签收情况", "注册情况", "刷卡情况", "备注"}
		//now         = time.Now()
		customers = make([]*models.SysCustomer, 0)
	)

	open, err := file.Open()
	if err != nil {
		return err
	}

	f, err := excelize.OpenReader(open)
	if err != nil {
		return err
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	for _, name := range sheetNames {
		rows, err := f.GetRows(name)
		if err != nil {
			return err
		}
		for i, row := range rows {
			if i == 0 {
				if len(row) < len(sheetTitles) {
					return errors.New("文件不规范，缺少字段")
				}
				for i2, title := range sheetTitles {
					if title != row[i2] {
						return errors.New("文件不规范，请检查字段顺序")
					}
				}
				continue
			}
			if len(row) < len(sheetTitles) {
				return errors.New("文件不规范，请检查字段顺序")
			}
			cust := &models.SysCustomer{
				ExpressDate:    replaceNullString(row[0]),
				CustomerName:   replaceNullString(row[1]),
				PhoneNumber:    replaceNullString(row[2]),
				Address:        replaceNullString(row[3]),
				OperatorName:   replaceNullString(row[4]),
				MachineNumber:  replaceNullString(row[5]),
				ModelName:      replaceNullString(row[6]),
				Issue:          replaceNullString(row[7]),
				ExpressNumber:  replaceNullString(row[8]),
				ExpressResult:  replaceNullString(row[9]),
				RegisterResult: replaceNullString(row[10]),
				UseResult:      replaceNullString(row[11]),
				SheetName:      name,
				FileName:       file.Filename,
				Remark:         replaceNullString(row[12]),
			}
			cust.CreateBy = createBy
			customers = append(customers, cust)
		}
	}

	if len(customers) == 0 {
		return errors.New("未找到有效数据")
	}

	err = e.Orm.CreateInBatches(customers, len(customers)).Error
	//err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertSysCustomer error:%s", err)
		return err
	}
	return nil
}

func replaceNullString(ss string) string {
	if len(ss) > 0 {
		return ss
	}
	return "<空>"
}
