package download

import (
	"testing"
)

func Test_New_Downloader(t *testing.T) {

	downloader := New()
	downloader.Html(url)
	downloader.Image(url)
}
