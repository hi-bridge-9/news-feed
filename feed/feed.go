package feed

import (
	"errors"
	"log"
	"sort"
	"time"

	"github.com/hi-bridge-9/news-feed/target"
	"github.com/mmcdole/gofeed"
)

func GetNewInfo(urls *[]target.Info, start, end *time.Time) (*[]News, error) {
	// 取得対象の開始時刻が入力にない場合、エラーを返却
	if start == nil {
		return nil, errors.New("start of range is not exist")
	}

	// 終了時刻が入力にない場合、現在時刻を終了時刻とみなす
	if end == nil {
		now := time.Now().UTC()
		end = &now
	}

	var newsList *[]News
	for _, url := range *urls {
		for _, site := range url.Sites {
			news, err := checkUpdate(site, start, end)
			if err != nil {
				return nil, err
			}
			if news != nil {
				*newsList = append(*newsList, *news)
			}
		}
	}

	return newsList, nil
}

func checkUpdate(site target.Site, start, end *time.Time) (*News, error) {
	// URL先からフィード用コンテンツを取得
	feed, err := gofeed.NewParser().ParseURL(site.FeedURL)
	if err != nil {
		return nil, err
	}

	// 新しい情報のみを抽出
	// もし新しい情報がなくても、エラーは出さずに初期値を返す
	return extract(feed, start, end), nil
}

func extract(f *gofeed.Feed, start, end *time.Time) *News {
	// フィード用ファイルの更新がされていなければ、記事の確認処理を行わない
	if f.UpdatedParsed.Unix() < start.Unix() || f.UpdatedParsed.Unix() < end.Unix() {
		log.Println("feed file is not updated")
		return nil
	}

	// 記事の投稿日の新しい順にソート
	var items ByPublishedParsed = f.Items
	sort.Sort(items)

	// 投稿日が取得対象の範囲内なら抽出
	var news *News
	for _, article := range f.Items {
		if article.PublishedParsed.Unix() >= start.Unix() {
			if article.PublishedParsed.Unix() < end.Unix() {
				news.Articles = append(news.Articles, *article)
			}
		} else {
			// 取得開始時刻よりも記事が古い場合、抽出処理を終了
			break
		}
	}

	if len(news.Articles) == 0 {
		log.Println("article is not updated")
		return nil
	}

	// 記事が抽出されている場合のみ、サイト名やURLを取得する
	news.SiteTitle = f.Title
	news.SiteURL = f.Link

	return news
}


func (p ByPublishedParsed) Len() int {
	return len(p)
}

func (p ByPublishedParsed) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByPublishedParsed) Less(i, j int) bool {
	return p[j].PublishedParsed.Unix() < p[i].PublishedParsed.Unix()
}
