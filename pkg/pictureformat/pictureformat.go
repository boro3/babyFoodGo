package pictureformat

var fileFormats = []string{
	"image/jpeg",
	"image/pjpeg",
	"image/png",
	"image/gif",
}

//Function for checking file format from given Array of permited formats.
//As input string of the file format is given. The fucintion returns true if the format is permited and false if not.
func CheckFileFormat(file string) bool {
	for _, format := range fileFormats {
		if format == file {
			return true
		}
	}
	return false
}
