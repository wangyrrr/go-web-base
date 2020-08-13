package service

import (
	"go-web-base/core/bo"
	"go-web-base/core/entity"
	"go-web-base/global"
	"go-web-base/global/response"
)

func GetDemo(pageInfo bo.PageInfo) response.PageResult {
	db := global.DB.Model(&entity.Demo{})
	var total int
	db.Count(&total)
	pages :=  (total + pageInfo.Size - 1) / pageInfo.Size
	records := make([]entity.Demo, 0)
	db.Limit(pageInfo.Size).Offset((pageInfo.Page - 1) * pageInfo.Size).Find(&records)
	return response.PageResult{
		Total: total,
		Size: pageInfo.Size,
		Pages: pages,
		Current: pageInfo.Page,
		Records: records,
	}
}
