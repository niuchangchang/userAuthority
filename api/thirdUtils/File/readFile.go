package File

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 读取文件
func (f *FileOperate) ReadFile() string {
	var result string
	var fileSrc = f.FilePath + "/" + f.FileName + "." + f.FileType
	if fileObj, err := os.Open(fileSrc); err == nil {
		defer fileObj.Close()
		if contents,err := ioutil.ReadAll(fileObj); err == nil {
			result = strings.Replace(string(contents),"\n","",1)
			return result
		}
	}
	return result
}

// 判断文件夹存在 (不存在, 建）
func (f *FileOperate) IsDirExist() bool {
	var  tmpDir = f.FilePath
	_, err := os.Stat(tmpDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
				log.Printf("创建 文件夹 %s 失败![%v]\n", tmpDir, err)
				return false
			}
		}
	}
	return true
}

// 判断文件是否存在 （不存在, 建)
func (f *FileOperate) IsFileExist() bool {
	var fileSrc = f.FilePath + "/" + f.FileName + "." + f.FileType
	_, err := os.Stat(fileSrc)
	if err != nil {
		if os.IsNotExist(err) {
			if fileObj , err := os.Create(fileSrc); err != nil {
				log.Printf("创建 文件 %s  失败![%v]\n", fileSrc, err)
				defer fileObj.Close()
				return false
			}
		}
	}
	return true
}

/**
 * @Author: STONE
 * @Description: 一键返回文件内容:  判断目录 -> 创建目录 -> 判断文件 -> 创建文件 -> 读文件
 * @Date: 2019-12-20 10:15
*/
func ReadFileFlow(f FileOperates) string {
	var t string
	if isDirExist := f.IsDirExist(); isDirExist {
		if isFileExist := f.IsFileExist(); isFileExist {
			t = f.ReadFile()
			return t
		}
	}
	return t
}