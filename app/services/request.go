package services

// Request 普通携带secucode的请求
type Request struct {
	Secucode string `form:"secucode"`
}

// SearchRequest 携带search关键字的请求
type SearchRequest struct {
	Search string `form:"search"`
}

// NetValueRequest 基金净值请求
type NetValueRequest struct {
	Secucode  string `form:"secucode"`
	Style     string `form:"style"`
	Benchmark string `form:"benchmark"`
}