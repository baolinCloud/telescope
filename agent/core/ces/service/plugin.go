package services

import (
	"github.com/huaweicloud/telescope/agent/core/logs"

	"github.com/huaweicloud/telescope/agent/core/ces/config"
	"github.com/huaweicloud/telescope/agent/core/ces/model"
	"github.com/huaweicloud/telescope/agent/core/ces/report"
	cesUtils "github.com/huaweicloud/telescope/agent/core/ces/utils"
)

// CollectPluginTask cron job for collecting plugin data
func CollectPluginTask(data chan *model.InputMetric) {
	if !config.GetConfig().Enable || !config.GetConfig().EnablePlugin {
		return
	}

	if config.GetDefaultPluginConfig() == nil {
		return
	}

	plugins := config.GetDefaultPluginConfig()

	if len(plugins) > cesUtils.MaxPluginNum {
		plugins = plugins[:cesUtils.MaxPluginNum]
	}

	for _, eachPlugin := range plugins {
		logs.GetCesLogger().Debugf("Plugin is default type, info is %v", *eachPlugin)

		eachPluginSchedule := model.NewPluginScheduler(eachPlugin)
		if eachPluginSchedule == nil {
			return
		}
		go eachPluginSchedule.Schedule(data)
	}
}

// SendPluginTask task for post plugin data
func SendPluginTask(data chan *model.InputMetric) {
	for {

		pluginData := <-data
		logs.GetCesLogger().Debugf("Plugin metric data is %v", *pluginData)
		report.SendMetricData(BuildURL(cesUtils.PostAggregatedMetricDataURI), pluginData, true)

	}
}
