package utils

type InfoSite struct {
	URL      string
	SiteName string
}

const (
	URL_200LAB       = "https://200lab.io/blog/"
	SITE_NAME_200LAB = "200lab"
)

const (
	SITE_NAME_VIBLO = "viblo"
	URL_VIBLO       = "https://viblo.asia"
	PAGE_SIZE_VIBLO = 20
)
const (
	NUM_OF_WORKER = 10
)

var SITES = []InfoSite{
	//{URL: URL_200LAB, SiteName: SITE_NAME_200LAB},
	{URL: URL_VIBLO + "/newest", SiteName: SITE_NAME_VIBLO},
}
