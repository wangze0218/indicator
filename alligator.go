package indicator

type Alligator struct {
	lips  *Ema
	teeth *Ema
	jaw   *Ema
}

func NewAlligator(fast int32, slow int32) *Alligator {
	return &Alligator{
		lips:  NewEma(fast),
		teeth: NewEma(slow),
		jaw:   NewEma(slow),
	}
}

func (this *Alligator) Update(price float64) (lips, teeth, jaw float64) {
	this.lips.Update(price)
	this.teeth.Update(price)
	this.jaw.Update(price)
	return this.lips.GetPrice(), this.teeth.GetPrice(), this.jaw.GetPrice()
}

func (this *Alligator) GetPrice() (lips, teeth, jaw float64) {
	return this.lips.GetPrice(), this.teeth.GetPrice(), this.jaw.GetPrice()
}

func (this *Alligator) Clone() *Alligator {
	return &Alligator{
		lips:  this.lips.Clone(),
		teeth: this.teeth.Clone(),
		jaw:   this.jaw.Clone(),
	}
}
