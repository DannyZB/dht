package peer_store

import (
	"github.com/dannyzb/torrent/metainfo"

	"github.com/dannyzb/dht/v2/krpc"
)

type InfoHash = metainfo.Hash

type Interface interface {
	AddPeer(InfoHash, krpc.NodeAddr)
	GetPeers(InfoHash) []krpc.NodeAddr
}
