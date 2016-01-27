package download

import (
	"testing"
)

func Test_one_html(t *testing.T) {

	url := "http://www.baidu.com"
	content := Html(url)
	if len(content) < 1 {
		t.Error("download html page failed")
	}
}

func Test_multi_html(t *testing.T) {
	var urls []string
	urls = append(urls, "http://www.speedtest.net/")
	urls = append(urls, "http://www.speedtest.net/")
	urls = append(urls, "http://www.speedtest.net/")
	urls = append(urls, "http://www.speedtest.net/")

	contents := Htmls(urls)

	for index, content := range contents {
		if len(content) < 1 {
			t.Error("download ", urls[index], " page failed")
		}
	}
}
