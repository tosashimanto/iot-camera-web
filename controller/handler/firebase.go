package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/tosashimanto/iot-camera-web/model/api"
	"log"
	"net/http"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func CheckFirebaseToken(c echo.Context) error {

	tokenPost := new(api.TokenJSONPost)
	if err := c.Bind(tokenPost); err != nil {
		return err
	}
	fmt.Println("CheckFirebaseToken Token=", tokenPost.Token)

	idToken := tokenPost.Token
	opt := option.WithCredentialsFile("./firebase/iotcamera-f4d88-firebase-adminsdk-fgwzw-517cc16744.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// return nil, fmt.Errorf("error initializing app: %v", err)
		log.Fatalf("error initializing app: %v\n", err)
	}

	// --------------------------------------
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	ctx := c.Request().Context()
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)

	return c.JSON(http.StatusOK, map[string]string{
		"token": "OK",
	})
}
