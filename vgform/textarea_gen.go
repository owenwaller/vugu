// Code generated by vugu via vugugen DO NOT EDIT.
// Please regenerate instead of editing or add additional code in a separate file.

package vgform

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"
import "log"

func (c *Textarea) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "textarea", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	vgn.AddAttrList(c.AttrMap)
	{
		b, err := vjson.Marshal(c.Value.StringValue())
		if err != nil {
			panic(err)
		}
		vgn.Prop = append(vgn.Prop, vugu.VGProperty{Key: "value", JSONVal: vjson.RawMessage(b)})
	}
	vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
		EventType:	"change",
		Func:		func(event vugu.DOMEvent) { c.handleChange(event) },
		// TODO: implement capture, etc. mostly need to decide syntax
	})
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
var _ log.Logger
