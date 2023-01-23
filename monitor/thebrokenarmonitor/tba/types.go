package thebrokenarmonitor

var (
	firstRun     = true
	frontendLink = "https://www.the-broken-arm.com/1_en_0_sitemap.xml"
)

type Urlset struct {
	Image string `xml:"image,attr"`
	URL   []struct {
		Loc   string `xml:"loc"`
		Image []struct {
			Loc   string `xml:"loc"`
			Title string `xml:"title"`
		} `xml:"image"`
	} `xml:"url"`
}
