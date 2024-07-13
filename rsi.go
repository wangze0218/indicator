package indicator

import "sync"

// Rsi 表示相对强弱指数 (RSI)
type Rsi struct {
	period     int
	gainSma    *Sma
	lossSma    *Sma
	prevPrice  float64
	m          sync.Mutex
	currentRsi float64
	prices     []float64 // 存储价格数据
}

// NewRsi 创建一个新的 Rsi 实例
func NewRsi(period int) *Rsi {
	return &Rsi{
		period:  period,
		gainSma: NewSma(period),
		lossSma: NewSma(period),
	}
}

// Update 更新 RSI 并返回当前值
func (r *Rsi) Update(price float64) float64 {
	defer r.m.Unlock()
	r.m.Lock()
	if r.prevPrice == 0 {
		r.prevPrice = price
		return 0
	}

	change := price - r.prevPrice
	r.prevPrice = price

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
		prevPrice:  r.prevPrice,
		currentRsi: r.currentRsi,
		prices:     append([]float64{}, r.prices...), // 复制切片内容，避免共享底层数组
	}
}

// IdentifyDivergence 识别背离
func (r *Rsi) IdentifyDivergence(lookback int) bool {
	if len(r.prices) < lookback+1 {
		return false
	}

	currentRsi := r.GetCurrentRsi()
	currentPrice := r.prices[len(r.prices)-1]
	previousPrice := r.prices[len(r.prices)-1-lookback]

	// 检查牛市背离
	if currentPrice < previousPrice && currentRsi > r.GetRsiForIndex(len(r.prices)-1-lookback) {
		return true
	}

	// 检查熊市背离
	if currentPrice > previousPrice && currentRsi < r.GetRsiForIndex(len(r.prices)-1-lookback) {
		return true
	}

	return false
}

// GetRsiForIndex 获取指定索引位置的 RSI 值
func (r *Rsi) GetRsiForIndex(index int) float64 {
	if index < 0 || index >= len(r.prices) {
		return 0.0 // 或者返回一个合适的默认值
	}
	return r.prices[index]
}

// SetPrices 设置价格数据
func (r *Rsi) SetPrices(prices []float64) {
	r.prices = prices
}
