/**
 * @Author: kwens
 * @Date: 2022-08-10 16:11:08
 * @Description:
 */
package oprlog

import (
	"io"
	"log"
	"os"
)

type OprLogRepository interface {
	Write(log OprLog) error
}

var (
	oprLogRepo OprLogRepository
)

func RegisterRepo(repo OprLogRepository) {
	oprLogRepo = repo
}

// defaultRepo 默认操作日志的仓储，直接输出到日志
type defaultRepo struct {
	log *log.Logger
}

func (r *defaultRepo) Write(log OprLog) error {
	r.log.Println(log.String())
	return nil
}

func defaultRepository() OprLogRepository {
	mw := io.MultiWriter(os.Stdout)
	l := log.Default()
	l.SetOutput(mw)
	return &defaultRepo{log: l}
}
