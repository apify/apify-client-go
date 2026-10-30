package apify

import (
	"encoding/json"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	cases := []struct{ raw, want string }{
		// The scheme and host are lowercased and the default port dropped.
		{"HTTPS://Run-1.Example.COM:443", "https://run-1.example.com/"},
		{"http://example.com:80/a", "http://example.com/a"},
		{"https://example.com:8443/a", "https://example.com:8443/a"},
		// An empty path becomes "/", an existing path is kept.
		{"https://example.com", "https://example.com/"},
		{"https://example.com/v2/datasets/abc/items?format=json", "https://example.com/v2/datasets/abc/items?format=json"},
		{"https://example.com?q=1", "https://example.com/?q=1"},
		// A non-ASCII host becomes punycode.
		{"https://bücher.example/", "https://xn--bcher-kva.example/"},
		// Unsafe characters are percent-encoded; existing escapes are kept.
		{"https://example.com/a b?q=1 2#f g", "https://example.com/a%20b?q=1%202#f%20g"},
		{"https://example.com/a%20b/é", "https://example.com/a%20b/%C3%A9"},
		{"https://example.com/a?x='\"'", "https://example.com/a?x=%27%22%27"},
		// Dot segments are resolved.
		{"https://example.com/a/b/../c/./d", "https://example.com/a/c/d"},
		{"https://example.com/a/..", "https://example.com/"},
		// Userinfo and IPv6 hosts survive.
		{"https://user:pa ss@Example.com/", "https://user:pa%20ss@example.com/"},
		{"http://[::1]:80/x", "http://[::1]/x"},
		// The host of a URL that is not special is kept as given; its path is still resolved.
		{"mailto:Someone@Example.com", "mailto:Someone@Example.com"},
		{"mailto:a b", "mailto:a b"},
		{"foo://Example.COM/A", "foo://Example.COM/A"},
		{"foo://host/a/../b", "foo://host/b"},
		{"foo:/a/./b", "foo:/a/b"},
		// An empty password, and an empty userinfo, are dropped.
		{"https://u:@ex.com/", "https://u@ex.com/"},
		{"https://:@ex.com/", "https://ex.com/"},
		{"https://:p@ex.com/", "https://:p@ex.com/"},
		// A special URL has an authority, whatever the number of slashes after its scheme.
		{"https:foo", "https://foo/"},
		{"https:///foo/a", "https://foo/a"},
		{"https:\\\\foo\\a", "https://foo/a"},
		// The host is percent-decoded, and numeric hosts are written canonically.
		{"https://ex%41mple.com/", "https://example.com/"},
		{"http://0x7f.1/", "http://127.0.0.1/"},
		{"http://127.1/", "http://127.0.0.1/"},
		{"http://2130706433/", "http://127.0.0.1/"},
		{"http://[2001:db8:0:0:0:0:0:1]/", "http://[2001:db8::1]/"},
		{"http://[0:0:0:0:0:0:0:0]/", "http://[::]/"},
		{"http://[2001:db8:0:1:0:0:0:1]:8080/", "http://[2001:db8:0:1::1]:8080/"},
		// file URLs.
		{"file:///tmp/a", "file:///tmp/a"},
		{"file://localhost/tmp/a", "file:///tmp/a"},
		{"file:tmp/a", "file:///tmp/a"},
		// Tabs and line breaks are removed.
		{"https://exa\tmple.com/pa\nth", "https://example.com/path"},
		// More punycode.
		{"https://m\u00fcnchen.example/", "https://xn--mnchen-3ya.example/"},
		{"https://\u65e5\u672c\u8a9e.jp/", "https://xn--wgv71a119e.jp/"},
		{"https://\u30c9\u30e1\u30a4\u30f3\u540d\u4f8b.example/", "https://xn--eckwd4c7cu47r2wf.example/"},
		// Surrounding whitespace is dropped, a backslash is a slash.
		{"  https://example.com\\a  ", "https://example.com/a"},
	}
	for _, c := range cases {
		got, err := normalizeURL(c.raw)
		if err != nil {
			t.Errorf("normalizeURL(%q) failed: %v", c.raw, err)
			continue
		}
		if got != c.want {
			t.Errorf("normalizeURL(%q) = %q, want %q", c.raw, got, c.want)
		}
		again, err := normalizeURL(got)
		if err != nil || again != got {
			t.Errorf("normalizeURL is not idempotent for %q: got %q, %v", got, again, err)
		}
	}
}

func TestNormalizeURLRejectsInvalidURLs(t *testing.T) {
	for _, raw := range []string{"not a url", "/relative/path", "https://", "http://exa mple.com/", "http://example.com:99999/", "http://example.com:abc/", "http://[::1/", "http://1.2.3.4.5/", "http://256.1.1.1/", "http://foo.123/", "https://%ff.example/", "http://ex.com:+80/", "http://ex.com:-0/", "http://[fe80::1%25eth0]/", "file://host:80/x"} {
		if got, err := normalizeURL(raw); err == nil {
			t.Errorf("normalizeURL(%q) = %q, want an error", raw, got)
		}
	}
}

func TestEveryURLFieldOfTheModelsIsNormalized(t *testing.T) {
	const raw = "HTTP://Example.COM:80"
	const want = "http://example.com/"

	var run ActorRun
	mustUnmarshal(t, `{"id":"r","containerUrl":"`+raw+`","consoleUrl":"`+raw+`","output":{"results":"`+raw+`"}}`, &run)
	if run.ContainerURL != want || run.ConsoleURL != want || run.Output["results"] != want {
		t.Errorf("run URLs not normalized: %+v", run)
	}

	var task Task
	mustUnmarshal(t, `{"id":"t","standbyUrl":"`+raw+`"}`, &task)
	if task.StandbyURL == nil || *task.StandbyURL != want {
		t.Errorf("task standbyUrl not normalized: %v", task.StandbyURL)
	}
	var noStandby Task
	mustUnmarshal(t, `{"id":"t","standbyUrl":null}`, &noStandby)
	if noStandby.StandbyURL != nil {
		t.Errorf("a null standbyUrl must stay nil, got %q", *noStandby.StandbyURL)
	}

	var dataset Dataset
	mustUnmarshal(t, `{"id":"d","consoleUrl":"`+raw+`","itemsPublicUrl":"`+raw+`"}`, &dataset)
	if dataset.ConsoleURL != want || dataset.ItemsPublicURL != want {
		t.Errorf("dataset URLs not normalized: %+v", dataset)
	}

	var store KeyValueStore
	mustUnmarshal(t, `{"id":"s","consoleUrl":"`+raw+`","keysPublicUrl":"`+raw+`","recordsPublicUrl":"`+raw+`"}`, &store)
	if store.ConsoleURL != want || store.KeysPublicURL != want || store.RecordsPublicURL != want {
		t.Errorf("store URLs not normalized: %+v", store)
	}

	var key KeyValueStoreKey
	mustUnmarshal(t, `{"key":"k","size":1,"recordPublicUrl":"`+raw+`"}`, &key)
	if key.RecordPublicURL != want {
		t.Errorf("key recordPublicUrl not normalized: %q", key.RecordPublicURL)
	}

	var queue RequestQueue
	mustUnmarshal(t, `{"id":"q","consoleUrl":"`+raw+`"}`, &queue)
	if queue.ConsoleURL != want {
		t.Errorf("queue consoleUrl not normalized: %q", queue.ConsoleURL)
	}

	var webhook Webhook
	mustUnmarshal(t, `{"id":"w","requestUrl":"`+raw+`"}`, &webhook)
	if webhook.RequestURL != want {
		t.Errorf("webhook requestUrl not normalized: %q", webhook.RequestURL)
	}
	var noURL Webhook
	mustUnmarshal(t, `{"id":"w","requestUrl":null}`, &noURL)
	if noURL.RequestURL != "" {
		t.Errorf("a null requestUrl must stay empty, got %q", noURL.RequestURL)
	}

	var dispatch WebhookDispatch
	mustUnmarshal(t, `{"id":"x","webhookId":"w","webhook":{"_id":"w","actionType":"HTTP","isAdHoc":false,"requestUrl":"`+raw+`","slackMessage":{}}}`, &dispatch)
	if dispatch.Webhook == nil || dispatch.Webhook.RequestURL != want || dispatch.Webhook.ID != "w" {
		t.Errorf("dispatch webhook summary not parsed and normalized: %+v", dispatch.Webhook)
	}
	if _, ok := dispatch.Webhook.Extra["slackMessage"]; !ok {
		t.Errorf("an unmodelled field of the summary must land in Extra: %v", dispatch.Webhook.Extra)
	}

	var item ActorStoreListItem
	mustUnmarshal(t, `{"id":"a","url":"`+raw+`","userPictureUrl":"`+raw+`"}`, &item)
	if item.URL == nil || *item.URL != want || item.UserPictureURL == nil || *item.UserPictureURL != want {
		t.Errorf("store item URLs not normalized: %+v", item)
	}

	var user User
	mustUnmarshal(t, `{"id":"u","username":"me","profile":{"name":"N","pictureUrl":"`+raw+`","websiteUrl":null}}`, &user)
	if user.Profile == nil || user.Profile.PictureURL == nil || *user.Profile.PictureURL != want || user.Profile.WebsiteURL != nil {
		t.Errorf("user profile URLs not normalized: %+v", user.Profile)
	}
}

func TestAnInvalidURLFieldFailsTheResponse(t *testing.T) {
	var run ActorRun
	if err := json.Unmarshal([]byte(`{"id":"r","consoleUrl":"not a url"}`), &run); err == nil {
		t.Fatal("expected an error for an invalid URL field")
	}
}

func mustUnmarshal(t *testing.T, data string, v any) {
	t.Helper()
	if err := json.Unmarshal([]byte(data), v); err != nil {
		t.Fatalf("unmarshal %s: %v", data, err)
	}
}
