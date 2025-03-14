package locales

import (
	"context"
	"embed"
	"github.com/invopop/ctxi18n/i18n"
)

//go:embed translations.yaml
var Content embed.FS

func LocaleCode(ctx context.Context) string {
	return i18n.GetLocale(ctx).Code().String()
}
