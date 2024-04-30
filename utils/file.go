package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// FileManager 结构体用于管理文件的读写操作
type FileManager struct {
	FileName string     // 文件名
	Mutex    sync.Mutex // 互斥锁
}

// NewFileManager 创建一个新的 FileManager 实例
func NewFileManager(filename string) *FileManager {
	return &FileManager{FileName: filename}
}

// WriteToFile 将内容写入文件
func (fm *FileManager) WriteToFile(data []byte) error {
	fm.Mutex.Lock()
	defer fm.Mutex.Unlock()

	return ioutil.WriteFile(fm.FileName, data, 0644)
}

// AppendToFile 将内容追加到文件
func (fm *FileManager) AppendToFile(data string) error {
	fm.Mutex.Lock()
	defer fm.Mutex.Unlock()

	f, err := os.OpenFile(fm.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return nil
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	_, err = writer.WriteString(data + "\n")
	return err
}

// ReadFromFile 从文件中读取内容
func (fm *FileManager) ReadFromFile() ([]byte, error) {
	fm.Mutex.Lock()
	defer fm.Mutex.Unlock()

	if _, err := os.Stat(fm.FileName); os.IsNotExist(err) {
		return nil, err //文件不存在
	}

	return os.ReadFile(fm.FileName)
}

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
