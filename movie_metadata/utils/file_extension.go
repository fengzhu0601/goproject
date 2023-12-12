package utils

var extensions = []string{
	".mp4",
	".MP4",
	".wmv",
	".WMV",
	".rmvb",
	".RMVB",
	".avi",
	".AVI",
	".flv",
	".FLV",
	".mov",
	".MOV",
	".mkv",
	".MKV",
	".mpeg",
	".MPEG",
	".mpg",
	".MPG",
	".3gp",
	".3GP",
	".m4v",
	".M4V",
	".rm",
	".RM",
	".ts",
	".TS",
}

func IsMovieFile(fileName string) bool {
	for _, ext := range extensions {
		if GetFileExtension(fileName) == ext {
			return true
		}
	}
	return false
}
