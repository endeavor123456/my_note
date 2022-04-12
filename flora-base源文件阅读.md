# flora-base源文件阅读(以下为几段比较经典的代码段)

## from handler.go

```go
func shouldRefreshFunc() func(ok bool) bool {//返回值为func(ok bool)类型的数据
	mu := sync.RWMutex{}  //var mu sync.RWMutex  mu = sync.RWMutex{}     RWMutex为sync包里的结构体
	attempted := 0
	maxRetry := 6

	return func(pageOk bool) bool {
		mu.RLock()
		if (pageOk && attempted == 0) || (!pageOk && attempted > maxRetry) {//这一句也挺有意思
			mu.RUnlock()
			return false
		}
		mu.RUnlock()

		mu.Lock()
		if pageOk {
			attempted = 0
			mu.Unlock()
			return false
		}
		attempted++
		mu.Unlock()
		return true
	}
}
```
```go
func MakeErrorHandler(code int, data ErrorPageData) func(ginCtx *gin.Context) {
	renderer := render.HTML{
		Template: errorPageTemplate,
		Data:     data,
	}
	return func(ginCtx *gin.Context) {
		ginCtx.Render(code, renderer)
	}
}//plainpage包里的MakeErrorHandler函数

errPage := plainpage.MakeErrorHandler(503, plainpage.ErrorPageData{//plainpage包里的MakeErrorHandler函数进行调用
				Title:           "Unavailable 服务器无响应",
				StatusCodeTitle: "503",
				SubTitleEN:      "The frontend server is unavailable",
				SubTitleCN:      "前端服务器无响应",
				MessageEN:       "<p>Sorry the frontend server is unavailalbe, please try again later.</p>",
				MessageCN:       "<p>抱歉，前端服务器暂时无响应，请稍后刷新重试。</p>",
				Action:          autoRefresh,
			})

/*
上面变量errPage的定义是
var errPage func(ginCtx *gin.Context) errPage = plainpage.MakeErrorHandler(503, plainpage.ErrorPageData{//plainpage包里的MakeErrorHandler函数进行调用
				Title:           "Unavailable 服务器无响应",
				StatusCodeTitle: "503",
				SubTitleEN:      "The frontend server is unavailable",
				SubTitleCN:      "前端服务器无响应",
				MessageEN:       "<p>Sorry the frontend server is unavailalbe, please try again later.</p>",
				MessageCN:       "<p>抱歉，前端服务器暂时无响应，请稍后刷新重试。</p>",
				Action:          autoRefresh,
			})
*/
```
