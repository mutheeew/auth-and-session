package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	databaseUrl := "postgres://postgres:181818@localhost:5432/Personal-Web-Mute"

	var err error
	//
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	// panggil kembali pgx, isinya context backgroun supaya server berjalan tidak sekali tapi terus menerus terhubung ke database selama aplikasi berjalan
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connert database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success connect to database")
}
