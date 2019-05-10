package configuration

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ory/x/viperx"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/ory/fosite"
	"github.com/ory/x/corsx"
)

var _ Provider = new(ViperProvider)

type ViperProvider struct {
	l logrus.FieldLogger
}

func NewViperProvider(l logrus.FieldLogger) *ViperProvider {
	return &ViperProvider{l: l}
}

func (v *ViperProvider) CORSEnabled(iface string) bool {
	return corsx.IsEnabled(v.l, "serve."+iface)
}

func (v *ViperProvider) CORSOptions(iface string) cors.Options {
	return corsx.ParseOptions(v.l, "serve."+iface)
}

func (v *ViperProvider) ProxyReadTimeout() time.Duration {
	return viperx.GetDuration(v.l, "serve.proxy.timeout.read", time.Second*5, "PROXY_SERVER_READ_TIMEOUT")
}

func (v *ViperProvider) ProxyWriteTimeout() time.Duration {
	return viperx.GetDuration(v.l, "serve.proxy.timeout.write", time.Second*10, "PROXY_SERVER_WRITE_TIMEOUT")
}

func (v *ViperProvider) ProxyIdleTimeout() time.Duration {
	return viperx.GetDuration(v.l, "serve.proxy.timeout.idle", time.Second*120, "PROXY_SERVER_IDLE_TIMEOUT")
}

func (v *ViperProvider) ProxyServeAddress() string {
	return fmt.Sprintf(
		"%s:%d",
		viperx.GetString(v.l, "serve.proxy.host", ""),
		viperx.GetInt(v.l, "serve.proxy.port", 4455),
	)
}

func (v *ViperProvider) APIServeAddress() string {
	return fmt.Sprintf(
		"%s:%d",
		viperx.GetString(v.l, "serve.api.host", ""),
		viperx.GetInt(v.l, "serve.api.port", 4456),
	)
}

func (v *ViperProvider) getURL(value string, key string) *url.URL {
	u, err := url.ParseRequestURI(value)
	if err != nil {
		v.l.WithError(err).Errorf(`Configuration key "%s" is missing or malformed.`, key)
		return nil
	}

	return u
}

func (v *ViperProvider) toScopeStrategy(value string, key string) fosite.ScopeStrategy {
	switch strings.ToLower(value) {
	case "hierarchic":
		return fosite.HierarchicScopeStrategy
	case "exact":
		return fosite.ExactScopeStrategy
	case "wildcard":
		return fosite.WildcardScopeStrategy
	case "none":
		return nil
	default:
		v.l.Errorf(`Configuration key "%s" declares unknown scope strategy "%s", only "hierarchic", "exact", "wildcard", "none" are supported. Falling back to strategy "none".`, key, value)
		return nil
	}
}
