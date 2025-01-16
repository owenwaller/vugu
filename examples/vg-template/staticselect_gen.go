// Code generated by vugu via vugugen DO NOT EDIT.
// Please regenerate instead of editing or add additional code in a separate file.

package main

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"
import "log"

func (c *Staticselect) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "cards"}}}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{{Namespace: "", Key: "for", Val: "cars"}}}
		vgparent.AppendChild(vgn)
		vgn.SetInnerHTML(vugu.HTML("Choose a car:"))
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Note 'c' in this case is the Staticselect component itself. "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " We call Staticselect.Change when the value changes to read the new value "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "select", Attr: []vugu.VGAttribute{{Namespace: "", Key: "name", Val: "cars"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "cars"}}}
		vgparent.AppendChild(vgn)
		vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
			EventType:	"change",
			Func:		func(event vugu.DOMEvent) { c.Change(event) },
			// TODO: implement capture, etc. mostly need to decide syntax
		})
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "option", Attr: []vugu.VGAttribute{{Namespace: "", Key: "value", Val: "volvo"}}}
			vgparent.AppendChild(vgn)
			vgn.SetInnerHTML(vugu.HTML("Volvo"))
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "option", Attr: []vugu.VGAttribute{{Namespace: "", Key: "value", Val: "saab"}}}
			vgparent.AppendChild(vgn)
			vgn.SetInnerHTML(vugu.HTML("Saab"))
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "option", Attr: []vugu.VGAttribute{{Namespace: "", Key: "value", Val: "mercedes"}}}
			vgparent.AppendChild(vgn)
			vgn.SetInnerHTML(vugu.HTML("Mercedes"))
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "option", Attr: []vugu.VGAttribute{{Namespace: "", Key: "value", Val: "audi"}, vugu.VGAttribute{Namespace: "", Key: "selected", Val: "true"}}}
			vgparent.AppendChild(vgn)
			vgn.SetInnerHTML(vugu.HTML("Audi"))
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " the default in the code is set to Audi. See staticselect.go "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " \n"}
		vgparent.AppendChild(vgn)
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
var _ log.Logger