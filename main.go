package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/osm/mvdspawn/internal/mvdparser"
)

func main() {
	locsPath := flag.String("locs-path", "", "path to directory with loc files")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("usage: %s [-locs-path /quake/qw/locs] <demo.mvd>\n", os.Args[0])
		os.Exit(1)
	}
	mvdPath := args[0]

	mvdData, err := ioutil.ReadFile(mvdPath)
	if err != nil {
		fmt.Printf("unable to open %v, %v\n", mvdPath, err)
		os.Exit(1)
	}

	p := mvdparser.New()
	if *locsPath != "" {
		p.SetLocsPath(*locsPath)
	}

	spawns, err := p.Parse(mvdData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when parsing %v, %v\n", mvdPath, err)
		os.Exit(1)
	}

	for _, s := range spawns {
		fmt.Printf("%s;%s;%s;%s\n", p.Level, s.Player.Team, s.Player.Name, s.Location.Name)
	}
}
