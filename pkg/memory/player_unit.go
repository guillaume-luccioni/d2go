package memory

import (
	"github.com/guillaume-luccioni/d2go/pkg/data"
	"github.com/guillaume-luccioni/d2go/pkg/data/area"
	"github.com/guillaume-luccioni/d2go/pkg/data/mode"
	"github.com/guillaume-luccioni/d2go/pkg/data/stat"
	"github.com/guillaume-luccioni/d2go/pkg/data/state"
)

type RawPlayerUnit struct {
	data.UnitID
	Address      uintptr
	Name         string
	IsMainPlayer bool
	IsCorpse     bool
	Area         area.ID
	Position     data.Position
	IsHovered    bool
	States       state.States
	Stats        stat.Stats
	BaseStats    stat.Stats
	Mode         mode.PlayerMode
}

type RawPlayerUnits []RawPlayerUnit

func (pu RawPlayerUnits) GetMainPlayer() RawPlayerUnit {
	for _, p := range pu {
		if p.IsMainPlayer {
			return p
		}
	}
	return RawPlayerUnit{}
}

func (pu RawPlayerUnits) GetCorpse() RawPlayerUnit {
	for _, p := range pu {
		if p.IsCorpse {
			return p
		}
	}
	return RawPlayerUnit{}
}
