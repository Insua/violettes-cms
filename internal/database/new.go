package database

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"path/filepath"
	"time"
)

var (
	errNullFileName = errors.New("need migration file name")
)

func newMigration(fileName string) error {
	dir := "internal/database/migrations"
	if len(fileName) == 0 {
		return errNullFileName
	}

	timeStamp := fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), nanoString())

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%s_%s.%s.%s", timeStamp, fileName, direction, "sql")
		filename := filepath.Join(dir, basename)

		if err := createFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func nanoString() string {
	nano := gconv.String(time.Now().Nanosecond())
	if len(nano) > 4 {
		return nano[0:4]
	} else {
		str := nano[0:]
		for i := 0; i < 4-len(nano); i++ {
			str += "0"
		}
		return str
	}
}

func createFile(filename string) error {
	fmt.Println(filename)
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	err = os.Chmod(filename, 0666)
	if err != nil {
		return err
	}

	return f.Close()
}
