package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/eggysetiawan/fiber-starterkit/errs"
)

func NewLogFile() (*os.File, *errs.AppError) {
	d := time.Now().Format(time.DateOnly)

	file, err := os.OpenFile(fmt.Sprintf("./storage/logs/ejol-%s.log", d), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	return file, nil
}

func LogToElastic() {

}
