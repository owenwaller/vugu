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
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "main", Attr: []vugu.VGAttribute{{Namespace: "", Key: "role", Val: "main"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "container text-center"}}}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "class", Val: "mt-5"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "h1", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				vgn.SetInnerHTML(vugu.HTML("Embeded files and Translation Example"))
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("Select one of the radio boxes below to change the lanuage of the sentence. "))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " c.Msg() returns the message translated to the language set by the radio buttons "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(c.Msg())
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " c.SelectedLanguage() returns the BCP 47 langauge (in its short form) "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " \n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "The current language is: "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(c.SelectedLanguage())
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "radioboxes"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "h3", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("Select a language:"))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "input", Attr: []vugu.VGAttribute{{Namespace: "", Key: "type", Val: "radio"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "english"}, vugu.VGAttribute{Namespace: "", Key: "name", Val: "language"}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "English"}, vugu.VGAttribute{Namespace: "", Key: "checked", Val: "checked"}}}
					vgparent.AppendChild(vgn)
					vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
						EventType:	"change",
						Func:		func(event vugu.DOMEvent) { c.Change(event) },
						// TODO: implement capture, etc. mostly need to decide syntax
					})
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{{Namespace: "", Key: "for", Val: "english"}}}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("English"))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "br", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML(""))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "input", Attr: []vugu.VGAttribute{{Namespace: "", Key: "type", Val: "radio"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "french"}, vugu.VGAttribute{Namespace: "", Key: "name", Val: "language"}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "Français"}}}
					vgparent.AppendChild(vgn)
					vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
						EventType:	"change",
						Func:		func(event vugu.DOMEvent) { c.Change(event) },
						// TODO: implement capture, etc. mostly need to decide syntax
					})
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{{Namespace: "", Key: "for", Val: "french"}}}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("Français"))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "br", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML(""))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "input", Attr: []vugu.VGAttribute{{Namespace: "", Key: "type", Val: "radio"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "italian"}, vugu.VGAttribute{Namespace: "", Key: "name", Val: "language"}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "Italiano"}}}
					vgparent.AppendChild(vgn)
					vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
						EventType:	"change",
						Func:		func(event vugu.DOMEvent) { c.Change(event) },
						// TODO: implement capture, etc. mostly need to decide syntax
					})
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{{Namespace: "", Key: "for", Val: "italian"}}}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("Italiano"))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "br", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML(""))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
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
