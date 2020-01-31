package ch3

type Config struct {
	Name     string `json:"SiteName"`
	URL      string `json:"SiteUrl"`
	Database struct {
		Name     string
		Host     string
		Port     int
		Username string
		Password string
	}
}
