package zip

import (
	"archive/zip"
	"context"
	"fmt"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"io"
	"log"
	"os"
)

type Storage struct {
	zipPath string
}

func (s Storage) Provider() string {
	return "zip-storage"
}

func (s Storage) GetChangedDocuments(ctx context.Context, filesystem file.Filesystem) (file.ChangeList, error) {
	read, err := zip.OpenReader(s.zipPath)
	if err != nil {
		log.Printf("Could not open zip file: %s\n", err)
		return file.ChangeList{}, err
	}
	defer read.Close()

	files := make([]file.File, len(read.File))
	for i, zf := range read.File {
		df, err := s.zipFileEntryToDomainFile(zf)

		if err == nil {
			files[i] = df
		} else {
			log.Printf("Coould not convert zip file %s error: %s\n", zf.Name, err)
		}
	}

	fl, err := file.NewList(files)

	if err != nil {
		log.Printf("Could not create file list: %s\n", err)
		return file.ChangeList{}, err
	}

	changeList, err := filesystem.Sync(fl)

	if err != nil {
		log.Printf("Could not create change list: %s\n", err)
		return file.ChangeList{}, err
	}

	return changeList, nil
}

func (s Storage) zipFileEntryToDomainFile(zf *zip.File) (file.File, error) {
	path, err := file.NewPath(zf.Name)
	if err != nil {
		return file.File{}, err
	}

	provider, err := file.NewProvider(s.Provider())
	if err != nil {
		return file.File{}, err
	}

	created, err := file.NewCreated(zf.Modified)
	if err != nil {
		return file.File{}, err
	}

	updated, err := file.NewUpdated(zf.Modified)
	if err != nil {
		return file.File{}, err
	}

	return file.NewFile(path, provider, created, updated), nil
}

func (s Storage) DownloadDocument(ctx context.Context, path file.Path, target *os.File) {
	read, err := zip.OpenReader(s.zipPath)
	if err != nil {
		log.Printf("Could not open zip file: %s\n", err)
		return
	}
	defer read.Close()

	var foundFile *zip.File = nil
	for _, f := range read.File {
		if f.Name == string(path) {
			foundFile = f
			break
		}
	}

	if foundFile != nil {
		v, err := foundFile.Open()
		if err != nil {
			log.Printf("Could not open file of zip: %s\n", err)
			return
		}
		defer v.Close()

		_, err = io.Copy(target, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func NewStorage(zipPath string) Storage {
	return Storage{zipPath}
}
