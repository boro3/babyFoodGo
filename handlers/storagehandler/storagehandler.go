package storagehandler

import (
	"babyFood/pkg/functions/pictureformat"
	"babyFood/pkg/functions/storagepath"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadProfilePicture(c echo.Context) error {
	file, err := c.FormFile("document")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	allowedSize := 1 * 1024 * 1024

	if allowedSize <= int(file.Size) {
		return c.JSON(http.StatusBadRequest, "Bad Request!: File too large!")
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	buff := make([]byte, 512)

	_, err = src.Read(buff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	_, err = src.Seek(0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	defer src.Close()

	filetype := http.DetectContentType(buff)
	fmt.Println(filetype)

	dst, err := storagepath.Create(`./uploads/profile/` + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	return c.JSON(http.StatusOK, file.Filename)
}

func UploadRecipePicture(c echo.Context) error {
	file, err := c.FormFile("document")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	allowedSize := 1 * 1024 * 1024

	if allowedSize <= int(file.Size) {
		return c.JSON(http.StatusBadRequest, "Bad Request!: File too large!")
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	buff := make([]byte, 512)

	_, err = src.Read(buff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	_, err = src.Seek(0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	defer src.Close()

	filetype := http.DetectContentType(buff)
	check := pictureformat.CheckFileFormat(filetype)
	if !check {
		return c.JSON(http.StatusBadRequest, "Bad Request!: Unsupported file format!")
	}

	dst, err := storagepath.Create(`./uploads/recipe/` + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	return c.JSON(http.StatusOK, file.Filename)
}

func DownloadProfilePicture(c echo.Context) error {
	img := c.Param("img")
	return c.File(`./uploads/profile/` + img)
}

func DownloadRecipePicture(c echo.Context) error {
	img := c.Param("img")
	return c.File(`./uploads/recipe/` + img)
}
