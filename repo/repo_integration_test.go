package repo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	driver = "postgres"

	host     = "localhost"
	user     = "postgres"
	password = "localdb"
	dbname   = "nakama"

	port = 5432
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open(driver,
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	m.Run()
}

func Test_SaveRequestDetails(t *testing.T) {
	t.Parallel()

	id, err := SaveUserRequestDetails(context.Background(), db, gofakeit.UUID(), gofakeit.BuzzWord(), time.Now())

	require.Nil(t, err)
	require.NotEmpty(t, id)
}
