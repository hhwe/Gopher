// 这个示例程序展示如何写基础单元测试
package test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认 http 包的 Get 函数可以下载内容
func TestDownload(t *testing.T) {
	var urls = []struct {
		url string
		statusCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK,
		},
		{
			"http://www.bing.com",
			http.StatusNotFound,
		},
	}
	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url",
					checkMark)
				defer resp.Body.Close()
				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
