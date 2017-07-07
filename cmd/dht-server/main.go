package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/anacrolix/dht"
	"github.com/anacrolix/dht/krpc"
	"github.com/anacrolix/tagflag"
)

var (
	flags = struct {
		TableFile string `help:"name of file for storing node info"`
		Addr      string `help:"local UDP address"`
	}{
		Addr: ":0",
	}
	s *dht.Server
)

func loadTable() error {
	if flags.TableFile == "" {
		return nil
	}
	f, err := os.Open(flags.TableFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("error opening table file: %s", err)
	}
	defer f.Close()
	added := 0
	for {
		b := make([]byte, krpc.CompactIPv4NodeInfoLen)
		_, err := io.ReadFull(f, b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading table file: %s", err)
		}
		var ni krpc.NodeInfo
		err = ni.UnmarshalCompactIPv4(b)
		if err != nil {
			return fmt.Errorf("error unmarshaling compact node info: %s", err)
		}
		s.AddNode(ni)
		added++
	}
	log.Printf("loaded %d nodes from table file", added)
	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tagflag.Parse(&flags)
	var err error
	s, err = dht.NewServer(&dht.ServerConfig{
		Addr: flags.Addr,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = loadTable()
	if err != nil {
		log.Fatalf("error loading table: %s", err)
	}
	log.Printf("dht server on %s, ID is %x", s.Addr(), s.ID())
}

func saveTable() error {
	goodNodes := s.Nodes()
	if flags.TableFile == "" {
		if len(goodNodes) != 0 {
			log.Printf("discarding %d good nodes!", len(goodNodes))
		}
		return nil
	}
	f, err := os.OpenFile(flags.TableFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("error opening table file: %s", err)
	}
	defer f.Close()
	for _, nodeInfo := range goodNodes {
		var b [krpc.CompactIPv4NodeInfoLen]byte
		err := nodeInfo.PutCompact(b[:])
		if err != nil {
			return fmt.Errorf("error compacting node info: %s", err)
		}
		_, err = f.Write(b[:])
		if err != nil {
			return fmt.Errorf("error writing compact node info: %s", err)
		}
	}
	log.Printf("saved %d nodes to table file", len(goodNodes))
	return nil
}

func setupSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch)
	<-ch
	s.Close()

	if flags.TableFile != "" {
		if err := saveTable(); err != nil {
			log.Printf("error saving node table: %s", err)
		}
	}
}

func main() {
	setupSignals()
}