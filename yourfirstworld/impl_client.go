// Code generated by mediator boilerplate; DO NOT EDIT.
package main

import (
	"github.com/celskeggs/mediator/platform/datum"
	"github.com/celskeggs/mediator/platform/types"
	"github.com/celskeggs/mediator/platform/world"
)

type ClientImpl struct {
	world.ClientData
	datum.DatumData
}

func NewClient(realm *types.Realm, params ...types.Value) *types.Datum {
	i := &ClientImpl{}
	d := realm.NewDatum(i)
	datum.NewDatumData(d, &i.DatumData, params...)
	world.NewClientData(d, &i.ClientData, params...)
	return d
}

func (t *ClientImpl) Type() types.TypePath {
	return "/client"
}

func (t *ClientImpl) Var(src *types.Datum, name string) (types.Value, bool) {
	switch name {
	case "type":
		return types.TypePath("/client"), true
	case "parent_type":
		return types.TypePath("/datum"), true
	case "key":
		return types.String(t.ClientData.VarKey), true
	case "view":
		return types.Int(t.ClientData.VarView), true
	case "eye":
		return t.ClientData.GetEye(src), true
	case "mob":
		return t.ClientData.GetMob(src), true
	case "virtual_eye":
		return t.ClientData.GetVirtualEye(src), true
	default:
		return nil, false
	}
}

func (t *ClientImpl) SetVar(src *types.Datum, name string, value types.Value) types.SetResult {
	switch name {
	case "type":
		return types.SetResultReadOnly
	case "parent_type":
		return types.SetResultReadOnly
	case "key":
		t.ClientData.VarKey = types.Unstring(value)
		return types.SetResultOk
	case "view":
		t.ClientData.VarView = types.Unint(value)
		return types.SetResultOk
	case "eye":
		t.ClientData.SetEye(src, value)
		return types.SetResultOk
	case "mob":
		t.ClientData.SetMob(src, value)
		return types.SetResultOk
	case "virtual_eye":
		return types.SetResultReadOnly
	default:
		return types.SetResultNonexistent
	}
}

func (t *ClientImpl) Proc(src *types.Datum, usr *types.Datum, name string, params ...types.Value) (types.Value, bool) {
	switch name {
	case "<<":
		return t.ClientData.OperatorWrite(src, usr, types.Param(params, 0)), true
	case "Del":
		return t.ClientData.ProcDel(src, usr), true
	case "East":
		return t.ClientData.ProcEast(src, usr), true
	case "Move":
		return t.ClientData.ProcMove(src, usr, types.Param(params, 0), types.Param(params, 1)), true
	case "New":
		return t.ClientData.ProcNew(src, usr, types.Param(params, 0)), true
	case "North":
		return t.ClientData.ProcNorth(src, usr), true
	case "South":
		return t.ClientData.ProcSouth(src, usr), true
	case "West":
		return t.ClientData.ProcWest(src, usr), true
	default:
		return nil, false
	}
}

func (t *ClientImpl) Chunk(ref string) interface{} {
	switch ref {
	case "github.com/celskeggs/mediator/platform/world.ClientData":
		return &t.ClientData
	case "github.com/celskeggs/mediator/platform/datum.DatumData":
		return &t.DatumData
	default:
		return nil
	}
}
