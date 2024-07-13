package indicator

type FastSlowEma struct {
	fastEma *Ema
	midEma  *Ema
	slowEma *Ema
}

func NewFastSlowEma(fast int32, mid int32, slow int32) *FastSlowEma {
	return &FastSlowEma{
		fastEma: NewEma(fast),
		midEma:  NewEma(mid),
		slowEma: NewEma(slow),
	}
}

func (this *FastSlowEma) Update(price float64) (fast, mid, slow float64) {
	this.fastEma.Update(price)
	this.midEma.Update(price)
	this.slowEma.Update(price)
	return this.fastEma.GetPrice(), this.midEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) GetPrice() (fast, mid, slow float64) {
	return this.fastEma.GetPrice(), this.midEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) Clone() *FastSlowEma {
	return &FastSlowEma{
		fastEma: this.fastEma.Clone(),
		midEma:  this.midEma.Clone(),
		slowEma: this.slowEma.Clone(),
	}
}
