//go:build !dev

package worker

import "log"

func (c *Controller) handleError(err error) {
	log.Println(err)
}

func Start() {}
