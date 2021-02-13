package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "todoApp/server/controllers"
)

var _ = Describe("TodoController", func() {
	var (
		one int
		two int
	)

	BeforeEach(func() {
		one = 1
		two = 2
	})

	Describe("Test1", func() {
		Context("context 1", func() {
			It("should be one", func() {
				Expect(two - one).To(Equal(one))
			})
		})
	})

	Describe("Test2", func() {
		Context("context 2", func() {
			It("should be tree", func() {
				Expect(two + one).To(Equal(3))
			})
		})
	})
})
