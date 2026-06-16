package remote

import (
	"time"
)

func (t *queueManager) runShard(shardID int, stop chan struct{}) {
	backoff := time.Second
	for {
		select {
		case <-stop:
			return
		default:
			// Sending logic here
			err := t.sendSamples()
			if err != nil && isRetriable(err) {
				select {
				case <-stop:
					return
				case <-time.After(backoff):
					backoff *= 2
					continue
				}
			}
			backoff = time.Second
		}
	}
}
