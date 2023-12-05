/**
 * @Author: kwens
 * @Date: 2022-10-09 17:39:09
 * @Description:
 */
package cron

import (
	cron2 "github.com/robfig/cron/v3"

	"sync"
)

type cron struct {
	cronMap map[string]cron2.EntryID
	mutex   sync.Mutex
	cron    *cron2.Cron
}

func New() *cron {
	return &cron{
		cronMap: make(map[string]cron2.EntryID),
		cron:    cron2.New(cron2.WithSeconds()),
	}
}

// 0 */30 * * * * 每隔30分
// 30 59 23 1 3 * 3月1号23:29:30秒
// 30 59 23 * * 7 周日的23:29:30秒
func (c *cron) Add(cronName string, spec string, cmd func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cronMap[cronName] > 0 {
		c.cron.Remove(c.cronMap[cronName])
	}
	id, err := c.cron.AddFunc(spec, cmd)
	if err == nil {
		c.cronMap[cronName] = id
	}
	return err
}

func (c *cron) Delete(cronName string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cronMap[cronName] > 0 {
		c.cron.Remove(c.cronMap[cronName])
		delete(c.cronMap, cronName)
	}
}

func (c *cron) Start() {
	c.cron.Start()
}

var (
	Cron *cron
)

func Start() {
	Cron = New()
	Cron.Start()
}
