package download

type downloader struct {
	Proxy    bool
	Cache    bool
	Parallel bool
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

func (d *downloader) Htmls(url []string) {

	d.Parallel = true
}

func (d *downloader) Image(url string) {
	d.Cache = true
}

func (d *downloader) Images(url []string) {
	d.Cache = true
	d.Parallel = true
}
