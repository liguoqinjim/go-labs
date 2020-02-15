package foo_test

import . "github.com/onsi/ginkgo"
import "github.com/stretchr/testify/assert"

var _ = Describe("test001", func() {
	It("should testify to its correctness", func() {
		assert.Equal(GinkgoT(), "foo", "foo")
	})
})
