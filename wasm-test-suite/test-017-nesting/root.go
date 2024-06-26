package main

type Root struct {
	// ShowWasm bool `vugu:"data"`
	// ShowGo bool   `vugu:"data"`
	// ShowVugu bool `vugu:"data"`
}

/*

- Fix broken component references (same comp gives same id each time https://github.com/vugu/vugu/issues/109).  And it should honor vg-key(?)
- Get @Event working on components, keep it simple, no propagation yet.  (Doc should encourage the use of interfaces)
- Get slots working and related functionality:
  <vg-slot name="DefaultSlot" - convention is to put "Slot" after the name
  <vg-slot-ref expr="SomeSlot()"/> - ? need some way to call the slot
  <vg-placeholder (not output to DOM, just a list of nodes)
- <vg-js-value="vugu.JSValueSetter(&c.SomeJSValue)" - to get a reference to an element after render
  oh - what if this is a callback and there is a render instruction telling it to call back into wasm
  with the js.Value - could even have two phases - one right after element is created, and one
  after it's children are created - vg-element-created= and vg-element-populated and each with different interfaces/
  func names so one thing could implement both
- // AfterRender() so components can perform manual stuff as needed - may not be needed if vg-js-value is a callback

*/
