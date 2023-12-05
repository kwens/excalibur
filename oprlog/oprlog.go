/**
 * @Author: kwens
 * @Date: 2022-08-10 15:51:18
 * @Description:
 */
package oprlog

import (
	"fmt"
	"sync"
	"time"
)

type OprLog struct {
	ID        string
	OprModule string
	OprType   string
	Operator  string
	OprTime   string
	OprIp     string
	OprValue  string
	OriValue  string
}

func NewOprLog(
	OprModule string,
	OprType string,
	Operator string,
	OprIp string,
	OprValue string,
	OriValue string,
) OprLog {
	id := genUuid()
	return OprLog{
		ID:        id,
		OprModule: OprModule,
		OprType:   OprType,
		Operator:  Operator,
		OprTime:   fmt.Sprintf("%d", time.Now().Unix()),
		OprIp:     OprIp,
		OprValue:  OprValue,
		OriValue:  OriValue,
	}
}

func (l OprLog) String() string {
	return fmt.Sprintf(
		"OPRLOG: [OPRTIME]%s [ID]%s [MODULE]%s [TYPE]%s [OPERATOR]%s [OPRIP]%s \n[OPRVALUE]%s \n[ORIVALUE]%s",
		l.OprTime,
		l.ID,
		l.OprModule,
		l.OprType,
		l.Operator,
		l.OprIp,
		l.OprValue,
		l.OriValue,
	)
}

type OprValue struct {
	OprModule string
	OprType   string
}

type OprLogClient struct {
	OprLogChan chan OprLog
	repo       OprLogRepository
	mu         sync.Mutex

	OprHandle map[string]OprValue
}

func NewOprLogClient() *OprLogClient {
	return &OprLogClient{
		OprLogChan: make(chan OprLog, 1024),
		OprHandle:  make(map[string]OprValue),
	}
}

func (c *OprLogClient) Write(log OprLog) {
	c.OprLogChan <- log
}

func (c *OprLogClient) RegisterValue(k string, v OprValue) {
	c.OprHandle[k] = v
}

func (c *OprLogClient) GetValue(k string) OprValue {
	if v, ok := c.OprHandle[k]; ok {
		return v
	}
	return OprValue{}
}

func (c *OprLogClient) watch() {
	for oprlog := range c.OprLogChan {
		c.mu.Lock()
		c.repo.Write(oprlog)
		c.mu.Unlock()
	}
}

func (c *OprLogClient) applyRepo() {
	if oprLogRepo == nil {
		c.repo = defaultRepository()
		return
	}
	c.repo = oprLogRepo
}

var (
	once      sync.Once
	OprLogCli *OprLogClient
)

func Start() {
	once.Do(func() {
		if OprLogCli == nil {
			OprLogCli = NewOprLogClient()
		}
		OprLogCli.applyRepo()
		go OprLogCli.watch()
	})
}
