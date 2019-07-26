package handler

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/tosashimanto/iot-camera-web/controller/dto"
	"github.com/tosashimanto/iot-camera-web/service"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// File Upload (Put)テスト用
func UploadImage(c echo.Context) error {

	fmt.Println("UploadImage")
	index, _ := strconv.Atoi(c.Param("index"))
	fmt.Println("index=", index)

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	fmt.Println("bodyBytes len=", len(bodyBytes))

	_, err := os.Stat("./tmp")
	if err != nil {
		if err := os.Mkdir("./tmp", 0777); err != nil {
			panic(err)
		}
	}

	fileName := createFilePath(index)
	ioutil.WriteFile(fileName, bodyBytes, os.ModePerm)

	locDto := parseExif(fileName)
	image := resizeImage(fileName)
	buf := new(bytes.Buffer)
	err0 := jpeg.Encode(buf, image, nil)
	if err0 != nil {
		fmt.Println("jpeg.Encode Error=", err0)
	}
	locDto.ImageData = buf.Bytes()
	service.InsertImage(&locDto)
	service.GetImageList()

	mimeType := c.Request().Header.Get("Content-Type")
	fmt.Println("mimeType=", mimeType)

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", fileName))

}

func createFilePath(index int) string {
	const format = "20060102_150405" // 24h表現、0埋めあり
	now_date := time.Now().Format(format)
	var buff = bytes.NewBuffer(make([]byte, 0, 100))
	buff.WriteString("./tmp/")
	buff.WriteString(fmt.Sprintf("ID%08d_", index) + now_date)
	buff.WriteString(".jpg")
	return buff.String()
}

func parseExif(fileName string) dto.LocationDto {

	dto := dto.LocationDto{}
	fmt.Println("fileName=", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("exif.Decode")
	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("x.DateTime err=", err)
		// panic(err)
	} else {
		time, _ := x.DateTime()
		println("time=", time.String())

		lat, lng, _ := x.LatLong()
		println("lat/lng=", lat, lng)

		dto.ExifDateTime = time
		dto.ExifLat = lat
		dto.ExifLon = lng
	}

	// println("x.String=", x.String())
	return dto
}

func resizeImage(fileName string) image.Image {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(200, 0, img, resize.Lanczos3)
	//out, err := os.Create(fileName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer out.Close()

	//err = jpeg.Encode(out, m, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return m
}
