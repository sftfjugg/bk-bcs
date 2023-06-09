/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package k8sclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	clusternet "github.com/clusternet/clusternet/pkg/generated/clientset/versioned"
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	k8sVersion "k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/storage"
)

// GetEnvByClusterId 获取集群所属环境, 目前通过集群ID前缀判断
func GetEnvByClusterId(clusterId string) config.BCSClusterEnv {
	if strings.HasPrefix(clusterId, "BCS-K8S-1") {
		return config.UatCluster
	}
	if strings.HasPrefix(clusterId, "BCS-K8S-2") {
		return config.DebugCLuster
	}
	if strings.HasPrefix(clusterId, "BCS-K8S-4") {
		return config.ProdEnv
	}
	return config.ProdEnv
}

// GetBCSConfByClusterId 通过集群ID, 获取不同admin token 信息
func GetBCSConfByClusterId(clusterId string) *config.BCSConf {
	env := GetEnvByClusterId(clusterId)
	conf, ok := config.G.BCSEnvMap[env]
	if ok {
		return conf
	}
	// 默认返回bcs配置
	return config.G.BCS
}

// GetK8SClientByClusterId 通过集群 ID 获取 k8s client 对象
func GetK8SClientByClusterId(clusterId string) (*kubernetes.Clientset, error) {
	bcsConf := GetBCSConfByClusterId(clusterId)
	host := fmt.Sprintf("%s/clusters/%s", bcsConf.Host, clusterId)
	config := &rest.Config{
		Host:        host,
		BearerToken: bcsConf.Token,
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return k8sClient, nil
}

// GetClusterNetClientByClusterId 通过集群 ID 获取 clusternet client 对象
func GetClusterNetClientByClusterId(clusterId string) (*clusternet.Clientset, error) {
	bcsConf := GetBCSConfByClusterId(clusterId)
	host := fmt.Sprintf("%s/clusters/%s", bcsConf.Host, clusterId)
	config := &rest.Config{
		Host:        host,
		BearerToken: bcsConf.Token,
	}
	k8sClient, err := clusternet.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return k8sClient, nil
}

// parseVersion 解析版本, 优先使用gitVersion, 回退到 v{Major}.{Minor}.0
func parseVersion(versionInfo *k8sVersion.Info) (*version.Version, error) {
	v, err := version.NewVersion(versionInfo.GitVersion)
	if err == nil {
		return v, nil
	}

	// 回退使用 v{Major}.{Minor}.0
	majorVersion := fmt.Sprintf("v%s.%s.0", versionInfo.Major, versionInfo.Minor)
	v, err = version.NewVersion(majorVersion)
	if err == nil {
		return v, nil
	}

	return nil, errors.Errorf("Malformed version: GitVersion %s, MajorVersion %s", versionInfo.GitVersion, majorVersion)
}

// GetK8SVersion 获取 k8s 版本
func GetK8SVersion(ctx context.Context, clusterId string) (*version.Version, error) {
	cacheKey := fmt.Sprintf("k8sclient.GetK8SVersion:%s", clusterId)
	if cacheResult, ok := storage.LocalCache.Slot.Get(cacheKey); ok {
		return cacheResult.(*version.Version), nil
	}

	k8sClient, err := GetK8SClientByClusterId(clusterId)
	if err != nil {
		return nil, err
	}

	info, err := k8sClient.ServerVersion()
	if err != nil {
		return nil, err
	}

	v, err := parseVersion(info)
	if err != nil {
		return nil, err
	}

	// 缓存2个小时
	storage.LocalCache.Slot.Set(cacheKey, v, time.Hour*2)

	return v, nil
}

// K8SLessThan 对比版本, 如果异常, 按向下兼容
func K8SLessThan(ctx context.Context, clusterId, ver string) bool {
	k8sV, err := GetK8SVersion(ctx, clusterId)
	if err != nil {
		klog.Warningf("GetK8SVersion error, %s, %s, %s", clusterId, err, err == nil)
		return false
	}

	rawV, err := version.NewVersion(ver)
	if err != nil {
		klog.Warningf("parse raw version, %s, %s, %s", clusterId, ver, err)
		return false
	}

	return k8sV.LessThan(rawV)
}
