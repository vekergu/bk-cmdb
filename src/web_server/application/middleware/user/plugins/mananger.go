/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package plugins

import (
	"github.com/gin-gonic/gin"

	"configcenter/src/common"
	"configcenter/src/common/core/cc/api"
	"configcenter/src/common/metadata"
	"configcenter/src/web_server/application/middleware/user/plugins/manager"
	_ "configcenter/src/web_server/application/middleware/user/plugins/register"
)

func CurrentPlugin(c *gin.Context) metadata.LoginUserPluginInerface {
	ccapi := api.NewAPIResource()
	config, _ := ccapi.ParseConfig()
	version, ok := config["login.version"]
	if !ok {
		version = common.BKDefaultLoginUserPluginVersion
	}

	var selfPlugin *metadata.LoginPluginInfo
	for _, plugin := range manager.LoginPluginInfo {
		if plugin.Version == version {
			return plugin.HandleFunc
		}
		if common.BKDefaultLoginUserPluginVersion == plugin.Version {
			selfPlugin = plugin
		}
	}
	if nil != selfPlugin {
		return selfPlugin.HandleFunc
	}

	return nil
}
