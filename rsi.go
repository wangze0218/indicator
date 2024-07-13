package indicator

import "sync"

type Rsi struct {
	period     int
	gainSma    *Sma
	lossSma    *Sma
	prevPrice  float64
	m          sync.Mutex
	currentRsi float64
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
	this.currentRsi = 100 - (100 / (1 + rs))
	return this.currentRsi
}

func (this *Rsi) Clone() *Rsi {
	return &Rsi{
		period:     this.period,
		gainSma:    this.gainSma.Clone(),
		lossSma:    this.lossSma.Clone(),
		prevPrice:  this.prevPrice,
		currentRsi: this.currentRsi,
	}
}

func (this *Rsi) GetCurrentRsi() float64 {
	return this.currentRsi
}
