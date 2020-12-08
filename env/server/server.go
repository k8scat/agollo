/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package serverlist

import (
	"github.com/zouyx/agollo/v4/env/config"
	"strings"
	"sync"
	"time"
)

// ip -> server
var (
	ipMap      map[string]*Server
	serverLock sync.Mutex
	//next try connect period - 60 second
	nextTryConnectPeriod int64 = 30
)

type Server struct {
	//real servers ip
	serverMap       map[string]*config.ServerInfo
	nextTryConnTime int64
}

//GetServersLen 获取服务器数组
func GetServers(configIp string) map[string]*config.ServerInfo {
	serverLock.Lock()
	defer serverLock.Unlock()
	if ipMap[configIp] == nil {
		return nil
	}
	return ipMap[configIp].serverMap
}

//GetServersLen 获取服务器数组长度
func GetServersLen(configIp string) int {
	serverLock.Lock()
	defer serverLock.Unlock()
	s := ipMap[configIp]
	if s == nil || len(s.serverMap) == 0 {
		return 0
	}
	return len(s.serverMap)
}

func SetServers(configIp string, serverMap map[string]*config.ServerInfo) {
	serverLock.Lock()
	defer serverLock.Unlock()
	ipMap[configIp] = &Server{
		serverMap: serverMap,
	}
}

//SetDownNode 设置失效节点
func SetDownNode(host string, configIp string) {
	serverLock.Lock()
	defer serverLock.Unlock()
	s := ipMap[configIp]
	if host == "" || s == nil {
		return
	}

	if host == configIp {
		s.nextTryConnTime = nextTryConnectPeriod
	}

	for k, server := range s.serverMap {
		// if some node has down then select next node
		if strings.Index(k, host) > -1 {
			server.IsDown = true
		}
	}
}

//IsConnectDirectly is connect by ip directly
//false : no
//true : yes
func IsConnectDirectly(configIp string) bool {
	serverLock.Lock()
	defer serverLock.Unlock()
	s := ipMap[configIp]
	if s.nextTryConnTime >= 0 && s.nextTryConnTime > time.Now().Unix() {
		return true
	}

	return false
}

//SetNextTryConnTime if this connect is fail will set this time
func setNextTryConnTime(configIp string) {
	serverLock.Lock()
	defer serverLock.Unlock()
	s := ipMap[configIp]
	if s == nil || len(s.serverMap) == 0 {
		return
	}
	s.nextTryConnTime = time.Now().Unix() + nextTryConnectPeriod
}
