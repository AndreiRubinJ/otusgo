package main

import (
	"strings"

	"github.com/AndreiRubinJ/otusgo/hw12_log_util/config"
	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log"
)

func main() {
	log.GenerateLogs()
	var statistics map[string]map[string]int
	conf := config.GetConfig()
	println(conf.LogFile, conf.LogLevel, conf.Output)
	if conf.LogFile == "" {
		log.GenerateLogs()
		statistics = log.GetStatistic(log.GetFileNameByDefoult(), strings.ToUpper(conf.LogLevel))
	} else {
		statistics = log.GetStatistic(conf.LogFile, strings.ToUpper(conf.LogLevel))
	}

	if conf.Output == "" {
		log.OutputToConsole(statistics)
	} else {
		log.OutputStatisticToFile(statistics, conf.Output)
	}
}
