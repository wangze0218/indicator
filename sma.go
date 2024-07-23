package indicator

// Sma 结构体定义
type Sma struct {
	period int
	count  int
	sum    float64
}

// NewSma 创建一个新的 Sma 对象
func NewSma(period int) *Sma {
	return &Sma{
		period: period,
	}
}

// Update 更新 Sma 值
func (s *Sma) Update(price float64) float64 {
	if s.count >= s.period {
		// 先减去最早的值
		s.sum -= s.sum / float64(s.count)
	} else {
		s.count++
	}
	s.sum += price
	return s.sum / float64(s.count)
}

// Clone 创建并返回当前 Sma 对象的克隆副本
func (s *Sma) Clone() *Sma {
	return &Sma{
		period: s.period,
		count:  s.count,
		sum:    s.sum,
	}
}
