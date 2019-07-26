package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tosashimanto/iot-camera-web/model/api"
	"github.com/tosashimanto/iot-camera-web/service"
	"github.com/tosashimanto/iot-camera-web/util"
	"net/http"
	"strconv"
	"time"
)

func CheckToken(c echo.Context) error {
	tokenPost := new(api.TokenJSONPost)
	if err := c.Bind(tokenPost); err != nil {
		return err
	}
	fmt.Println("Token=", tokenPost.Token)

	// Create token(JWT)
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	//claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	fmt.Println("token=" + t)
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func GetImageList(c echo.Context) error {

	image_manage_array := service.GetImageList()
	image_manages := make([]api.Image_manage, len(image_manage_array))
	for i := 0; i < len(image_manage_array); i++ {
		image_manages[i] = api.Image_manage{
			ID:           image_manage_array[i].ID,
			ImageData:    util.Base64Encoding(image_manage_array[i].ImageData),
			ExifLat:      image_manage_array[i].ExifLat,
			ExifLon:      image_manage_array[i].ExifLon,
			ExifDateTime: image_manage_array[i].ExifDateTime.String(),
			CreatedBy:    image_manage_array[i].CreatedBy,
			CreatedAt:    image_manage_array[i].CreatedAt.String(),
			UpdatedBy:    image_manage_array[i].UpdatedBy,
			UpdatedAt:    image_manage_array[i].UpdatedAt.String(),
		}
	}

	getImageListResponse := &api.GetImageListResponse{
		ImageManegeArray: image_manages[:len(image_manage_array)],
	}

	jsonString, _ := json.Marshal(getImageListResponse)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, getImageListResponse)

}

func GetUplaodURL(c echo.Context) error {

	const arrayLen = 5
	var imageUplaodURLArray [arrayLen]api.ImageUploadURL

	uploadUrl := "http://10.0.2.2:8080/iot_camera/v1/upload"
	for i := 0; i < len(imageUplaodURLArray); i++ {
		imageUplaodURLArray[i] = api.ImageUploadURL{
			Index:     i + 1,
			UploadUrl: uploadUrl + "/" + strconv.Itoa(i+1),
		}
	}

	putUploadURLResponse := &api.PutUploadURLResponse{
		ImageUplaodURLArray: imageUplaodURLArray[:arrayLen],
	}

	jsonString, _ := json.Marshal(putUploadURLResponse)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, putUploadURLResponse)
}
