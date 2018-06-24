package main

type iCat interface {
	getName() string
	meow() string
	scratch() string
}
type iWildCat interface {
	getBread() string
	growl() string
	scratch() string
}

type yardcat struct {
	name string
}

func (cat yardcat) getName() string {
	return cat.name
}
func (cat *yardcat) setName(name string) {
	cat.name = name
}
func (cat yardcat) meow() string {
	return "Yard meow"
}
func (cat yardcat) scratch() string {
	return "Yard scratch"
}

type homeCat struct {
	name string
}

func (cat homeCat) getName() string {
	return cat.name
}
func (cat *homeCat) setName(name string) {
	cat.name = name
}
func (cat homeCat) meow() string {
	return "Home meow"
}
func (cat homeCat) scratch() string {
	return "Home scratch"
}

type tigr struct {
	bread string
}

func (iwc tigr) getBread() string {
	return iwc.bread
}
func (iwc tigr) growl() string {
	return "GRRRRR"
}
func (iwc tigr) scratch() string {
	return "die"
}

type iWildCatAdaprer struct {
	iwc iWildCat
}

func (cat iWildCatAdaprer) getName() string {
	return cat.iwc.getBread()
}
func (cat iWildCatAdaprer) meow() string {
	return cat.iwc.growl()
}

func (cat iWildCatAdaprer) scratch() string {
	return cat.iwc.scratch()
}


func main() {
	yc := yardcat{}
	yc.setName("Yard")

	hc := homeCat{}
	hc.setName("Hme")

	tc := tigr{"tigr"}

	tca := iWildCatAdaprer{tc}

	catInfo(yc)
	catInfo(hc)
	catInfo(tca)
}

func catInfo(cat iCat) {
	println(cat.getName())
	println(cat.meow())
	println(cat.scratch())
	println("")
}
