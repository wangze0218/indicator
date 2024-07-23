package indicator

import "container/list"

type Kdj struct {
	n1     int
	n2     int
	n3     int
	kSma   *Sma       // 用于计算 K 值的 SMA
	dSma   *Sma       // 用于计算 D 值的 SMA
	dequeH *list.List // 用于存储最近 n1 个 Kline 的 High 值
	dequeL *list.List // 用于存储最近 n1 个 Kline 的 Low 值
}

// NewKdj 创建一个新的 Kdj 对象
func NewKdj(n1 int, n2 int, n3 int) *Kdj {
	return &Kdj{
		n1:     n1,
		n2:     n2,
		n3:     n3,
		kSma:   NewSma(),
		dSma:   NewSma(),
		dequeH: list.New(),
		dequeL: list.New(),
	}
}

// maxHigh 计算最近 n1 个 Kline 的最高值
func (k *Kdj) maxHigh() float64 {
	for k.dequeH.Len() > 1 && k.dequeH.Front().Value.(float64) < k.dequeH.Back().Value.(float64) {
		k.dequeH.Remove(k.dequeH.Front())
	}
	return k.dequeH.Front().Value.(float64)
}

// minLow 计算最近 n1 个 Kline 的最低值
func (k *Kdj) minLow() float64 {
	for k.dequeL.Len() > 1 && k.dequeL.Front().Value.(float64) > k.dequeL.Back().Value.(float64) {
		k.dequeL.Remove(k.dequeL.Front())
	}
	return k.dequeL.Front().Value.(float64)
}

func (k *Kdj) Update(bid Kline) (float64, float64, float64) {
	// 更新 dequeH 和 dequeL
	if k.dequeH.Len() >= k.n1 {
		k.dequeH.Remove(k.dequeH.Front())
	}
	if k.dequeL.Len() >= k.n1 {
		k.dequeL.Remove(k.dequeL.Front())
	}
	k.dequeH.PushBack(bid.High)
	k.dequeL.PushBack(bid.Low)

	// 计算 RSV
	if k.dequeH.Len() < k.n1 || k.dequeL.Len() < k.n1 {
		// 如果数据点不足 n1 个，无法计算 KDJ，返回默认值
		return 0, 0, 0
	}
	h := k.maxHigh()
	l := k.minLow()
	var rsv float64
	if h != l {
		rsv = (bid.Close - l) * 100.0 / (h - l)
	} else {
		rsv = 50.0 // 如果高低相等，RSV 为 50
	}

	// 更新 K 和 D
	K := k.kSma.Update(rsv)
	D := k.dSma.Update(K)
	J := 3.0*K - 2.0*D

	return K, D, J
}
