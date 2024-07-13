// rsi.go

package indicator

import (
	"math"
	"sync"
)

// Rsi 表示相对强弱指数 (RSI)
type Rsi struct {
	period     int
	gainSma    *Sma
	lossSma    *Sma
	m          sync.Mutex
	currentRsi float64
	prices     []float64 // 存储价格数据
}

// NewRsi 创建一个新的 Rsi 实例
func NewRsi(period int) *Rsi {
	return &Rsi{
		period:  period,
		gainSma: NewSma(1500),
		lossSma: NewSma(1500),
	}
}

func (r *Rsi) setPrice(price float64) {
	// 添加价格到 prices 切片
	r.prices = append(r.prices, price)

	// 如果 prices 的长度大于两倍 period，则删除最旧的价格
	if len(r.prices) > 2*r.period {
		r.prices = r.prices[1:]
	}
}

// Update 更新 RSI 并返回当前值
func (r *Rsi) Update(price float64) float64 {
	defer r.m.Unlock()
	r.m.Lock()
	r.setPrice(price)
	if len(r.prices) < 2 {
		return 0 // 如果只有一个价格数据，无法计算 RSI，返回默认值
	}
	change := price - r.prices[len(r.prices)-1]

	var gain, loss float64
	if change > 0 {
		gain = change
		loss = 0
	} else {
		gain = 0
		loss = -change
	}

	avgGain := r.gainSma.Update(gain)
	avgLoss := r.lossSma.Update(loss)

	if avgLoss == 0 {
		return 100
	}

	rs := avgGain / avgLoss
	r.currentRsi = 100 - (100 / (1 + rs))
	return r.currentRsi
}

// GetCurrentRsi 返回当前的 RSI 值
func (r *Rsi) GetCurrentRsi() float64 {
	return r.currentRsi
}

// Clone 返回当前 Rsi 的克隆实例
func (r *Rsi) Clone() *Rsi {
	return &Rsi{
		period:     r.period,
		gainSma:    r.gainSma.Clone(),
		lossSma:    r.lossSma.Clone(),
		currentRsi: r.currentRsi,
		prices:     append([]float64{}, r.prices...), // 复制切片内容，避免共享底层数组
	}
}

// IdentifyBullishDivergence 识别牛市背离，并返回背离程度
func (r *Rsi) IdentifyBullishDivergence(lookbackLeft, lookbackRight, rangeLower, rangeUpper int) (bool, float64) {
	if len(r.prices) < lookbackRight+1 {
		return false, 0
	}

	currentRsi := r.GetCurrentRsi()
	previousRsi := r.GetRsiForIndex(len(r.prices) - 1 - lookbackRight)

	// 计算 RSI 差异的百分比
	rsiChange := currentRsi - previousRsi
	divergenceDegree := math.Abs(rsiChange / currentRsi * 100)

	// 检查牛市背离条件
	rsiHL := currentRsi > r.GetRsiForIndex(len(r.prices)-1-lookbackLeft) && inRange(r, len(r.prices)-1-lookbackLeft, rangeLower, rangeUpper)
	priceLL := r.prices[len(r.prices)-1] < r.prices[len(r.prices)-1-lookbackLeft]
	if priceLL && rsiHL {
		return true, divergenceDegree
	}

	return false, 0
}

// IdentifyBearishDivergence 识别熊市背离，并返回背离程度
func (r *Rsi) IdentifyBearishDivergence(lookbackLeft, lookbackRight, rangeLower, rangeUpper int) (bool, float64) {
	if len(r.prices) < lookbackRight+1 {
		return false, 0
	}

	currentRsi := r.GetCurrentRsi()
	previousRsi := r.GetRsiForIndex(len(r.prices) - 1 - lookbackRight)

	// 计算 RSI 差异的百分比
	rsiChange := currentRsi - previousRsi
	divergenceDegree := math.Abs(rsiChange / currentRsi * 100)

	// 检查熊市背离条件
	rsiLH := currentRsi < r.GetRsiForIndex(len(r.prices)-1-lookbackLeft) && inRange(r, len(r.prices)-1-lookbackLeft, rangeLower, rangeUpper)
	priceHH := r.prices[len(r.prices)-1] > r.prices[len(r.prices)-1-lookbackLeft]
	if priceHH && rsiLH {
		return true, divergenceDegree
	}

	return false, 0
}

// GetRsiForIndex 获取指定索引位置的 RSI 值
func (r *Rsi) GetRsiForIndex(index int) float64 {
	if index < 0 || index >= len(r.prices) {
		return 0.0 // 或者返回一个合适的默认值
	}
	return r.prices[index]
}

// inRange 检查当前 RSI 是否在指定范围内
func inRange(r *Rsi, index, lower, upper int) bool {
	rsi := r.GetRsiForIndex(index)
	return rsi >= float64(lower) && rsi <= float64(upper)
}
