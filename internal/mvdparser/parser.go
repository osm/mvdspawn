package mvdparser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/osm/quake/common/args"
	"github.com/osm/quake/common/ascii"
	"github.com/osm/quake/common/context"
	"github.com/osm/quake/common/infostring"
	"github.com/osm/quake/common/loc"
	"github.com/osm/quake/demo/mvd"
	"github.com/osm/quake/packet/command/playerinfo"
	"github.com/osm/quake/packet/command/stufftext"
	"github.com/osm/quake/packet/command/updateuserinfo"
	"github.com/osm/quake/packet/svc"
	"github.com/osm/quake/protocol"
)

type Parser struct {
	isStarted bool
	locs      *loc.Locations
	locsPath  string

	Level   string
	Players map[byte]*Player
	Spawns  []Spawn
}

type Player struct {
	isDead bool

	Team string
	Name string
}

type Spawn struct {
	Location *loc.Location
	Player   *Player
}

func New() *Parser {
	return &Parser{
		Players: make(map[byte]*Player),
	}
}

func (p *Parser) SetLocsPath(locsPath string) {
	p.locsPath = locsPath
}

func (p *Parser) Parse(data []byte) ([]Spawn, error) {
	demo, err := mvd.Parse(context.New(), data)
	if err != nil {
		return nil, err
	}

	for _, d := range demo.Data {
		if d.Read == nil {
			continue
		}

		gd, ok := d.Read.Packet.(*svc.GameData)
		if !ok {
			continue
		}

		var err error
		for _, cmd := range gd.Commands {
			switch c := cmd.(type) {
			case *stufftext.Command:
				err = p.handleStufftext(c)
			case *updateuserinfo.Command:
				err = p.handleUpdateUserinfo(c)
			case *playerinfo.Command:
				err = p.handlePlayerinfo(c)
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return p.Spawns, nil
}

func (p *Parser) handleStufftext(cmd *stufftext.Command) error {
	for _, a := range args.Parse(cmd.String) {
		switch a.Cmd {
		case "//ktx":
			return p.handleKTX(a.Args)
		case "fullserverinfo":
			return p.handleFullServerInfo(a.Args)
		}
	}

	return nil
}

func (p *Parser) handleKTX(args []string) error {
	if len(args) != 1 || args[0] != "matchstart" {
		return nil
	}

	p.isStarted = true
	for _, pl := range p.Players {
		pl.isDead = true
	}

	return nil
}

func (p *Parser) handleFullServerInfo(args []string) error {
	if len(args) == 0 {
		return nil
	}

	info := infostring.Parse(args[0])
	p.Level = info.Get("map")

	return p.loadLocs()
}

func (p *Parser) loadLocs() error {
	var data []byte
	var err error

	if p.locsPath != "" {
		data, err = ioutil.ReadFile(filepath.Join(p.locsPath, fmt.Sprintf("%s.loc", p.Level)))
		if err != nil {
			return err
		}
	} else {
		data, _ = locs[p.Level]
	}

	if len(data) == 0 {
		return fmt.Errorf("unable to find loc data for %s", p.Level)
	}

	locs, err := loc.Parse(data)
	if err != nil {
		return err
	}

	p.locs = locs
	return nil
}

func (p *Parser) handleUpdateUserinfo(cmd *updateuserinfo.Command) error {
	is := infostring.Parse(cmd.UserInfo)
	name := ascii.Parse(is.Get("name"))
	if name == "" {
		return nil
	}

	if _, exists := p.Players[cmd.PlayerIndex]; !exists {
		p.Players[cmd.PlayerIndex] = &Player{}
	}

	pl := p.Players[cmd.PlayerIndex]
	pl.Name = name
	pl.Team = strings.TrimSpace(ascii.Parse(is.Get("team")))

	return nil
}

func (p *Parser) handlePlayerinfo(cmd *playerinfo.Command) error {
	if !p.isStarted {
		return nil
	}

	pl, exists := p.Players[cmd.Index]
	if !exists {
		return nil
	}

	if !pl.isDead && cmd.MVD != nil && cmd.MVD.Bits&protocol.DFDead != 0 {
		pl.isDead = true
	} else if pl.isDead && cmd.MVD != nil && cmd.MVD.Bits&protocol.DFDead == 0 {
		pl.isDead = false
		p.Spawns = append(p.Spawns, Spawn{p.locs.Get(cmd.MVD.Coord), pl})
	}

	return nil
}
