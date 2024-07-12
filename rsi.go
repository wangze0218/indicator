package indicator

import "sync"

type Rsi struct {
	period    int
	gainSma   *Sma
	lossSma   *Sma
	prevPrice float64
	m         sync.Mutex
}

func NewRsi(period int) *Rsi {
	return &Rsi{
		period:  period,
		gainSma: NewSma(period),
		lossSma: NewSma(period),
	}
}

func (this *Rsi) Update(price float64) float64 {
	defer this.m.Unlock()
	this.m.Lock()
	if this.prevPrice == 0 {
		this.prevPrice = price
		return 0
	}

	change := price - this.prevPrice
	this.prevPrice = price

	var gain, loss float64
	if change > 0 {
		gain = change
		loss = 0
	} else {
		gain = 0
		loss = -change
	}

	avgGain := this.gainSma.Update(gain)
	avgLoss := this.lossSma.Update(loss)

	if avgLoss == 0 {
		return 100
	}

	rs := avgGain / avgLoss
	rsi := 100 - (100 / (1 + rs))
	return rsi
}
