package service

import (
	"io"
	"mime/multipart"
	"mywebmall/conf"
	"os"
	"strconv"
	"strings"
)

// 上传头像到本地的函数
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string, ftype string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	// 目录不存在
	if !DirExistOrNot(basePath) {
		// 创建
		CreateDir(basePath)
	} else {
		// 目录存在，先删除之前的，再存储现在的头像
		os.RemoveAll(basePath)
		CreateDir(basePath)
	}
	// 用户的头像目录已经存在
	// 这里路径直接设为jpg，后续可对文件的格式进行检查，支持多种图片格式等
	avatarPath := basePath + userName + "." + ftype // todo: 把file的后缀提取出来
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0777)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + "." + ftype, nil
}

// 上传商品到本地
func UploadProductToLocalStatic(file multipart.File, userId uint, productName string, photoName string, ftype string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.ProductPath + "boss" + bId + "/" + productName + "/"
	// 目录不存在
	if !DirExistOrNot(basePath) {
		// 创建
		CreateDir(basePath)
	}
	// 用户的目录已经存在
	productPath := basePath + photoName + "." + ftype
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0777)
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + "/" + photoName + "." + ftype, nil
}

// 检查文件类型是否是png,jpg,jpeg类型
func CheckPhotoType(fileheader *multipart.FileHeader) (bool, string) {
	fileType := fileheader.Header.Get("Content-Type")
	strs := strings.Split(fileType, "/")
	if strs[0] != "image" {
		return false, ""
	} else if strs[1] != "jpeg" && strs[1] != "png" {
		return false, ""
	} else {
		return true, strs[1]
	}
}

// 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	// 如果返回了路径错误则目录不存在
	if err != nil {
		return false
	}
	// 再判断是否是一个目录
	return s.IsDir()
}

// 创建文件夹
func CreateDir(dirName string) bool {
	// MkdirAll会创建多级目录包括dirName中的父目录，后面的参数是文件访问许可
	err := os.MkdirAll(dirName, 0777)
	return err == nil
}
