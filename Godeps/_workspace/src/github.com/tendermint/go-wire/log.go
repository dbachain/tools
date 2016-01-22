package wire

import (
	"github.com/tendermint/netmon/Godeps/_workspace/src/github.com/tendermint/go-logger"
)

var log = logger.New("module", "binary")

func init() {
	log.SetHandler(
		logger.LvlFilterHandler(
			logger.LvlWarn,
			//logger.LvlDebug,
			logger.MainHandler(),
		),
	)
}