/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package cacheimpls

import (
	"errors"

	"github.com/TencentBlueKing/gopkg/cache"

	"github.com/TencentBlueKing/bk-iam/pkg/service"
	"github.com/TencentBlueKing/bk-iam/pkg/service/types"
)

func retrieveSubject(key cache.Key) (interface{}, error) {
	k := key.(SubjectPKCacheKey)
	svc := service.NewSubjectService()
	return svc.Get(k.PK)
}

// GetSubjectByPK ...
func GetSubjectByPK(pk int64) (subject types.Subject, err error) {
	key := SubjectPKCacheKey{
		PK: pk,
	}

	var value interface{}
	value, err = LocalSubjectCache.Get(key)
	if err != nil {
		return
	}

	var ok bool
	subject, ok = value.(types.Subject)
	if !ok {
		err = errors.New("not types.Subject in cache")
		return
	}

	return
}
