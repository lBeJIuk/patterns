package main
import (
	"./news"
	"./chanels"
	"./iobserver"
	"fmt"
)

func main() {
	n := news.NewsSource{make(map[string]iobserver.IObserver), []news.News{}}
	c1 := chanels.Chanel1{"Chanel1"}
	c2 := chanels.Chanel2{"Chanel2"}
	c3 := chanels.Chanel3{"Chanel3"}

	n.RegisterObserver("observer1", c1)
	n.RegisterObserver("observer2", c2)

	fmt.Println("-------New news")
	n.AddNews(news.News{"News1", "Text1"})

	n.RegisterObserver("observer3", c3)

	fmt.Println("-------New news")
	n.AddNews(news.News{"News2", "Text2"})

	n.RemoveObserver("observer1")
	n.RemoveObserver("observer2")

	fmt.Println("-------New news")
	n.AddNews(news.News{"News3", "Text3"})

}