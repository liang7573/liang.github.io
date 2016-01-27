package download

type downloader struct {
}

func New() *downloader {
	d := new(downloader)
	d.init()
	return d
}

func (d *downloader) init() {

}

func (d *downloader) Html(url string) {

}

func (d *downloader) Image(url string) {

}
