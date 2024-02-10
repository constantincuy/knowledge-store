package files

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
)

type GetFilesystemReq struct {
	KnowledgeBase string
	Provider      string
}

type GetFilesystemRes struct {
	Filesystem file.Filesystem
}

func (s Service) GetFilesystem(ctx context.Context, req GetFilesystemReq) (GetFilesystemRes, error) {
	name, err := knowledgebase.NewName(req.KnowledgeBase)

	if err != nil {
		return GetFilesystemRes{}, service.NewError(service.ErrBadRequest, err)
	}

	provider, err := file.NewProvider(req.Provider)

	if err != nil {
		return GetFilesystemRes{}, service.NewError(service.ErrBadRequest, err)
	}

	files, err := s.fileRepo.GetAllProviderFiles(ctx, name, provider)
	if err != nil {
		return GetFilesystemRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	fileList, err := file.NewList(files)
	if err != nil {
		return GetFilesystemRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	return GetFilesystemRes{file.NewFilesystem(fileList)}, nil
}
