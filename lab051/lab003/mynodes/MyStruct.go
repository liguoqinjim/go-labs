package mynodes

type Player struct {
	Pid int
	Px  int //位置
}

func (p *Player) Move() {
	p.Px += 2
}
