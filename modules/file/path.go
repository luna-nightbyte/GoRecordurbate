package file

import "strings"

var (
	Config_path          string = "./internal/settings/app.json"
	YoutubeDL_configPath string = "./internal/settings/youtube-dl.config"
	Index_path           string = "./internal/web/index.html"
	Login_path           string = "./internal/web/login.html"
	Videos_folder        string = "./output/videos"
	Log_path             string = "./output/GoRecordurbate.log"
)

// isVideoFile returns true if the file extension indicates a video file.
func IsVideoFile(filename string) bool {
	extensions := []string{".mp4", ".avi", ".mov", ".mkv"}
	lower := strings.ToLower(filename)
	for _, ext := range extensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}
