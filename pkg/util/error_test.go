package util_test

import (
	. "github.com/onsi/ginkgo/v2"

	"github.com/TencentBlueKing/bk-iam/pkg/util"
)

var _ = Describe("Error", func() {
	BeforeEach(func() {
		util.InitErrorReport(false)
	})

	Describe("ReportToSentry", func() {
		It("send", func() {
			util.ReportToSentry("test", map[string]interface{}{
				"hello": "world",
			})
		})
	})
})
