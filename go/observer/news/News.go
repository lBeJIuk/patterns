package news

import (
	"../iobserver"
)

type NewsSource struct {
	Observers map[string]iobserver.IObserver
	News []News
}

type News struct{
Name string
Text string
}
func (news NewsSource) RegisterObserver(name string, observer iobserver.IObserver) {
	news.Observers[name] = observer
}
func (news NewsSource) RemoveObserver(name string) {
	delete(news.Observers, name)
}

func (news NewsSource) NotifyObservers() {
	for k := range news.Observers {
		news.Observers[k].Update();
	}
}

func (newsSource NewsSource) AddNews(news News){
	newsSource.News = append(newsSource.News, news)
	newsSource.NotifyObservers()
}
