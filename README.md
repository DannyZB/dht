# dht

[![Go Reference](https://pkg.go.dev/badge/github.com/dannyzb/dht/v2.svg)](https://pkg.go.dev/github.com/dannyzb/dht/v2)

dht implements the mainline DHT protocol used by BitTorrent. It provides server and client implementations, handling the
protocol level stuff, and allowing the user to provide their own store and routing table implementations.

## Installation

Get the library package with `go get github.com/dannyzb/dht/v2`, or the provided cmds with `go install github.com/dannyzb/dht/v2/cmd/...@latest`.

## Commands

Several commands are provided that use the library:

### dht

Supports various commands operating on the DHT.

    % go run github.com/dannyzb/dht/v2/cmd/dht --help
    valid arguments at this point:
      --help|-h
      --network <string>
      --secure
      --bootstrap-addr <[]string>
      --query-resend-delay <time.Duration>
      derive-put-target
      put
      put-mutable-infohash
      get
      ping
      get-peers
      query
      ping-nodes

## Downstream projects

Projects that uses this repo in novel ways.

* [cove](https://coveapp.info): Torrent browser with streaming, DHT search, video transcoding and casting.
* [btlink](https://github.com/anacrolix/btlink): btlink supports DNS records stored on the DHT.
