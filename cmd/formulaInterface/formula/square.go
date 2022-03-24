package formula

// Square 正方形的属性，属性是边长。
type Square struct {
	Side float32
}

// Area 正方形求面积的方法，接收了正方形的结构体并使用结构体中的边长属性来计算面积
func (sq *Square) Area() float32 {
	return sq.Side * sq.Side
}

// Perimeter 正方形求周长的方法
func (sq *Square) Perimeter() float32 {
	return sq.Side * 4
}
