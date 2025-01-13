package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1



// 先获取文件后缀筛出原始文件名进行 SHA256 加密, 然后返回经过加密处理后的文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeSHA256(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	// 循环查找”.“符号, 然后通过切片索引返回后缀
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

func CheckSavePath(dst string) bool {

	// 获取 FileInfo, 利用 os.Stat 返回的 error 值与系统定义的 oserror.ErrNotExist 进行校验
	_, err := os.Stat(dst)

	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext) // 将文件后缀统一转换为大写
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size < global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true 
		}
	}

	return false // 超出限制
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	// FileInfo 与 oserror.ErrPermission 校验
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm) // 递归创建目录
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}