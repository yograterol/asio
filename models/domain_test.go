package models_test

import (
	. "github.com/Clowl/asio/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain", func() {
	var (
		domainOne        Domain
		domainTwo        Domain
		domainOneSetting DomainSetting
		domainTwoSetting DomainSetting
	)

	Describe("Domain and Domain settings struct type", func() {
		BeforeEach(func() {
			domainOne = Domain{Name: "Clowl 1", URL: "http://clowl.org"}
			domainTwo = Domain{Name: "Clowl 2", URL: "http://clowl.com"}
			domainOneSetting = DomainSetting{Domain: &domainOne, Interval: 10, Retry: 10, Timeout: 10}
			domainTwoSetting = DomainSetting{Domain: &domainTwo, Interval: 60, Retry: 3, Timeout: 5}
		})

		Context("Domain struct", func() {
			It("Domain name", func() {
				Expect(domainOne.Name).To(Equal("Clowl 1"))
			})
			It("Domain URL", func() {
				Expect(domainTwo.URL).To(Equal("http://clowl.com"))
			})
		})

		Context("Domain Setting struct", func() {
			It("Domain Setting is valid", func() {
				Expect(domainTwoSetting.Validate()).To(BeTrue())
			})
			It("Domain Setting is not valid", func() {
				Expect(domainOneSetting.Validate()).ToNot(BeTrue())
			})
		})
	})
})
