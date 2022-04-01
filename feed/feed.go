package feed

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"time"

	"github.com/hi-bridge-9/news-feed/target"
)

func GetNewInfo(urls *[]target.Info, start, end *time.Time) (*[]Output, error) {
	// 取得対象の開始時刻が入力にない場合、エラーを返却
	if start == nil {
		return nil, errors.New("start of range is not exist")
	}

	// 終了時刻が入力にない場合、現在時刻を終了時刻とみなす
	if end == nil {
		now := time.Now().UTC()
		end = &now
	}

	var outList *[]Output
	for _, url := range *urls {
		for _, site := range url.Sites {
			out, err := findUpdate(site, start, end)
			if err != nil {
				return nil, err
			}
			if out != nil {
				*outList = append(*outList, *out)
			}
		}
	}

	return outList, nil
}

func findUpdate(site target.Site, start, end *time.Time) (*Output, error) {
	// URL先からフィード用コンテンツを取得
	bytes, err := fetch(site.FeedURL)
	if err != nil {
		return nil, fmt.Errorf("error in fetch feed contents: %w", err)
	}

	// コンテンツをパース
	feed, err := parse(bytes, filepath.Ext(site.FeedURL))
	if err != nil {
		return nil, fmt.Errorf("error in parse feed contents: %w", err)
	}

	// 新しい情報のみを抽出
	// もし新しい情報がなくても、エラーは出さずに初期値を返す
	return extract(feed, start, end), nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// リクエストボディを読み取り、返却
	return ioutil.ReadAll(resp.Body)
}

func parse(data []byte, ext string) (*Feed, error) {
	var feed *Feed
	var err error

	// フィードの形式がatomの時は専用のパースを実施
	switch ext {
	case ".atom":
		feed, err = atom.Parse(data)
	default:
		feed, err = rss.Parse(data)
	}

	if err != nil {
		return nil, err
	}

	return feed, nil
}

func extract(f *Feed, start, end *time.Time) *Output {
	// フィード用ファイルの更新がされていなければ、記事の確認処理を行わない
	if f.UpdatedParsed.Unix() < start.Unix() || f.UpdatedParsed.Unix() < end.Unix() {
		log.Println("feed file is not updated")
		return nil
	}

	// 記事の投稿日の新しい順にソート
	var items ByPublishedParsed = f.Items
	sort.Sort(items)

	// 投稿日が取得対象の範囲内なら抽出
	var out *Output
	for _, article := range f.Items {
		if article.PublishedParsed.Unix() >= start.Unix() {
			if article.PublishedParsed.Unix() < end.Unix() {
				out.Articles = append(out.Articles, article)
			}
		} else {
			// 取得開始時刻よりも記事が古い場合、抽出を終了（新しい順にソートしているため、これ以降抽出することはない）
			break
		}
	}

	if len(out.Articles) == 0 {
		log.Println("article is not updated")
		return nil
	}

	// 記事が抽出されている場合のみ、サイト名やURLを取得する
	out.SiteTitle = f.Title
	out.SiteURL = f.Link

	return out
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
