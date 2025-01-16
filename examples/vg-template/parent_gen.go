// Code generated by vugu via vugugen DO NOT EDIT.
// Please regenerate instead of editing or add additional code in a separate file.

package main

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"
import "log"

func (c *Parent) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

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
				vgn.SetInnerHTML(vugu.HTML("Simple Form Example"))
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "parent-component"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " \"c\" is the Parent control that containts the form "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Note the inversion of control - we pass the Parent control, 'c' to the the Simpleform as its 'Parent' attribute "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " We need the Simpleform to be able to call into the Parent to set state in the parent in response to DOM events that the form handles "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					{
						vgcompKey := vugu.MakeCompKey(0xB55457FE74F75491^vgin.CurrentPositionHash(), vgiterkey)
						// ask BuildEnv for prior instance of this specific component
						vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Sampleform)
						if vgcomp == nil {
							// create new one if needed
							vgcomp = new(Sampleform)
							vgin.BuildEnv.WireComponent(vgcomp)
						}
						vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
						vgcomp.Parent = c
						vgout.Components = append(vgout.Components, vgcomp)
						vgn = &vugu.VGNode{Component: vgcomp}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "br", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML(""))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "hr", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML(""))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Note: there is a way to remove the repeated 'vg-if' statements. See the 'vg-template' example "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					if c.Submitted() {
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    The submit button has been pressed.\n                "}
							vgparent.AppendChild(vgn)
						}
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					if c.Submitted() {
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    What follows is the JSON encoded form contents that would be sent to the web server for further processing.\n                "}
							vgparent.AppendChild(vgn)
						}
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " We call c.EncodedFormData() to get the JSON encoded version of the form data "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					if c.Submitted() {
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "code", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(c.EncodedFormData())
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{{Namespace: "", Key: "id", Val: "paragragraphs"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						for i := 0; i < 5; i++ {
							var vgiterkey interface{} = i
							_ = vgiterkey
							i := i
							_ = i
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3)}	// <vg-template>
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
								vgparent.AppendChild(vgn)
								vgn.SetInnerHTML(fmt.Sprintf("%d", i))
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
								vgparent.AppendChild(vgn)
							}
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
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
