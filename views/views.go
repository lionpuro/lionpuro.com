package views

import "time"

const (
	github_url   = "https://github.com/lionpuro"
	linkedin_url = "https://www.linkedin.com/in/lion-puro-09bb4b1a1/"
	email        = "mailto:connect@lionpuro.com"
)

func dateString(t time.Time) string {
	return t.Format("January _2, 2006")
}
