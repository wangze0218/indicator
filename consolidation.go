package indicator

type Consolidation struct {
	upper   *Ema // 盘整上边界
	marking *Ema // 对标线
	lower   *Ema // 盘整下边界
	through bool // 是否盘整
}

func NewConsolidation(boundary int32, marking int32) *Consolidation {
	return &Consolidation{
		upper:   NewEma(boundary),
		marking: NewEma(marking),
		lower:   NewEma(boundary),
	}
}

func (this *Consolidation) Update(bHPrice, bLPrice, mPrice float64) bool {
	this.upper.Update(bHPrice)
	this.marking.Update(mPrice)
	this.lower.Update(bLPrice)
	return true
}

func (this *Consolidation) GetThrough(point float64) bool {
	if this.marking.GetPrice() < this.upper.GetPrice()+point && this.marking.GetPrice() > this.lower.GetPrice()-point {
		return true
	}
	return false
}

func (this *Consolidation) Clone() *Consolidation {
	return &Consolidation{
		upper:   this.upper.Clone(),
		marking: this.marking.Clone(),
		lower:   this.lower.Clone(),
	}
}
