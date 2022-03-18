package service

import (
	"testing"
)

func mock_get_page(url string) string {
	// mock your 'get_page()' function here
	if len(url) == 0 {
		return ""
	}
	return "this is the result for geting url "
}

func TestGetUrlMd5(t *testing.T) {
	d := NewDownloader(mock_get_page)
	testCase := []struct {
		url     string
		wantMd5 string
	}{
		{
			url:     "google.com",
			wantMd5: "058c16e9a1dfed3889b95d189bbee6f5",
		},
		{
			url:     "facebook.com",
			wantMd5: "058c16e9a1dfed3889b95d189bbee6f5",
		},
		{
			url:     "",
			wantMd5: "d41d8cd98f00b204e9800998ecf8427e",
		},
	}

	for _, tc := range testCase {
		got := d.GetUrlMd5(tc.url)
		if got != tc.wantMd5 {
			t.Errorf("got %v, wanted %v ,%v", got, tc.wantMd5, tc.url)
		}
	}
}
