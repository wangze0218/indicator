package indicator

type Ema struct {
	Weight float64
	result float64
	age    uint32
	prices []float64
}

func NewEma(weight int32) *Ema {
	return &Ema{Weight: float64(weight)}
}

func (this *Ema) Update(price float64) float64 {
	alpha := 2.0 / (this.Weight + 1.0)
	this.prices = append(this.prices, price)
	if this.age > uint32(this.Weight) {
		this.prices = this.prices[1:]
	}
	// 计算
	for i, v := range this.prices {
		if i == 0 {
			// 初次更新
			this.result = v
		} else {
			// 后续更新
			this.result = alpha*v + (1-alpha)*this.result
		}
	}

	this.age += 1
	return this.result
}

func (this *Ema) GetPrice() float64 {
	return this.result
}

func (this *Ema) Clone() *Ema {
	return &Ema{Weight: this.Weight, result: this.result, age: this.age}
}
