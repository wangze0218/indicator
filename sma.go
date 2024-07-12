package indicator

type Sma struct {
	maxLength int       // 控制切片的最大长度
	values    []float64 // 存储价格数据的切片
	sum       float64   // 用于累加计算移动平均值的总和
}

func NewSma(maxLength int) *Sma {
	return &Sma{
		maxLength: maxLength,
		values:    make([]float64, 0, maxLength),
	}
}

func (this *Sma) Update(price float64) float64 {
	if len(this.values) >= this.maxLength {
		// 如果超过最大长度，则移除最旧的数据
		this.sum -= this.values[0]
		this.values = this.values[1:]
	} else {
		// 如果未达到最大长度，则添加新数据
		this.values = append(this.values, price)
	}

	// 添加新数据并更新总和
	this.values = append(this.values, price)
	this.sum += price

	// 返回移动平均值
	return this.sum / float64(len(this.values))
}

func (this *Sma) GetAverage() float64 {
	// 计算并返回移动平均值
	return this.sum / float64(len(this.values))
}

func (this *Sma) Clone() *Sma {
	// 创建并返回当前 SMA 对象的克隆副本
	clone := &Sma{
		maxLength: this.maxLength,
		values:    make([]float64, len(this.values)),
		sum:       this.sum,
	}
	copy(clone.values, this.values)
	return clone
}
