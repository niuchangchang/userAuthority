package File

//定义文件操作接口
type FileOperates interface {
	ReadFile() string
	WriteFile() bool
	IsDirExist() bool
	IsFileExist() bool
}

//定义文件操作结构体
type FileOperate struct {
	FilePath string
	FileName string
	FileType string
	FileContent string
}
