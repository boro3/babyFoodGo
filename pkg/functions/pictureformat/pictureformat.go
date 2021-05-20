package pictureformat

var fileFormats = []string{
	"image/jpeg",
	"image/pjpeg",
	"image/png",
	"image/gif",
}

func CheckFileFormat(file string) bool {
	for _, format := range fileFormats {
		if format == file {
			return true
		}
	}
	return false
}
