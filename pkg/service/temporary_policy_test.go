/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package service

import (
	"errors"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/bk-iam/pkg/database"
	"github.com/TencentBlueKing/bk-iam/pkg/database/dao"
	"github.com/TencentBlueKing/bk-iam/pkg/database/dao/mock"
	"github.com/TencentBlueKing/bk-iam/pkg/service/types"
)

var _ = Describe("TemporaryPolicyService", func() {
	Describe("ListThinBySubjectAction", func() {
		var ctl *gomock.Controller
		BeforeEach(func() {
			ctl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctl.Finish()
		})

		It("manager.ListThinBySubjectAction fail", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().ListThinBySubjectAction(
				int64(1), int64(2), gomock.Any(),
			).Return(
				nil, errors.New("list fail"),
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			_, err := manager.ListThinBySubjectAction(1, 2)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "ListThinBySubjectAction")
		})

		It("ok", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().ListThinBySubjectAction(
				int64(1), int64(2), gomock.Any(),
			).Return(
				[]dao.ThinTemporaryPolicy{
					{
						PK: 1,
					},
				}, nil,
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			ps, err := manager.ListThinBySubjectAction(1, 2)
			assert.NoError(GinkgoT(), err)
			assert.Equal(GinkgoT(), []types.ThinTemporaryPolicy{{PK: 1}}, ps)
		})
	})

	Describe("ListByPKs", func() {
		var ctl *gomock.Controller
		BeforeEach(func() {
			ctl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctl.Finish()
		})

		It("ListByPKs fail", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().ListByPKs(
				[]int64{1},
			).Return(
				nil, errors.New("list fail"),
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			_, err := manager.ListByPKs([]int64{1})
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "ListByPKs")
		})

		It("ok", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().ListByPKs(
				[]int64{1},
			).Return(
				[]dao.TemporaryPolicy{{
					PK: 1,
				}}, nil,
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			ps, err := manager.ListByPKs([]int64{1})
			assert.NoError(GinkgoT(), err)
			assert.Equal(GinkgoT(), []types.TemporaryPolicy{{PK: 1}}, ps)
		})
	})

	Describe("Create", func() {
		var ctl *gomock.Controller
		BeforeEach(func() {
			ctl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctl.Finish()
		})

		It("Create fail", func() {
			policies := []types.TemporaryPolicy{{}}
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkCreateWithTx(
				gomock.Any(), []dao.TemporaryPolicy{{}},
			).Return(
				nil, errors.New("create fail"),
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			_, err := manager.Create(policies)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "fail")
		})

		It("ok", func() {
			policies := []types.TemporaryPolicy{{}}
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkCreateWithTx(
				gomock.Any(), []dao.TemporaryPolicy{{}},
			).Return(
				[]int64{1}, nil,
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			pks, err := manager.Create(policies)
			assert.NoError(GinkgoT(), err)
			assert.Equal(GinkgoT(), []int64{1}, pks)
		})
	})

	Describe("DeleteByPKs", func() {
		var ctl *gomock.Controller
		BeforeEach(func() {
			ctl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctl.Finish()
		})

		It("DeleteByPKs fail", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkDeleteByPKs(
				int64(1), []int64{1, 2},
			).Return(
				int64(0), errors.New("delete fail"),
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			err := manager.DeleteByPKs(int64(1), []int64{1, 2})
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "fail")
		})

		It("ok", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkDeleteByPKs(
				int64(1), []int64{1, 2},
			).Return(
				int64(2), nil,
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			err := manager.DeleteByPKs(int64(1), []int64{1, 2})
			assert.NoError(GinkgoT(), err)
		})
	})
	Describe("DeleteBeforeExpiredAt", func() {
		var ctl *gomock.Controller
		BeforeEach(func() {
			ctl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctl.Finish()
		})

		It("BulkDeleteBeforeExpiredAtWithTx fail", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkDeleteBeforeExpiredAtWithTx(
				gomock.Any(), int64(1), int64(10000),
			).Return(
				int64(0), errors.New("delete fail"),
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			err := manager.DeleteBeforeExpiredAt(int64(1))
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "fail")
		})

		It("ok", func() {
			mockTemporaryPolicyService := mock.NewMockTemporaryPolicyManager(ctl)
			mockTemporaryPolicyService.EXPECT().BulkDeleteBeforeExpiredAtWithTx(
				gomock.Any(), int64(1), int64(10000),
			).Return(
				int64(1), nil,
			).AnyTimes()

			manager := &temporaryPolicyService{
				manager: mockTemporaryPolicyService,
			}

			db, dbMock := database.NewMockSqlxDB()
			dbMock.ExpectBegin()
			dbMock.ExpectCommit()

			patches := gomonkey.ApplyFunc(database.GenerateDefaultDBTx, db.Beginx)
			defer patches.Reset()

			err := manager.DeleteBeforeExpiredAt(int64(1))
			assert.NoError(GinkgoT(), err)
		})
	})
})
