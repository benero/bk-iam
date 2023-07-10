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
	"github.com/TencentBlueKing/gopkg/collection/set"
	"github.com/gin-gonic/gin"

	"github.com/TencentBlueKing/bk-iam/pkg/util"
)

// ShareSystemInfoQuery godoc
// @Summary system info query
// @Description query the system
// @ID api-model-share-system-query
// @Tags model_share
// @Accept json
// @Produce json
// @Param system_id path string true "System ID"
// @Param body body querySerializer true "the policy request"
// @Success 200 {object} util.Response{data=gin.H}
// @Header 200 {string} X-Request-Id "the request id"
// @Security AppCode
// @Security AppSecret
// @Router /api/v1/model/share/systems/{system_id}/query [get]
func ShareSystemInfoQuery(c *gin.Context) {
	var query querySerializer
	if err := c.ShouldBindQuery(&query); err != nil {
		util.BadRequestErrorJSONResponse(c, util.ValidationErrorMessage(err))
		return
	}

	systemID := c.Param("system_id")
	fields := query.Fields
	if fields == "" {
		fields = "base_info,resource_types,actions,action_groups,instance_selections,resource_creator_actions," +
			"common_actions,feature_shield_rules"
	}
	fieldSet := set.SplitStringToSet(fields, ",")

	BuildSystemInfoQueryResponse(c, systemID, fieldSet, true)
}
