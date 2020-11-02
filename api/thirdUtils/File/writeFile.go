package File

import (
	"log"
	"os"
)

// 写文件
func (f *FileOperate) WriteFile() bool {

	var fileSrc = f.FilePath + "/" + f.FileName + "." + f.FileType
	var bytes = []byte(f.FileContent)

	if file, err := os.OpenFile(fileSrc, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666); err != nil {
		log.Fatalf("打开文件: %s, 失败%s", fileSrc,  err)
		return false
	} else {
		if bw, err := file.Write(bytes);err != nil {
			log.Fatalf("打开文件: %s, 失败%s", fileSrc,  err)
			return false
		} else {
			log.Printf("写入 %d bytes. => 内容: %s ", bw, string(bytes))
		}
		defer file.Close()
		return true
	}
}

/**
 * @Author: STONE
 * @Description: 一键写文件内容 -> 判断目录 -> 创建目录 -> 判断文件 -> 创建文件 -> 写文件
 * @Date: 2019-12-20 10:15
 */
func WriteFileFlow(f FileOperates) bool {
	var t bool
	if isDirExist := f.IsDirExist(); isDirExist {
		if isFileExist := f.IsFileExist(); isFileExist {
			t = f.WriteFile()
			return t
		}
	}
	return t
}