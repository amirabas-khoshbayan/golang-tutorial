package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"os"
)

var redditSession *geddit.LoginSession

type Story struct {
	title string
	url   string
}

func init() {
	var err error
	redditSession, err = geddit.NewLoginSession("g_d_bot", "K417k4FTua52", "gdAgent v0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	inPutProRedditCh := make(chan Story, 10)
	inPutITRedditCh := make(chan Story, 10)

	outConsoleCh := make(chan Story, 10)
	outFileCh := make(chan Story, 10)

	go loadReddit("programming", inPutProRedditCh)
	go loadReddit("IT ", inPutITRedditCh)

	file, err := os.Create("story.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	go outConsole(outConsoleCh)
	go outFile(outFileCh, file)

	proRedditOpen := true
	itRedditOpen := true

	for proRedditOpen || itRedditOpen {
		select {
		case proStor, open := <-inPutProRedditCh:
			if open {
				outConsoleCh <- proStor
				outFileCh <- proStor
			} else {
				proRedditOpen = false
			}
		case ITStor, open := <-inPutITRedditCh:
			if open {
				outConsoleCh <- ITStor
				outFileCh <- ITStor
			} else {
				itRedditOpen = false
			}
		}
	}

	//var stories []Story
	//
	//proStory := loadReddit("programming")
	//ITStory := loadReddit("it")
	//if proStory != nil {
	//	for _, proStor := range proStory {
	//		stories = append(stories, proStor)
	//	}
	//}
	//
	//if ITStory != nil {
	//	for _, ITStor := range ITStory {
	//		stories = append(stories, ITStor)
	//	}
	//}
	//
	//file, err := os.Create("story.txt")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//for _, s := range stories {
	//	fmt.Fprintf(file, "%s \n \t %s \n\n", s.title, s.url)
	//}
	//for _, s := range stories {
	//	fmt.Printf("%s \n \t %s \n\n", s.title, s.url)
	//}
	//
	//fmt.Printf("%v \n\n\n", stories)

}

func loadReddit(filter string, storiesCh chan<- Story) {
	defer close(storiesCh)

	var listOption geddit.ListingOptions
	submission, err := redditSession.SubredditSubmissions(filter, "", listOption)
	if err != nil {
		fmt.Println(err)
	}

	for _, s := range submission {
		newStory := Story{
			title: s.Title,
			url:   s.URL,
		}
		storiesCh <- newStory
	}

}

func outConsole(storyCh <-chan Story) {
	for {
		s := <-storyCh
		fmt.Printf("%s \n \t %s \n\n", s.title, s.url)
	}
}

func outFile(storyCh <-chan Story, file *os.File) {
	for {
		s := <-storyCh
		fmt.Fprintf(file, "%s \n \t %s \n\n", s.title, s.url)
	}
}
