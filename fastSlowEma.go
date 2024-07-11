package indicator

type FastSlowEma struct {
	fastEma *Ema
	slowEma *Ema
}

func NewFastSlowEma(fast int32, slow int32) *FastSlowEma {
	return &FastSlowEma{
		fastEma: NewEma(fast),
		slowEma: NewEma(slow),
	}
}

func (this *FastSlowEma) Update(price float64) (fast, slow float64) {
	this.fastEma.Update(price)
	this.slowEma.Update(price)
	return this.fastEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) GetPrice() (fast, slow float64) {
	return this.fastEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) Clone() *FastSlowEma {
	return &FastSlowEma{
		fastEma: this.fastEma.Clone(),
		slowEma: this.slowEma.Clone(),
	}
}
