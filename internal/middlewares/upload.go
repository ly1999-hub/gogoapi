package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/util/ptime"
	"os"
	"path"
	"path/filepath"
)

const (
	PathUpload = "/uploads/"
)

func UploadSingleFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		fileNameOrigin := file.Filename
		src, err := file.Open()
		if err != nil {
			fmt.Println("err file.Open", err.Error())
			return err
		}
		dir, _ := os.Getwd()
		defer src.Close()
		// note...
		fileName := ptime.Format(ptime.Now(), constant.FormatTime) + filepath.Ext(file.Filename)
		uploadFilePath := path.Join(dir + PathUpload + fileName)

		// Destination
		dst, err := os.Create(uploadFilePath)
		if err != nil {
			fmt.Println("error : ", err)
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println("err  io.Copy", err.Error())
			return err
		}

		c.Set("file", model.FileUploadInfo{
			Filename:       fileName,
			FileNameOrigin: fileNameOrigin,
			Path:           uploadFilePath,
			Ext:            filepath.Ext(file.Filename),
		})
		return next(c)
	}
}
