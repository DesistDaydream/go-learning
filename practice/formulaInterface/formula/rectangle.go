package formula

// Rectangle 长方形的属性。属性是长和宽。
type Rectangle struct {
	Length float32
	Width  float32
}

// Area 矩形求面积的方法
func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

// Perimeter 矩形求周长的方法
func (r Rectangle) Perimeter() float32 {
	return r.Length*2 + r.Width*2
}
