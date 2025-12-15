package eventbus

import (
	"github.com/ProgrammerBuffalo/event-driven/exchange"
)

type EventBus struct {
	exchange exchange.Exchange
}
