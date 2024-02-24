package repository

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type FsMockRepository struct {
	Files []string
}

func (repo *FsMockRepository) Ext(path string) string {
	return filepath.Ext(path)
}

func (repo *FsMockRepository) IsExist(path string) bool {
	return slices.Contains(repo.Files, path)
}

func (repo *FsMockRepository) IsDir(path string) (bool, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/workdir/" + path
	}
	result := new(bool)
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			if filepath == path {
				// this is file. not dir.
				if *result != true {
					*result = false
				}
				continue
			}
			*result = true
			continue
		}
	}
	if result != nil {
		return *result, nil
	}

	return false, fmt.Errorf("file or dir does not exist.")
}

func (repo *FsMockRepository) Create(path string, body []byte) error {
	return nil
}

func (repo *FsMockRepository) HomeDir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) WorkDir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	repo.Files = append(repo.Files, path)
	return nil
}

func (repo *FsMockRepository) Remove(path string) error {
	list := make([]string, 0)
	for _, file := range repo.Files {
		if file != path {
			list = append(list, file)
		}
	}
	repo.Files = list
	return nil
}

func (repo *FsMockRepository) Read(path string) ([]byte, error) {
	return make([]byte, 0), nil
}

func (repo *FsMockRepository) ListDirs(path string) ([]string, error) {
	return make([]string, 0), fmt.Errorf("not implemented")
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	list := make([]string, 0)
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			rest := strings.TrimPrefix(filepath, path+"/")
			if !strings.Contains(rest, "/") {
				list = append(list, filepath)
			}
		}
	}

	return list, nil
}
