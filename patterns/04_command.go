package pattern

type icon struct {
    command Command
}

type Command interface {
    run()
}

type file interface {
    delete()
    open()
    show()
    close()
}

type delete_command struct {
    file File
}

type open_command struct {
    file File
}

type show_command struct {
    file File
}

type close_command struct {
    file File
}

func (c delete_command) run() {
    c.file.delete()
}

func (o open_command) run() {
    o.file.open()
}

func (s show_command) run() {
    s.file.show()
}

func (c close_command) run() {
    c.file.close()
}

func (i icon) click() {
    i.command.run()
}
