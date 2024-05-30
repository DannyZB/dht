package krpc

import (
	"testing"

	"github.com/dannyzb/torrent/bencode"
	"github.com/stretchr/testify/require"
)

// https://github.com/dannyzb/torrent/issues/166
func TestUnmarshalBadError(t *testing.T) {
	var e Error
	err := bencode.Unmarshal([]byte(`l5:helloe`), &e)
	require.Error(t, err)
}
