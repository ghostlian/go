package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"

	//1.请求 - 模拟用户访问浏览器
	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/")

	//2.检查 - 是否无错误且 200
	//t 为 testing 标准库里的 testing.T 对象
	//第二个参数为错误对象 err
	//第三个参数为出错时显示的信息（选填）
	assert.NoError(t, err, "有错误发生，err 不为空")

	// 会断言两个值相等
	// 第一个参数同上，
	// 第二个参数是期待的状态码，
	// 第三个参数是请求返回的状态码，
	// 第四个参数为出错时显示的信息（选填）。
	assert.Equal(t, 200, resp.StatusCode, "返回状态码 200")
}

func TestAboutPage(t *testing.T) {
	baseURL := "http://localhost:3000"

	//1.请求 - 模拟用户访问浏览器
	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/about")

	//2.检查 - 是否无措且200
	assert.NoError(t, err, "有错误发送，err不能为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码200")

}

func TestAllPages(t *testing.T) {
	baseURL := "http://localhost:3000"

	//1.声明+初始化测试数据
	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/3", 200},
		{"GET", "/articles/3/edit", 200},
		{"POST", "/articles/3", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 404},
	}

	//2.遍历当前所有测试

	for _, test := range tests {
		t.Logf("当前请求 URL: %v \n", test.url) //辅助方法打印数据

		var (
			resp *http.Response
			err  error
		)

		//2.1请求以获取响应
		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}

		//2,2 断言
		assert.NoError(t, err, "请求"+test.url+"时候报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+"应返回状态码 "+strconv.Itoa(test.expected))
	}
}
