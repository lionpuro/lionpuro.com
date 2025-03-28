package locales

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/invopop/ctxi18n/i18n"
)

//go:embed fi en
var Content embed.FS

func LocaleCode(ctx context.Context) string {
	return i18n.GetLocale(ctx).Code().String()
}

var finnishMonths = map[string]string{
	"January":   "Tammikuu",
	"February":  "Helmikuu",
	"March":     "Maaliskuu",
	"April":     "Huhtikuu",
	"May":       "Toukokuu",
	"June":      "Kesäkuu",
	"July":      "Heinäkuu",
	"August":    "Elokuu",
	"September": "Syyskuu",
	"October":   "Lokakuu",
	"November":  "Marraskuu",
	"December":  "Joulukuu",
}

func DateString(ctx context.Context, t time.Time) string {
	l := i18n.GetLocale(ctx).Code().String()
	if l == "fi" {
		y, _, d := t.Date()
		m := finnishMonths[t.Month().String()]
		return fmt.Sprintf("%d. %sta %d", d, m, y)
	}
	return t.Format("January _2, 2006")
}
