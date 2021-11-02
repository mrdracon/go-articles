package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(... string)){
	t.Helper()

	//testDSN := "postgres://127.0.0.1/go-url-shortener-test?sslmode=disable&user=postgres&password=postgres"
	st := Store{
		DatabaseURL:    databaseURL,
	}

	err := st.Open()
	if err != nil {
		t.Fatal(err)
	}

	return &st, func(tables ... string){
		if len(tables) > 0 {
			if _, err := st.dbConn.Exec(st.dbCtx, fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
			st.Close()
		}
	}
}
