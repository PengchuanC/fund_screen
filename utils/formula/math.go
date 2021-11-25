package formula

import "math"

// Array 数组
type Array interface {
	Var() float64
	Std() float64
	Mean() float64
	Sum() float64
	Cov(array Array) float64
	Pct() Array
}

// ArrayFloat64 float64 数组
type ArrayFloat64 []float64

// Sum 求和
func (a ArrayFloat64) Sum() float64 {
	var sum float64
	sum = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// Mean 求平均值
func (a ArrayFloat64) Mean() float64 {
	var (
		sum    float64
		length int
	)
	sum = a.Sum()
	length = len(a)
	return sum / float64(length)
}

// Var 计算方差
func (a ArrayFloat64) Var() float64 {
	var (
		sum    float64
		mean   float64
		length int
	)
	sum = 0
	mean = a.Mean()
	for _, v := range a {
		sum += math.Pow(v-mean, 2)
	}
	return sum / float64(length-1)
}

// Std 标准差
func (a ArrayFloat64) Std() float64 {
	var var_ float64
	var_ = a.Var()
	return math.Pow(var_, 1/2)
}

// Cov 计算协方差
func (a ArrayFloat64) Cov(b ArrayFloat64) float64 {
	var (
		meanA  float64
		meanB  float64
		sum    float64
		length int
	)
	length = len(a)
	meanA = a.Mean()
	meanB = b.Mean()
	for i := 0; i < length; i++ {
		sum += (a[i] - meanA) * (b[i] - meanB)
	}
	return sum / float64(length-1)
}

// Corr 计算相关系数
func (a ArrayFloat64) Corr(b ArrayFloat64) float64 {
	var (
		cov  float64
		stdA float64
		stdB float64
	)
	cov = a.Cov(b)
	stdA = a.Std()
	stdB = b.Std()
	return cov / (stdA * stdB)
}

// Pct 计算涨跌幅
func (a ArrayFloat64) Pct() ArrayFloat64 {
	var (
		pct ArrayFloat64
		length int
	)
	length = len(a)
	for i := 1; i < length; i++ {
		pct = append(pct, a[i] / a[i-1] - 1)
	}
	return pct
}
