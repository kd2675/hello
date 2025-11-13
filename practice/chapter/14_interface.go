package chapter

func Interface() {
	p1 := newPerson("hoge", 10)
	p1.run()
	p1.stop()
	p2 := newPerson("fuga", 20)
	p2.run()
	p2.stop()
}

type action interface {
	run()
	stop()
}

type person struct {
	name string
	age  int
	action
}

func newPerson(name string, age int) *person {
	p := person{
		name: name,
		age:  age,
	}
	return &p
}

func (p *person) run() {
	println(p.name + " run")
}

func (p *person) stop() {
	println(p.name + " stop")
}
