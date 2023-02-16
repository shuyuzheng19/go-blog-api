package service

import "vs-blog-api/response"

type FileService interface {
	GetCurrentFiles(userId int, page int, sort string, flag bool, keyword string) response.PageInfoResponse
	GetPublicFiles(page int, sort string, flag bool, keyword string) response.PageInfoResponse
	DeleteFiles(ids []string)
	DeleteUserFiles(userId int, ids []string)
}
