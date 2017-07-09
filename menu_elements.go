package obmenu

import "fmt"

//menuElement interface represents any element that could be wrapped in menu or pipe menu
type menuElement interface {
	PrintOut()
}

//PipeMenu is a representation of Openbox pipe menu element
//Its is a base type for any Openbox pipe menu.
//For additional info check out http://openbox.org/wiki/Help:Menus#Pipe_menu
type PipeMenu struct {
	inner []menuElement
}

//PrintOut prints pipemenu and all of its contains to stdout
func (p PipeMenu) PrintOut() {
	fmt.Println("<openbox_pipe_menu>")
	for _, v := range p.inner {
		v.PrintOut()
	}
	fmt.Println("</openbox_pipe_menu>")
}

//AddElement adds elements to the menu for later PrintOut
func (p *PipeMenu) AddElement(e menuElement) {
	p.inner = append(p.inner, e)
}

//Item is a representation of Openbox item element
type Item struct {
	Label   string
	Icon    string
	Actions []actionElement
}

//AddAction adds Action/ActionExecute element to run onclick
func (i *Item) AddAction(p actionElement) {
	i.Actions = append(i.Actions, p)
}

//SetIcon sets icon path
func (i *Item) SetIcon(path string) {
	i.Icon = path
}

//PrintOut prints Item and all of its contains to stdout
func (i Item) PrintOut() {
	fmt.Println("<Item label=\"", i.Label, "\">")
	for _, v := range i.Actions {
		v.PrintOut()
	}
	fmt.Println("</Item>")
	return
}

type actionElement interface {
	menuElement
	isAction()
}

//Action is a representation of Openbox Action element
//For additional info check out http://openbox.org/wiki/Help:Actions#Global_Actions
//Note: Window Actions are not applicable to menus(unless proven otherwise)
type Action struct {
	Name string
}

//PrintOut prints Action and all of its contains to stdout
func (a Action) PrintOut() {
	fmt.Print("<Action name=\"")
	fmt.Print(a.Name)
	fmt.Println("\"></Action>")
}

func (a Action) isAction() {
}

//ActionExecute is an Action for executing shell command
type ActionExecute struct {
	Command string
}

//PrintOut prints Action and all of its contains to stdout
func (a ActionExecute) PrintOut() {
	fmt.Println("<Action name=\"Execute\">")
	fmt.Print("<command>")
	fmt.Print(a.Command)
	fmt.Println("</command>")
	fmt.Println("</Action>")
}

func (a ActionExecute) isAction() {
}

//Separator is a representation of Openbox separator element
type Separator struct {
	Label string
}

//PrintOut prints Separator to stdout
func (s Separator) PrintOut() {
	fmt.Print("<separator label=\"")
	fmt.Print(s.Label)
	fmt.Println("\" />")
}

//NewPipeMenu creates new pipe menu
func NewPipeMenu() *PipeMenu {
	return &PipeMenu{}
}

//NewItem creates new "item" menu element
func NewItem(label string) *Item {
	return &Item{Label: label, Actions: nil}
}

//NewAction creates new "action" menu element
func NewAction(name string) *Action {
	return &Action{Name: name}
}

//NewExecute creates new "action" menu element with name="Execute"
func NewExecute(command string) *ActionExecute {
	return &ActionExecute{Command: command}
}

//NewSeparator creates new "separator" menu element
func NewSeparator(label string) *Separator {
	return &Separator{Label: label}
}
