package main

import (
	"fmt"
	"os"
	"time"

	"github.com/devops-works/slowql"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s file.log\n", os.Args[0])
		os.Exit(1)
	}

	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	p := slowql.NewParser(fd)

	var count int
	start := time.Now()
	for {
		q, err := p.GetNext()
		if err != nil {
			logrus.Error(err)
		}
		if q.Query == "" {
			break
		}
		fmt.Printf("Time: %s\nUser: %s\nHost: %s\nID: %d\nSchema: %s\nLast_errno: %d\nKilled: %d\nQuery_time: %s\nLock_time: %s\nRows_sent: %d\nRows_examined: %d\nRows_affected: %d\nBytes_sent: %d\nQuery: %s\n",
			q.Time,
			q.User,
			q.Host,
			q.ID,
			q.Schema,
			q.LastErrNo,
			q.Killed,
			q.QueryTime,
			q.LockTime,
			q.RowsSent,
			q.RowsExamined,
			q.RowsAffected,
			q.BytesSent,
			q.Query,
		)
		count++
	}

	elapsed := time.Since(start)
	fmt.Printf("parsed %d queries in %s\n", count, elapsed)
}
