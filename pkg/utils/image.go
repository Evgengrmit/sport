package imageUtils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetStorageRootPath() string {
	storageRootPath := os.Getenv("STORAGE_PATH")
	return storageRootPath
}
func GetStorageRootUrl() string {
	return "https://crossfit-api.nihao.team/assets/"
}

/**
Скачиваем аватарку с сайта и изменяем её размер.
*/
func GetAvatarThumbUrl(url string) (error, string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15")
	req.Header.Set("Host", "crossfit1905.com")
	response, e := client.Do(req)

	if e != nil {
		log.Fatal(e)
		return e, ""
	}
	defer response.Body.Close()

	// уникальный хеш из названия URL на фотографию
	hashedFileName := GetMD5Hash(url)

	originImagePath := path.Join(GetStorageRootPath(), hashedFileName+".jpg")
	originImageFile, err := os.Create(originImagePath)
	if err != nil {
		log.Fatal("error on create file for avatar picture", err)
	}
	defer originImageFile.Close()

	_, err = io.Copy(originImageFile, response.Body)

	if err != nil {
		log.Fatal("error on save file avatar picture", err)
	}
	// сбрасываем позицию чтения для файла
	_, err = originImageFile.Seek(0, io.SeekStart)

	err, thumbUrl := resizePicture(originImageFile, hashedFileName)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	// если часто обращаться к серверу сайта за фоткой - происходит блокировка по IP.
	// засыпаем на 3 сек для обхода блокировки после завершения работы функции
	time.Sleep(3 * time.Second)

	return nil, thumbUrl
}

func resizePicture(imageData io.Reader, hashedFileName string) (error, string) {
	imageJpeg, err := jpeg.Decode(imageData)
	if err != nil {
		log.Fatal("error on image decode ", err)
	}

	resizedImage := resize.Resize(300, 0, imageJpeg, resize.Lanczos3)
	thumbFileName := hashedFileName + "_thumb.jpg"
	thumbFilePath := GetStorageRootPath() + "/" + thumbFileName

	resizedImageFile, err := os.Create(thumbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer resizedImageFile.Close()

	jpeg.Encode(resizedImageFile, resizedImage, nil)

	return nil, thumbFileName
}
