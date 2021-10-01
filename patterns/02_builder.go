package pattern

type burger struct{
    bun bool
    cheese bool
    salad bool
    chicken bool
    sauce bool
    beaf bool
}

type Burger_builder interface {
    add_bun()
    add_cheese()
    add_salad()
    add_chicken()
    add_sauce()
    add_beaf()
    make_burger()
    show_burger()
}

type first_burger struct {
    burger
}

func (b *first_burger) add_bun(){
    b.burger.bun = true
}

func (b *first_burger) add_cheese(){
    b.burger.cheese = true
}

func (b *first_burger) add_salad(){
    b.burger.salad = false
}

func (b *first_burger) add_chicken(){
    b.burger.chicken = true
}

func (b *first_burger) add_sauce(){
    b.burger.sauce = true
}

func (b *first_burger) add_beaf(){
    b.burger.beaf = false
}

func (b *first_burger) make_burger(){
    b.add_bun()
    b.add_cheese()
    b.add_salad()
    b.add_chicken()
    b.add_sauce()
    b.add_beaf()
}

func (b *first_burger) show_burger() {
    fmt.Printf("_ _ _-----_ _ _\nBun  ---  %t\nCheese  ---  %t\nSalad  ---  %t\nChicken  ---  %t\nSauce  ---  %t\nBeaf  ---  %t\nBun  ---  %t\n- - -_____- - -\n",b.bun, b.cheese, b.salad, b.chicken, b.sauce, b.beaf, b.bun)
}
