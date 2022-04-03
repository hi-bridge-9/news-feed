package feed

import (
	"errors"
	"log"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

func GetNewInfo(ts *[]Tartget, start, end *time.Time) (*[]News, error) {
	// 取得対象の開始時刻が入力にない場合、エラーを返却
	if start == nil {
		return nil, errors.New("start of range is not exist")
	}

	// 終了時刻が入力にない場合、実行日のAM0:00(UTC)を終了時刻とみなす
	if end == nil {
		now := time.Now()
		time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 00, time.UTC)
		end = &now
	}

	newsList := &[]News{}
	for _, t := range *ts {
		for _, site := range t.Sites {
			news, err := checkUpdate(&site, start, end)
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

func checkUpdate(site *Site, start, end *time.Time) (*News, error) {
	// URL先からフィード用コンテンツを取得
	feed, err := gofeed.NewParser().ParseURL(site.FeedURL)
	if err != nil {
		log.Printf("Name: %s", site.Name)
		log.Printf("URL : %s", site.FeedURL)
		return nil, err
	}

	// 新しい情報のみを抽出
	// もし新しい情報がなくても、エラーは出さずに初期値を返す
	return extractNew(feed, start, end), nil
}

func extractNew(f *gofeed.Feed, start, end *time.Time) *News {
	// フィード用ファイルの更新がされていなければ、記事の確認処理を行わない
	if f.UpdatedParsed.Unix() < start.Unix() && f.UpdatedParsed.Unix() >= end.Unix() {
		log.Printf("%s: feed file is not updated\n", f.Title)
		log.Printf("update date -> %s\n", f.UpdatedParsed)
		return nil
	}

	// 投稿日が新しい順に記事をソート
	var articles ByPublishedParsed = f.Items
	sort.Sort(articles)

	// 投稿日が取得対象の範囲内なら抽出
	news := &News{}
	for _, article := range articles {
		// 投稿日時と更新日時の両方が存在しない場合はスキップ
		// 更新日時のみが存在する場合、更新日時を投稿日時として扱う
		if article.UpdatedParsed == nil && article.PublishedParsed == nil {
			log.Printf("%s: not exist published date and updated date\n", f.Title)
			return nil
		} else if article.PublishedParsed == nil {
			article.PublishedParsed = article.UpdatedParsed
		}

		// 記事の投稿時刻が取得対象範囲内であれば抽出
		if article.PublishedParsed.Unix() >= start.Unix() {
			if article.PublishedParsed.Unix() < end.Unix() {
				news.Articles = append(news.Articles, *article)
			}
		} else {
			// 取得範囲の開始時刻よりも古い記事が出てきたら、抽出処理を終了
			break
		}
	}

	if len(news.Articles) == 0 {
		log.Printf("%s: article is not updated\n", f.Title)
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
