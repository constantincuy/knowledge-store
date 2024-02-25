package files

import (
	"bufio"
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
	"github.com/constantincuy/knowledgestore/internal/core/util"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
)

type SearchFilesReq struct {
	KnowledgeBase string
	Query         string
}

type SearchFilesRes struct {
	Files []FileResult
}

type FileResult struct {
	Meta    file.File
	Content []string
}

func (s Service) Search(ctx context.Context, req SearchFilesReq) (SearchFilesRes, error) {
	knowledgeBase, err := knowledgebase.NewName(req.KnowledgeBase)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrBadRequest, err)
	}

	data := make([]string, 1)
	data[0] = req.Query

	embed, err := s.embedding.Extract(data)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	files, err := s.fileRepo.Search(ctx, knowledgeBase, embed.Vectors)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	dir, _ := os.Getwd()
	downloadPath := path.Join(dir, "download", "fake", s.storage.Provider())
	_ = os.MkdirAll(downloadPath, os.ModePerm)

	result := make([]FileResult, len(files))
	for i, f := range files {
		downloadFile := path.Join(downloadPath, req.KnowledgeBase+"_"+uuid.UUID(f.Id).String()+".txt")
		downloaded, err := os.Create(downloadFile)

		if err != nil {
			log.Println(err)
		}

		if err == nil {
			s.storage.DownloadDocument(context.Background(), f.Path, downloaded)

			chunksToLoad := util.Map(f.Chunks, func(t document.Document) int {
				return int(t.Chunk)
			})

			df, err := os.Open(downloadFile)
			chunked, err := document.NewChunkedDocument(bufio.NewReader(df), 256)
			df.Close()
			if err != nil {
				log.Println(err)
			}

			content := util.Map(chunksToLoad, func(c int) string {
				return chunked.GetChunk(c)
			})

			result[i] = FileResult{
				Meta:    f,
				Content: content,
			}
		}

		downloaded.Close()
	}

	return SearchFilesRes{result}, nil
}
