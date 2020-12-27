package ghostclient

import "net/url"

// Options ...
type Option func(*C)
type ContentOption func(*url.URL)

func WithBaseURL(u *url.URL) Option {
	return func(client *C) {
		client.baseUrl = u
	}
}

func WithVerbose() Option {
	return func(client *C) {
		client.verbose = true
	}
}

func WithBasicAuth(username, password string) Option {
	return func(client *C) {
		client.basicAuthUsername = username
		client.basicAuthPassword = password
	}
}

// WithKey for authetication
func WithKey(key string) ContentOption {
	return func(u *url.URL) {
		q := u.Query()
		q.Set("key", key)
		u.RawQuery = q.Encode()
	}
}

// WithInclude request addiional parmeters from the api
func WithInclude(includes ...string) ContentOption {
	return func(u *url.URL) {
		var str string
		for _, include := range includes {
			if str == "" {
				str = include
				continue
			}
			str = str + "," + include
		}
		q := u.Query()
		q.Set("include", str)
		u.RawQuery = q.Encode()
	}
}
