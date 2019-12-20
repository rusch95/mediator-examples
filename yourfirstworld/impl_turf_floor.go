// Code generated by mediator boilerplate; DO NOT EDIT.
package main

import (
	"github.com/celskeggs/mediator/common"
	"github.com/celskeggs/mediator/platform/atoms"
	"github.com/celskeggs/mediator/platform/datum"
	"github.com/celskeggs/mediator/platform/types"
)

type TurfFloorImpl struct {
	TurfFloorData
	atoms.TurfData
	atoms.AtomData
	datum.DatumData
}

func NewTurfFloor(realm *types.Realm, params ...types.Value) *types.Datum {
	i := &TurfFloorImpl{}
	d := realm.NewDatum(i)
	datum.NewDatumData(d, &i.DatumData, params...)
	atoms.NewAtomData(d, &i.AtomData, params...)
	atoms.NewTurfData(d, &i.TurfData, params...)
	NewTurfFloorData(d, &i.TurfFloorData, params...)
	return d
}

func (t *TurfFloorImpl) Type() types.TypePath {
	return "/turf/floor"
}

func (t *TurfFloorImpl) Var(src *types.Datum, name string) (types.Value, bool) {
	switch name {
	case "type":
		return types.TypePath("/turf/floor"), true
	case "parent_type":
		return types.TypePath("/turf"), true
	case "appearance":
		return t.AtomData.VarAppearance, true
	case "density":
		return types.Int(t.AtomData.VarDensity), true
	case "dir":
		return t.AtomData.VarDir, true
	case "opacity":
		return types.Int(t.AtomData.VarOpacity), true
	case "verbs":
		return datum.NewListFromSlice(t.AtomData.VarVerbs), true
	case "contents":
		return t.AtomData.GetContents(src), true
	case "desc":
		return t.AtomData.GetDesc(src), true
	case "icon":
		return t.AtomData.GetIcon(src), true
	case "icon_state":
		return t.AtomData.GetIconState(src), true
	case "layer":
		return t.AtomData.GetLayer(src), true
	case "loc":
		return t.AtomData.GetLoc(src), true
	case "name":
		return t.AtomData.GetName(src), true
	case "suffix":
		return t.AtomData.GetSuffix(src), true
	case "x":
		return t.TurfData.GetX(src), true
	case "y":
		return t.TurfData.GetY(src), true
	case "z":
		return t.TurfData.GetZ(src), true
	default:
		return nil, false
	}
}

func (t *TurfFloorImpl) SetVar(src *types.Datum, name string, value types.Value) types.SetResult {
	switch name {
	case "type":
		return types.SetResultReadOnly
	case "parent_type":
		return types.SetResultReadOnly
	case "appearance":
		t.AtomData.VarAppearance = value.(atoms.Appearance)
		return types.SetResultOk
	case "density":
		t.AtomData.VarDensity = types.Unint(value)
		return types.SetResultOk
	case "dir":
		t.AtomData.VarDir = value.(common.Direction)
		return types.SetResultOk
	case "opacity":
		t.AtomData.VarOpacity = types.Unint(value)
		return types.SetResultOk
	case "verbs":
		t.AtomData.VarVerbs = datum.ElementsAsType([]atoms.Verb{}, value).([]atoms.Verb)
		return types.SetResultOk
	case "contents":
		return types.SetResultReadOnly
	case "desc":
		t.AtomData.SetDesc(src, value)
		return types.SetResultOk
	case "icon":
		t.AtomData.SetIcon(src, value)
		return types.SetResultOk
	case "icon_state":
		t.AtomData.SetIconState(src, value)
		return types.SetResultOk
	case "layer":
		t.AtomData.SetLayer(src, value)
		return types.SetResultOk
	case "loc":
		t.AtomData.SetLoc(src, value)
		return types.SetResultOk
	case "name":
		t.AtomData.SetName(src, value)
		return types.SetResultOk
	case "suffix":
		t.AtomData.SetSuffix(src, value)
		return types.SetResultOk
	case "x":
		t.TurfData.SetX(src, value)
		return types.SetResultOk
	case "y":
		t.TurfData.SetY(src, value)
		return types.SetResultOk
	case "z":
		t.TurfData.SetZ(src, value)
		return types.SetResultOk
	default:
		return types.SetResultNonexistent
	}
}

func (t *TurfFloorImpl) Proc(src *types.Datum, usr *types.Datum, name string, params ...types.Value) (types.Value, bool) {
	switch name {
	case "Bump":
		return t.AtomData.ProcBump(src, usr, types.Param(params, 0)), true
	case "Enter":
		return t.TurfData.ProcEnter(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "Entered":
		return t.TurfData.ProcEntered(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "Exit":
		return t.TurfData.ProcExit(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "Exited":
		return t.TurfData.ProcExited(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "Move":
		return t.AtomData.ProcMove(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "New":
		return t.DatumData.ProcNew(src, usr), true
	default:
		return nil, false
	}
}

func (t *TurfFloorImpl) ProcSettings(name string) (types.ProcSettings, bool) {
	switch name {
	case "Bump":
		return types.ProcSettings{}, true
	case "Enter":
		return types.ProcSettings{}, true
	case "Entered":
		return types.ProcSettings{}, true
	case "Exit":
		return types.ProcSettings{}, true
	case "Exited":
		return types.ProcSettings{}, true
	case "Move":
		return types.ProcSettings{}, true
	case "New":
		return types.ProcSettings{}, true
	default:
		return types.ProcSettings{}, false
	}
}

func (t *TurfFloorImpl) Chunk(ref string) interface{} {
	switch ref {
	case "github.com/celskeggs/mediator-examples/yourfirstworld.TurfFloorData":
		return &t.TurfFloorData
	case "github.com/celskeggs/mediator/platform/atoms.TurfData":
		return &t.TurfData
	case "github.com/celskeggs/mediator/platform/atoms.AtomData":
		return &t.AtomData
	case "github.com/celskeggs/mediator/platform/datum.DatumData":
		return &t.DatumData
	default:
		return nil
	}
}
