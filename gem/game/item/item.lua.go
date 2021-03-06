// Code generated by glua; DO NOT EDIT.
package item

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Binditem generates a lua binding for item
func Binditem(L *lua.LState) {
	L.PreloadModule("gem.game.item", lBinditem)
}

// lBinditem generates the table for the item module
func lBinditem(L *lua.LState) int {
	mod := L.NewTable()

	lBindContainer(L, mod)

	lBindDefinition(L, mod)

	lBindStack(L, mod)

	L.Push(mod)
	return 1
}

func lBindContainer(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("item.Container")

	L.SetField(mt, "__call", L.NewFunction(lNewContainer))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ContainerMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Container", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("item.Container", mt)
}

func lNewContainer(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	retVal := NewContainer(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var ContainerMethods = map[string]lua.LGFunction{

	"add": lBindContainerAdd,

	"capacity": lBindContainerCapacity,

	"find_stack_of": lBindContainerFindStackOf,

	"remove_all_from_slot": lBindContainerRemoveAllFromSlot,

	"remove_from_slot": lBindContainerRemoveFromSlot,

	"set_slot": lBindContainerSetSlot,

	"slot": lBindContainerSlot,

	"swap_slots": lBindContainerSwapSlots,
}

func lBindContainerAdd(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*Stack)
	L.Remove(1)
	retVal := self.Add(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerCapacity(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	retVal := self.Capacity()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerFindStackOf(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*Definition)
	L.Remove(1)
	retVal := self.FindStackOf(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerRemoveAllFromSlot(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	retVal := self.RemoveAllFromSlot(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerRemoveFromSlot(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(int)
	L.Remove(1)
	retVal := self.RemoveFromSlot(arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerSetSlot(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(*Stack)
	L.Remove(1)
	self.SetSlot(arg0, arg1)
	return 0

}

func lBindContainerSlot(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	retVal := self.Slot(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindContainerSwapSlots(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Container)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(int)
	L.Remove(1)
	self.SwapSlots(arg0, arg1)
	return 0

}

func lBindDefinition(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("item.Definition")

	L.SetField(mt, "__call", L.NewFunction(lNewDefinition))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), DefinitionMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Definition", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("item.Definition", mt)
}

func lNewDefinition(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	retVal := NewDefinition(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var DefinitionMethods = map[string]lua.LGFunction{

	"description": lBindDefinitionDescription,

	"id": lBindDefinitionId,

	"name": lBindDefinitionName,

	"noted_id": lBindDefinitionNotedId,

	"shop_value": lBindDefinitionShopValue,

	"stackable": lBindDefinitionStackable,
}

func lBindDefinitionDescription(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.Description()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindDefinitionId(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.Id()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindDefinitionName(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.Name()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindDefinitionNotedId(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.NotedId()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindDefinitionShopValue(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.ShopValue()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindDefinitionStackable(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Definition)
	L.Remove(1)
	retVal := self.Stackable()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindStack(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("item.Stack")

	L.SetField(mt, "__call", L.NewFunction(lNewStack))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), StackMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Stack", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("item.Stack", mt)
}

func lNewStack(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*Definition)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(int)
	L.Remove(1)
	retVal := NewStack(arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var StackMethods = map[string]lua.LGFunction{

	"count": lBindStackCount,

	"definition": lBindStackDefinition,
}

func lBindStackCount(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Stack)
	L.Remove(1)
	retVal := self.Count()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindStackDefinition(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Stack)
	L.Remove(1)
	retVal := self.Definition()
	L.Push(glua.ToLua(L, retVal))
	return 1

}
