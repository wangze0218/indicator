package indicator

import (
	"math"
)

type CurrBoll struct {
	n      int
	k      float64
	lines  []Kline
	smaSum float64
	mid    float64
	up     float64
	low    float64
}

// NewCurrBoll 初始化 CurrBoll 结构体
func NewCurrBoll(n int, k float64) *CurrBoll {
	return &CurrBoll{
		n:     n,
		k:     k,
		lines: make([]Kline, 0, n),
	}
}

// AddKline 添加一个新的 Kline 数据点，并更新布林带
func (b *CurrBoll) AddKline(line Kline) {
	if len(b.lines) >= b.n {
		// 移除最早的 kline 并更新 smaSum
		b.smaSum -= b.lines[0].Close
		b.lines = b.lines[1:]
	}
	// 添加新的 kline 并更新 smaSum
	b.lines = append(b.lines, line)
	b.smaSum += line.Close

	// 如果达到n个数据点，更新布林带
	if len(b.lines) == b.n {
		b.calculate()
	}
}

// calcSMA 计算简单移动平均值
func (b *CurrBoll) calcSMA() float64 {
	return b.smaSum / float64(b.n)
}

// calcSTD 计算标准差
func (b *CurrBoll) calcSTD(ma float64) float64 {
	var sum float64
	for _, line := range b.lines {
		sum += (line.Close - ma) * (line.Close - ma)
	}
	return math.Sqrt(sum / float64(b.n))
}

// calculate 计算布林带
func (b *CurrBoll) calculate() {
	b.mid = b.calcSMA()
	std := b.calcSTD(b.mid)
	b.up = b.mid + b.k*std
	b.low = b.mid - b.k*std
}

// GetBoll 返回当前布林带的值
func (b *CurrBoll) GetBoll() (mid, up, low float64) {
	// 如果数据不足，返回0
	if len(b.lines) < b.n {
		return 0, 0, 0
	}
	return b.mid, b.up, b.low
}
