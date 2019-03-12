package main

import (
	"context"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"os"
	"time"
)

func main() {
	// Create a binlog syncer with a unique server id, the server id must be different from other MySQL's.
	// flavor is mysql or mariadb
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// Start sync with specified binlog file and position
	streamer, _ := syncer.StartSync(mysql.Position{binlogFile, binlogPos})

	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"

	for {
		ev, _ := streamer.GetEvent(context.Background())
		// Dump event
		ev.Dump(os.Stdout)
	}

	// or we can use a timeout context
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		ev, err := s.GetEvent(ctx)
		cancel()

		if err == context.DeadlineExceeded {
			// meet timeout
			continue
		}

		ev.Dump(os.Stdout)
	}
}
