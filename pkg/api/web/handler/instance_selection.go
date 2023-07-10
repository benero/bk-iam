/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package handler

import (
	"github.com/TencentBlueKing/gopkg/errorx"
	"github.com/gin-gonic/gin"

	"github.com/TencentBlueKing/bk-iam/pkg/service"
	"github.com/TencentBlueKing/bk-iam/pkg/util"
)

// ListInstanceSelection 查询系统的所有实例视图
func ListInstanceSelection(c *gin.Context) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf("Handler", "ListInstanceSelection")

	systemID := c.Param("system_id")

	// 获取action信息
	svc := service.NewInstanceSelectionService()
	instanceSelections, err := svc.ListBySystem(systemID)
	if err != nil {
		err = errorWrapf(err, "systemID=`%s`", systemID)
		util.SystemErrorJSONResponse(c, err)
		return
	}

	util.SuccessJSONResponse(c, "ok", instanceSelections)
}
