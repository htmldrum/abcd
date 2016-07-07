package main_test

import (
	"os"
	. "github.com/htmldrum/abcd/cmd/abcd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/htmldrum/abcd/fs"
	"github.com/spf13/afero"
	"path/filepath"
	"gopkg.in/h2non/gock.v1"
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
				feed = Feed{1, f_name, f_description, f_URI, f_subjects, f_networks, f_last_contact_datetime, articles}
			})
			It("Builds a Feed", func(){
				Expect(feed.Name).Should(Not(BeZero()))
			})
		})
	})
	Describe("abcd", func() {})
	Describe("SaveFeeds", func(){
	})
	Describe("RefreshFeeds", func(){
		It("Veifies feeds from template", func(){
			defer gock.Off()

			gock.New("http://www.abc.net.au").
				Get("/services/rss/programs.htm").
				Reply(200).
				File("programs.htm")

			feeds := RefreshFeeds()
			Expect(len(feeds)).To(BeNumerically("==", 201))
		})
	})
	Describe("read_config", func() {
		It("Should use a unixey path to store its config", func(){
			stub_env := []string{"HOME=/west/wing", "MOUNT=/media"}
			Expect("/west/wing/.abcd").To(Equal(GetConfDir(&stub_env)))
		})
		Describe("EnsureConfFile", func(){
			var mockFS afero.Fs
			test_home :="/home/darmock"
			test_path := filepath.Join(test_home, DIRNAME, CONFNAME)
			test_dir := filepath.Join(test_home, DIRNAME)

			Context("When conffile doesn't exist", func(){
				BeforeEach(func(){
					mockFS = fs.NewMockFs()
					mockFS.Mkdir(test_dir, os.ModeDir)
				})
				It("It creates the required conf dir", func(){
					fi, _ := mockFS.Stat(test_path)
					Expect(fi).To(BeNil())
					EnsureConfFile(test_path, mockFS)
					new_fi, _ := mockFS.Stat(test_path)
					Expect(new_fi).NotTo(BeNil())
				})
			})
			Context("When confdir exists", func(){
				BeforeEach(func(){
					mockFS = fs.NewMockFs()
					mockFS.Mkdir(test_path, os.ModeDir)
				})
				It("", func(){
					fi, _ := mockFS.Stat(test_path)
					mod_time := fi.ModTime()
					EnsureConfDir(test_path, mockFS)
					new_fi, _ := mockFS.Stat(test_path)
					Expect(new_fi.ModTime().Equal(mod_time)).To(Equal(true))
				})
			})
		})
		Describe("EnsureConfDir", func(){
			var mockFS afero.Fs
			test_path := "/home/darmock/.and_jalad"

			Context("When confdir doesn't exist", func(){
				BeforeEach(func(){
					mockFS = fs.NewMockFs()
				})
				It("It creates the required conf dir", func(){
					fi, _ := mockFS.Stat(test_path)
					Expect(fi).To(BeNil())
					EnsureConfDir(test_path, mockFS)
					new_fi, _ := mockFS.Stat(test_path)
					Expect(new_fi).NotTo(BeNil())
				})
			})
			Context("When confdir exists", func(){
				BeforeEach(func(){
					mockFS = fs.NewMockFs()
					mockFS.Mkdir(test_path, os.ModeDir)
				})
				It("Doesn't create the required conf dir", func(){
					fi, _ := mockFS.Stat(test_path)
					mod_time := fi.ModTime()
					EnsureConfDir(test_path, mockFS)
					new_fi, _ := mockFS.Stat(test_path)
					Expect(new_fi.ModTime().Equal(mod_time)).To(Equal(true))
				})
			})
		})
	})
})
