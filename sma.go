package indicator

type Sma struct {
	sum   float64 // 用于累加计算移动平均值的总和
	count int     // 当前数据点的数量
}

// NewSma 创建一个新的 Sma 对象
func NewSma() *Sma {
	return &Sma{}
}

// Update 更新 SMA 对象的价格数据并返回当前的移动平均值
func (s *Sma) Update(price float64) float64 {
	// 添加新数据并更新总和
	s.sum += price
	s.count++

	// 返回移动平均值
	return s.sum / float64(s.count)
}

// GetAverage 返回当前的移动平均值
func (s *Sma) GetAverage() float64 {
	// 计算并返回移动平均值
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

// Clone 创建并返回当前 SMA 对象的克隆副本
func (s *Sma) Clone() *Sma {
	// 创建并返回当前 SMA 对象的克隆副本
	return &Sma{
		sum:   s.sum,
		count: s.count,
	}
}
