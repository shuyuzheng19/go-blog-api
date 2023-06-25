package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	db *gorm.DB
}

func (f FileRepositoryImpl) FindFileByPage(uid int, page int, keyword string, sort string) (files []models.FileInfo, count int64) {
	var query = f.db.Model(&models.FileInfo{})

	if uid == 0 {
		query.Where("is_public = true")
	} else {
		query.Where("user_id = ?", uid)
	}

	if keyword != "" {
		query.Where("old_name like ?", "%"+keyword+"%")
	}

	if err := query.Offset((page - 1) * common.FILE_PAGE_COUNT).Limit(common.FILE_PAGE_COUNT).Count(&count).Error; err != nil || count == 0 {
		return make([]models.FileInfo, 0), 0
	}

	query.Order(sort + " desc").Preload("FileMd5").Find(&files)

	return files, count
}

func (f FileRepositoryImpl) SaveFileInfo(fileInfo models.FileInfo) error {
	return f.db.Model(&models.FileInfo{}).Save(&fileInfo).Error
}

func (f FileRepositoryImpl) FindMD5(md5 string) string {
	var result string

	var sql = "select url from file_md5 where md5 = ? limit 1"

	f.db.Model(&models.FileMd5{}).Raw(sql, md5).Scan(&result)

	return result
}

func (f FileRepositoryImpl) SaveMd5(md5 models.FileMd5) error {
	return f.db.Model(&models.FileMd5{}).Create(&md5).Error
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return FileRepositoryImpl{db: db}
}
