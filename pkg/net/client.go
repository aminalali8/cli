package net

import (
	"net/http"

	"github.com/briandowns/spinner"
)

type SpinnerTransport struct {
	Disabled bool

	Proxied http.RoundTripper
}

var DefaultSpinnerTransport = SpinnerTransport{
	Disabled: false,
	Proxied:  http.DefaultTransport,
}

func (st SpinnerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if !st.Disabled {
		spinner := makeSpinner()
		spinner.Suffix = " Fetching API data..."

		spinner.Start()

		defer spinner.Stop()
	}

	return st.Proxied.RoundTrip(req)
}

func GetCLIClient() *http.Client {
	return &http.Client{
		Transport: DefaultSpinnerTransport,
	}
}

func makeSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], defaultDuration)
}
