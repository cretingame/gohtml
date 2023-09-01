package bsx

// https://getbootstrap.com/docs/5.3/components/buttons/

type Button struct {
	Label    string
	class    string
	children []interface{}
}

func (b Button) HTML() string {
	if b.class == "" {
		b.class = "btn btn-primary"
	}

	out := `<button type="button" class="`
	out += b.class
	out += `">` + b.Label + `</button>`

	return out
}

type Div struct {
	class    string
	id       string
	children []interface{}
}

type StringWidget string

func (s StringWidget) HTML() string {
	return string(s)
}

/*
To be improved
*/
