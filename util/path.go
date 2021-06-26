package util

import (
	"io/ioutil"
	"path/filepath"
)

//var ErrNotGitRepository = error.New("not found")

// path で指定したリポジトリのルートディレクトリを返す
func FindGitRoot(path string) (string, error) {
	files,err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range  files {
		if file.IsDir() && file.Name() == ".git" {
			return path, nil
		}
	}
	abs,err := filepath.Abs(path)
	if err != nil{
		return "", err
	}
	if abs == "/" {
		return "", err
	}

	return FindGitRoot(filepath.Join(path, "."))
}
