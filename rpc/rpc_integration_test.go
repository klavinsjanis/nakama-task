package rpc

import (
	"context"
	"log"
	"testing"

	"github.com/ascii8/nakama-go"
	"github.com/stretchr/testify/require"
)

const (
	id       = "00000000-0000-0000-0000-000000000000"
	severkey = "defaultkey"
)

var client *nakama.Client

func TestMain(m *testing.M) {
	// create client
	client = nakama.New(nakama.WithServerKey(severkey))

	m.Run()
}

func Test_GetFileContents(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// authenticate
		if err := client.AuthenticateDevice(ctx, id, true, ""); err != nil {
			log.Fatal(err)
		}

		hash, err := generateMD5Hash("../files/core/0.0.0.json")
		require.NoError(t, err)

		var resp getFileContentsResp

		err = client.Rpc(ctx, "getfilecontents", getFileContentsPayload{
			Type:    ptr("core"),
			Version: ptr("0.0.0"),
			Hash:    &hash,
		}, &resp)

		require.NoError(t, err)
		require.NotEmpty(t, resp)
	})

	t.Run("bad request, file doesn't exist", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// authenticate
		if err := client.AuthenticateDevice(ctx, id, true, ""); err != nil {
			log.Fatal(err)
		}

		var resp getFileContentsResp

		err := client.Rpc(ctx, "getfilecontents", getFileContentsPayload{
			Type:    ptr("xyz"),
			Version: ptr("1.2.3"),
		}, &resp)

		require.Error(t, err)
		require.Empty(t, resp)
	})
}

// helpers

// ptr returns pointer to the passed in value.
func ptr[T any](v T) *T {
	return &v
}
