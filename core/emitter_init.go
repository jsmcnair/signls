package core

type InitEmitter struct {
	note      Note
	direction Direction
	pulse     uint64
	armed     bool
	triggered bool
}

func NewInitEmitter(direction Direction) *InitEmitter {
	return &InitEmitter{
		direction: direction,
		armed:     true,
		note: Note{
			Channel:  uint8(1),
			Key:      60,
			Velocity: 100,
		},
	}
}

func (e *InitEmitter) Activated() bool {
	return e.armed || e.triggered
}

func (e *InitEmitter) Note() Note {
	return e.note
}

func (e *InitEmitter) Arm() {
	e.armed = true
}

func (e *InitEmitter) Trig(g *Grid, x, y int) {
	if e.armed {
		g.Trig(x, y)
		e.triggered = true
		e.armed = false
		e.pulse = g.Pulse
	}
	// TODO: handle length and note off
}

func (e *InitEmitter) Emit(g *Grid, x, y int) {
	if !e.updated(g.Pulse) && e.triggered {
		g.Emit(x, y, e.direction)
		e.triggered = false
		e.pulse = g.Pulse
	}
}

func (e *InitEmitter) Direction() Direction {
	return e.direction
}

func (e *InitEmitter) Symbol() string {
	return "I"
}

func (e *InitEmitter) Color() string {
	return "165"
}

func (e *InitEmitter) Reset() {
	e.pulse = 0
	e.armed = true
	e.triggered = false
}

func (e *InitEmitter) updated(pulse uint64) bool {
	return e.pulse == pulse
}
