package main

func (a Square) Area() float32{
	return float32(a.Sidelen)*float32(a.Sidelen)
}

func (a Circle) Area() float32{
	return float32(a.Rad)*float32(a.Rad)*3.14
}
