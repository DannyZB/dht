package krpc

import (
	"github.com/dannyzb/torrent/metainfo"
)

type Bep46Payload struct {
	Ih metainfo.Hash `bencode:"ih"`
}
