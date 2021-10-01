package pattern

type Talking_Timer struct {
    Current_state State
}

type State interface {
    print()
}

type Starting_state struct {}

type Running_state struct {}

type Finished_state struct {}

func (s Starting_state) print() {
    fmt.Println("just starting")
}

func (r Running_state) print() {
    fmt.Println("running")
}

func (f Finished_state) print() {
    fmt.Println("done")
}

func (t *Talking_Timer) set_state(state State) {
    t.Current_state = state
}

func (t *Talking_Timer) State_switching() {
    for {
        switch t.Current_state {
            case Starting_state{}:
                t.Current_state.print()
                t.set_state(Running_state{})
            case Running_state{}:
                t.Current_state.print()
                t.set_state(Finished_state{})
            case Finished_state{}:
                t.Current_state.print()
                return
        }
    }
}
