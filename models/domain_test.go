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
			domainOneSetting = DomainSetting{Interval: 10, Retry: 10, Timeout: 10}
			domainTwoSetting = DomainSetting{Interval: 60, Retry: 3, Timeout: 5}
			domainOne = Domain{Name: "Clowl 1", URL: "http://clowl.org", Setting: &domainOneSetting}
			domainTwo = Domain{Name: "Clowl 2", URL: "http://clowl.com", Setting: &domainTwoSetting}
		})

		Context("Domain struct", func() {
			It("Domain name", func() {
				Expect(domainOne.Name).To(Equal("Clowl 1"))
			})
			It("Domain URL", func() {
				Expect(domainTwo.URL).To(Equal("http://clowl.com"))
			})
			It("Domain Setting is OK", func() {
				Expect(domainOne.Setting).To(Equal(&domainOneSetting))
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
