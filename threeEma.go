package indicator

type FastSlowEma struct {
	moreFast *Ema
	fastEma  *Ema
	midEma   *Ema
	slowEma  *Ema
}

func NewFastSlowEma(moreFast, fast, mid, slow int32) *FastSlowEma {
	return &FastSlowEma{
		moreFast: NewEma(moreFast),
		fastEma:  NewEma(fast),
		midEma:   NewEma(mid),
		slowEma:  NewEma(slow),
	}
}

func (this *FastSlowEma) Update(price float64) (moreFast, fast, mid, slow float64) {
	this.moreFast.Update(price)
	this.fastEma.Update(price)
	this.midEma.Update(price)
	this.slowEma.Update(price)
	return this.moreFast.GetPrice(), this.fastEma.GetPrice(), this.midEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) GetPrice() (moreFast, fast, mid, slow float64) {
	return this.moreFast.GetPrice(), this.fastEma.GetPrice(), this.midEma.GetPrice(), this.slowEma.GetPrice()
}

func (this *FastSlowEma) Clone() *FastSlowEma {
	return &FastSlowEma{
		moreFast: this.moreFast.Clone(),
		fastEma:  this.fastEma.Clone(),
		midEma:   this.midEma.Clone(),
		slowEma:  this.slowEma.Clone(),
	}
}
