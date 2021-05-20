package storagehandler

import (
	"babyFood/pkg/pictureformat"
	"babyFood/pkg/randstring"
	"babyFood/pkg/storagepath"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//Handler fucntion for uploading profile picture
//If successful returns string with the name of the uploaded picture
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
	check := pictureformat.CheckFileFormat(filetype)
	if !check {
		return c.JSON(http.StatusBadRequest, "Bad Request!: Unsupported file format!")
	}

	filename := randstring.RandString(8) + strings.Replace(file.Filename, " ", "_", -1)
	dst, err := storagepath.Create(`./uploads/profile/` + filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	return c.String(http.StatusOK, filename)
}

//Handler fucntion for uploading recipe picture
//If successful returns string with the name of the uploaded picture
func UploadRecipePicture(c echo.Context) error {
	file, err := c.FormFile("document")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	allowedSize := 1 * 1024 * 1024

	if allowedSize <= int(file.Size) {
		return c.JSON(http.StatusBadRequest, `Bad Request!: File too large!`)
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

	filename := randstring.RandString(8) + strings.Replace(file.Filename, " ", "_", -1)
	dst, err := storagepath.Create(`./uploads/recipe/` + filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}

	return c.String(http.StatusOK, filename)
}

//Function for downloading profile picture
//Returns file if the picture requested is found on storage
func DownloadProfilePicture(c echo.Context) error {
	img := c.Param("img")
	return c.File(`./uploads/profile/` + img)
}

//Function for downloading recipe picture
//Returns file if the picture requested is found on storage
func DownloadRecipePicture(c echo.Context) error {
	img := c.Param("img")
	return c.File(`./uploads/recipe/` + img)
}
