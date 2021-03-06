// Code generated by mediator autocoder; DO NOT EDIT.
package main

import (
	"github.com/celskeggs/mediator/platform/atoms"
	"github.com/celskeggs/mediator/platform/datum"
	"github.com/celskeggs/mediator/platform/format"
	"github.com/celskeggs/mediator/platform/framework"
	"github.com/celskeggs/mediator/platform/procs"
	"github.com/celskeggs/mediator/platform/types"
	"github.com/celskeggs/mediator/platform/world"
)

//mediator:declare MobPlayerData /mob/player /mob
type MobPlayerData struct {
}

func NewMobPlayerData(src *types.Datum, _ *MobPlayerData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("player.dmi"))
	src.SetVar("desc", types.String("A handsome and dashing rogue."))
	src.SetVar("name", types.String("player"))
	src.SetVar("verbs", src.Var("verbs").Invoke(nil, "+", atoms.NewVerb("look", "/mob/player", "look")))
}

func (chunk *MobPlayerData) Shadow0ForProcBump(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	varobstacle := types.Param(allargs, 0)
	(varsrc).Invoke(varusr, "<<", types.String("You bump into "+format.FormatMacro("the", varobstacle)+"."))
	(varsrc).Invoke(varusr, "<<", procs.NewSound("ouch.wav"))
	return out
}

func (chunk *MobPlayerData) ProcStat(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	_ = procs.Invoke(atoms.WorldOf(varsrc), varusr, "statpanel", types.String("Inventory"), varsrc.Var("contents"))
	return out
}

func (chunk *MobPlayerData) Proclook(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	(varsrc).Invoke(varusr, "<<", types.String("You see..."))
	for _, varo := range datum.Elements(procs.Invoke(atoms.WorldOf(varsrc), varusr, "oview")) {
		if !types.IsType(varo, "/atom/movable") {
			continue
		}
		(varsrc).Invoke(varusr, "<<", types.String(format.FormatMacro("The", varo)+".  "+format.FormatMacro("The", (varo).Var("desc"))))
	}
	return out
}

func (chunk *MobPlayerData) ProcBump(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	varobstacle := types.Param(allargs, 0)
	_ = chunk.Shadow0ForProcBump(varsrc, varusr, allargs)
	_ = (varobstacle).Invoke(varusr, "Bumped", varsrc)
	return out
}

//mediator:declare MobRatData /mob/rat /mob
type MobRatData struct {
}

func NewMobRatData(src *types.Datum, _ *MobRatData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("rat.dmi"))
	src.SetVar("desc", types.String("It's quite large."))
	src.SetVar("name", types.String("rat"))
}

func (chunk *MobRatData) ProcBumped(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	varbumper := types.Param(allargs, 0)
	(varbumper).Invoke(varusr, "<<", types.String("The giant rat defends its territory ferociously!"))
	varsrc.SetVar("dir", procs.Invoke(atoms.WorldOf(varsrc), varusr, "get_dir", varsrc, varbumper))
	_ = procs.Invoke(atoms.WorldOf(varsrc), varusr, "flick", types.String("fight"), varsrc)
	return out
}

//mediator:declare TurfFloorData /turf/floor /turf
type TurfFloorData struct {
}

func NewTurfFloorData(src *types.Datum, _ *TurfFloorData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("floor.dmi"))
	src.SetVar("name", types.String("floor"))
}

//mediator:declare TurfWallData /turf/wall /turf
type TurfWallData struct {
}

func NewTurfWallData(src *types.Datum, _ *TurfWallData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("wall.dmi"))
	src.SetVar("density", types.Int(1))
	src.SetVar("opacity", types.Int(1))
	src.SetVar("name", types.String("wall"))
}

//mediator:extend ExtObjData /obj
type ExtObjData struct {
}

func NewExtObjData(src *types.Datum, _ *ExtObjData, _ ...types.Value) {
	src.SetVar("verbs", src.Var("verbs").Invoke(nil, "+", atoms.NewVerb("get", "/obj", "get")))
	src.SetVar("verbs", src.Var("verbs").Invoke(nil, "+", atoms.NewVerb("drop", "/obj", "drop")))
}

func (chunk *ExtObjData) Procget(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	(varusr).Invoke(varusr, "<<", types.String("You get "+format.FormatMacro("the", varsrc)+"."))
	_ = (varsrc).Invoke(varusr, "Move", varusr)
	return out
}

func (*ExtObjData) SettingsForProcget() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeOView,
			In:   true,
		},
	}
}

func (chunk *ExtObjData) Procdrop(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	(varusr).Invoke(varusr, "<<", types.String("You drop "+format.FormatMacro("the", varsrc)+"."))
	_ = (varsrc).Invoke(varusr, "Move", (varusr).Var("loc"))
	return out
}

func (*ExtObjData) SettingsForProcdrop() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeUsr,
			In:   true,
		},
	}
}

//mediator:declare ObjCheeseData /obj/cheese /obj
type ObjCheeseData struct {
}

func NewObjCheeseData(src *types.Datum, _ *ObjCheeseData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("cheese.dmi"))
	src.SetVar("desc", types.String("It is quite smelly."))
	src.SetVar("name", types.String("cheese"))
	src.SetVar("verbs", src.Var("verbs").Invoke(nil, "+", atoms.NewVerb("eat", "/obj/cheese", "eat")))
}

func (chunk *ObjCheeseData) Proceat(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	(varusr).Invoke(varusr, "<<", types.String("You take a bite of the cheese. Bleck!"))
	varsrc.SetVar("suffix", types.String("(nibbled)"))
	return out
}

func (*ObjCheeseData) SettingsForProceat() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeUsr,
			In:   true,
		},
	}
}

func (chunk *ObjCheeseData) ProcMove(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	out = varsrc.SuperInvoke(varusr, "github.com/celskeggs/mediator-examples/yourfirstworld.ObjCheeseData", "Move", allargs...)
	for _, varrat := range datum.Elements(procs.Invoke(atoms.WorldOf(varsrc), varusr, "view", varsrc.Var("loc"))) {
		if !types.IsType(varrat, "/mob/rat") {
			continue
		}
		_ = procs.Invoke(atoms.WorldOf(varsrc), varusr, "walk_to", varrat, varsrc.Var("loc"), types.Int(1), types.Int(5))
	}
	return out
}

func (*ObjCheeseData) SettingsForProcMove() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeUsr,
			In:   true,
		},
	}
}

//mediator:declare ObjScrollData /obj/scroll /obj
type ObjScrollData struct {
}

func NewObjScrollData(src *types.Datum, _ *ObjScrollData, _ ...types.Value) {
	src.SetVar("icon", atoms.WorldOf(src).Icon("scroll.dmi"))
	src.SetVar("desc", types.String("It looks to be rather old."))
	src.SetVar("name", types.String("scroll"))
	src.SetVar("verbs", src.Var("verbs").Invoke(nil, "+", atoms.NewVerb("read", "/obj/scroll", "read")))
}

func (chunk *ObjScrollData) Procread(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	(varusr).Invoke(varusr, "<<", types.String("You utter the phrase written on the scroll: \"Knuth\"!"))
	_ = atoms.WorldOf(varsrc).Realm().New("/mob/rat", varusr, (varusr).Var("loc"))
	(varusr).Invoke(varusr, "<<", types.String("A giant rat appears!"))
	types.Del(varsrc)
	return out
}

func (*ObjScrollData) SettingsForProcread() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeUsr,
			In:   true,
		},
	}
}

//mediator:extend ExtAreaData /area
type ExtAreaData struct {
	VarMusic types.Value
}

func NewExtAreaData(src *types.Datum, _ *ExtAreaData, _ ...types.Value) {
}

func (chunk *ExtAreaData) ProcEntered(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	varm := types.Param(allargs, 0)
	if types.AsBool(procs.OperatorNot(procs.Invoke(atoms.WorldOf(varsrc), varusr, "ismob", varm))) {
		return out
	}
	(varm).Invoke(varusr, "<<", varsrc.Var("desc"))
	(varm).Invoke(varusr, "<<", procs.KWInvoke(atoms.WorldOf(varsrc), varusr, "sound", map[string]types.Value{"channel": types.Int(1)}, varsrc.Var("music"), types.Int(1)))
	return out
}

func (*ExtAreaData) SettingsForProcEntered() types.ProcSettings {
	return types.ProcSettings{
		Src: types.SrcSetting{
			Type: types.SrcSettingTypeView,
		},
	}
}

//mediator:declare AreaOutsideData /area/outside /area
type AreaOutsideData struct {
}

func NewAreaOutsideData(src *types.Datum, _ *AreaOutsideData, _ ...types.Value) {
	src.SetVar("desc", types.String("Nice and jazzy, here..."))
	src.SetVar("music", procs.NewSound("jazzy.mid"))
	src.SetVar("name", types.String("outside"))
}

//mediator:declare AreaCaveData /area/cave /area
type AreaCaveData struct {
}

func NewAreaCaveData(src *types.Datum, _ *AreaCaveData, _ ...types.Value) {
	src.SetVar("desc", types.String("Watch out for the giant rat!"))
	src.SetVar("music", procs.NewSound("cavern.mid"))
	src.SetVar("name", types.String("cave"))
}

//mediator:extend ExtAtomData /atom
type ExtAtomData struct {
}

func NewExtAtomData(src *types.Datum, _ *ExtAtomData, _ ...types.Value) {
}

func (chunk *ExtAtomData) ProcBumped(varsrc *types.Datum, varusr *types.Datum, allargs []types.Value) (out types.Value) {
	return out
}

func BeforeMap(world *world.World) []string {
	world.Name = "Your First World"
	world.Mob = "/mob/player"
	return []string{
		"map.dmm",
	}
}

func BuildWorld() *world.World {
	world, _ := framework.BuildWorld(Tree, BeforeMap)
	return world
}

func main() {
	framework.Launch(Tree, BeforeMap)
}
