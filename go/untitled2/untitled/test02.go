package main

import "fmt"

type goods interface {
	goods()
}
type chinese struct {
}

func (person chinese) goods() {
	fmt.Println("我是中国人,我会说汉语")
}

type american struct {
}

func (person american) goods() {
	fmt.Println("i am american krese boll")
}

func greet(s goods) {
	s.goods()
}
func main() {
	var s int = 20
	fmt.Println(s)
	c := chinese{}
	m := american{}
	greet(c)
	greet(m)
}
