package qiniuyun

import (
	"context"
	"ginblog/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var Zone = config.Zone
var AccessKey = config.AccessKey
var SecretKey = config.SecretKey
var Bucket = config.Bucket
var ImgUrl = config.QiniuSever

// UpLoadFile 上传文件函数
func UpLoadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := setConfig()

	putExtra := storage.PutExtra{}
	ret := storage.PutRet{}

	formUploader := storage.NewFormUploader(&cfg)

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}

	url := ImgUrl + ret.Key
	return url, nil
}

func setConfig() storage.Config {
	cfg := storage.Config{
		Zone:          selectZone(Zone),
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	return cfg
}

func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
