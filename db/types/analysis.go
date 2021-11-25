package types

// PerformanceType 基金区间业绩表现-风控指标
type PerformanceType struct {
	Period    string  `json:"period"`
	Absolute  float64 `json:"absolute"`
	Annual    float64 `json:"annual"`
	Abnormal  float64 `json:"abnormal"`
	Beyond    float64 `json:"beyond"`
	Vol       float64 `json:"vol"`
	Drawback  float64 `json:"drawback"`
	Sharpe    float64 `json:"sharpe"`
	Calmar    float64 `json:"calmar"`
	Timing    float64 `json:"timing"`
	Selection float64 `json:"selection"`
}

func NewPerformance() map[int]PerformanceType {
	var perf = make(map[int]PerformanceType, 6)
	var periods = map[int]string{
		1: "近1月", 3: "近3月", 6: "近6月", 12: "近1年", 36: "近3年", 60: "近5年",
	}
	for key, value := range periods {
		perf[key] = PerformanceType{
			Period:    value,
		}
	}
	return perf
}
