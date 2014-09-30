package expand

import (
	"testing"
)

func TestGet(t *testing.T) {
	url, _ := Get("http://goo.gl/hsGGW0")
	if url != "http://blog.s-shun.net/" {
		t.Error("fail expand url", url)
	}

	url, _ = Get("http://bit.ly/1uYGBgk")
	if url != "http://blog.s-shun.net/" {
		t.Error("fail expand url", url)
	}
}

func TestGets(t *testing.T) {
	urls := []string{"http://goo.gl/hsGGW0"}
	urls = Gets(urls)
	if urls[0] != "http://blog.s-shun.net/" {
		t.Error("fail expand url", urls[0])
	}
}
