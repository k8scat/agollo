// Copyright 2025 Apollo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extension

import (
	"testing"

	. "github.com/tevid/gohamcrest"

	"github.com/apolloconfig/agollo/v4/env/config"
)

type TestFileHandler struct {
}

// WriteConfigFile 写入配置文件
func (r *TestFileHandler) WriteConfigFile(config *config.ApolloConfig, configPath string) error {
	return nil
}

// GetConfigFile 获得配置文件路径
func (r *TestFileHandler) GetConfigFile(configDir string, appID string, cluster string, namespace string) string {
	return ""
}

func (r *TestFileHandler) LoadConfigFile(configDir string, appID string, cluster string, namespace string) (*config.ApolloConfig, error) {
	return nil, nil
}

func TestSetFileHandler(t *testing.T) {
	SetFileHandler(&TestFileHandler{})

	fileHandler := GetFileHandler()

	Assert(t, fileHandler, NotNilVal())
}
