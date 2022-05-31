package queue

type MovingAverage struct {
	buf          []int
	p, size, sum int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{buf: make([]int, 0, size), size: size, p: -1}
}

func (this *MovingAverage) Next(val int) float64 {
	this.p = (this.p + 1) % this.size
	if len(this.buf) == this.size {
		this.sum -= this.buf[this.p]
	}
	this.buf[this.p] = val
	this.sum += val

	return float64(this.sum) / float64(len(this.buf))
}
