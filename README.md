# Single process

> [fork from postfinance/single](https://github.com/postfinance/single)

## Options

`singleproc.WithLockPath` - configures the path for the lockfile.

`singleproc.WithPidPath` - configures the path for the pidfile.

## Usage

```go
package main

import (
	"github.com/NovikovRoman/singleproc"
	"log"
	"os"
	"time"
)

func main() {
	var (
		err error
		s   *singleproc.Single
	)

	s, err = singleproc.New("your-app-name")
	if err != nil {
		log.Fatalf("failed create Single: %v", err)
	}

	err = s.Lock()

	if err == singleproc.ErrAlreadyRunning {
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}

	defer func(s *singleproc.Single) {
		if err = s.Unlock(); err != nil {
			log.Fatal(err)
		}
	}(s)

	log.Println("working")
	time.Sleep(60 * time.Second)
	log.Println("finished")
}
```