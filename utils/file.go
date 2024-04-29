package utils

import (
	"bufio"
	"io"
	"os"
)

// 如果目标目录不存在，创建目录
func MakeDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm) //创建多级目录
		return err
	}

	return err
}

// 按行读取文本文件
func ReadTxtLines(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	l := make([]string, 0)
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		l = append(l, string(line))
	}

	return l, nil
}
