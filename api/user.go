package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"

	"go_example_server/internal/csv"
	"go_example_server/log"
	"go_example_server/models"
)

// TODO: User用と分かる関数名に変更する
func UploadCSV(c echo.Context) error {
	fmt.Println()
	fh, err := c.FormFile("user_csv")
	if err != nil {
		log.AppLog.Error("CSV Upload error", zap.Error(err))
		return err
	}

	var users []models.User
	csv.ReadFile(fh, &users)
	fmt.Println(users)

	err = models.User{}.Save(&users)
	if err != nil {
		log.AppLog.Error("User Save Error", zap.Error(err))
	}

	return c.String(http.StatusOK, "Success!")
}

// TODO: User用と分かる関数名に変更する
func DownloadCSV(c echo.Context) error {
	users, err := models.User{}.FetchAll()
	if err != nil {
		log.AppLog.Error("User FetchAll Error", zap.Error(err))
	}
	fmt.Println(users)

	fp, fn, err := csv.WriteFile("user", users)
	defer os.Remove(fp)
	if err != nil {
		log.AppLog.Error("User CSV Write Error", zap.Error(err))
		return c.String(http.StatusInternalServerError, "CSV Output Error.")
	}

	return c.Attachment(fp, fn)
}
