// Code generated by vugu via vugugen DO NOT EDIT.
// Please regenerate instead of editing or add additional code in a separate file.

package main

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"
import "log"

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "content"}}}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "button", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "make2"}}}
		vgparent.AppendChild(vgn)
		vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
			EventType:	"click",
			Func:		func(event vugu.DOMEvent) { c.Make(2, event) },
			// TODO: implement capture, etc. mostly need to decide syntax
		})
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Make 2"}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "button", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "make6"}}}
		vgparent.AppendChild(vgn)
		vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
			EventType:	"click",
			Func:		func(event vugu.DOMEvent) { c.Make(6, event) },
			// TODO: implement capture, etc. mostly need to decide syntax
		})
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Make 6"}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n    "}
		vgparent.AppendChild(vgn)
		for _, n := range c.Numbers {
			var vgiterkey interface{} = n
			_ = vgiterkey
			n := n
			_ = n
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			vgn.AddAttrInterface("id", n)
			vgn.SetInnerHTML(n)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n"}
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