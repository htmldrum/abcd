package main_test

import (
	. "github.com/htmldrum/abcd/cmd/abcd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Abcd", func() {
	Describe("ustructs", func(){
		var article Article
		var feed Feed

		Context("Article", func(){
			BeforeEach(func(){
				a_name := "foo"
				a_url := "http://www.abc.net.au/news/2016-07-04/could-silent-fireworks-be-used-for-territory-day-celebrations/7567510"
				article = Article{a_name, a_url}
			})
			It("Builds an Article", func(){
				Expect(article.Name).Should(Not(BeZero()))
			})
		})
		Context("Feed", func(){
			BeforeEach(func(){
				f_name := "foo"
				f_description := "Debate, ideas and attitude"
				f_URI := "http://abc.net.au/bestof/bestofabc.xml"
				f_subjects := []string{"education", "health"}
				f_networks := []string{"Radio Australia", "Sport"}
				f_last_contact_datetime := "derp"
				names := []string{
					"Could silent fireworks bring an end to 'war zone' connotations on Territory Day?",
					"Sydney siege inquest: Police radios cut out as command given to storm cafe",
				}
				urls := []string{
					"http://www.abc.net.au/news/2016-07-04/could-silent-fireworks-be-used-for-territory-day-celebrations/7567510",
					"http://www.abc.net.au/news/2016-07-04/sydney-siege-inquest-police-radios-failed-as-cafe-stormed/7567740",
				}
				articles := []Article{}
				for i, _ := range names {
					articles = append(articles, Article{names[i], urls[i]})
				}
				feed = Feed{f_name, f_description, f_URI, f_subjects, f_networks, f_last_contact_datetime, articles}
			})
			It("Builds a Feed", func(){
				Expect(feed.Name).Should(Not(BeZero()))
			})
		})
	})
	Describe("abcd", func() {
	})
	Describe("read_config", func() {
		It("Should have a ReadConfig function", func(){
			ReadConfig()
			Expect(true).To(Equal(true))
		})
		It("Should use a unixey path to store its config", func(){
			stub_env := []string{"HOME=/west/wing", "MOUNT=/media"}
			Expect("/west/wing/.abcd").To(Equal(GetConfDir(&stub_env)))
		})
		Describe("EnsureConfDir", func(){
			Context("When confdir doesn't exist", func(){
			})
			Context("When confdir exists", func(){
				It("Should create the confidr", func(){
					// Assertions
				})
			})
		})
	})
})
