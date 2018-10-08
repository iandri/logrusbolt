# Logrus BoltDB Hook

With this hook logrus saves messages in the [BoltDB](https://github.com/coreos/bbolt)

## Install

```bash
$ go get github.com/iandri/logrusbolt
```

## Usage

```go
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/iandri/logrusbolt"
)

func init() {
	config := logrusbolt.BoltHook{
		DBLoc:     "/tmp/test.db",
		Bucket:    "test",
		Formatter: &logrus.JSONFormatter{},
        Level:     logrus.WarnLevel,
	}
	
	hook, err := logrusbolt.NewHook(config)
	
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Error(err)
	}
}


func main() {
	logrus.Info("test info")
}
```

