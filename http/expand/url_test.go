package expand

import (
	"testing"
)

func TestGet(t *testing.T) {
	url, _ := Get("http://goo.gl/1IW3oP")
	if url != "http://blog.satoshun.info/" {
		t.Error("fail expand url", url)
	}
}

func TestGets(t *testing.T) {
	urls := []string{"http://goo.gl/1IW3oP"}
	urls = Gets(urls)
	if urls[0] != "http://blog.satoshun.info/" {
		t.Error("fail expand url", urls[0])
	}
}
