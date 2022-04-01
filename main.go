package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hi-bridge-9/news-feed/feed"
	"github.com/hi-bridge-9/news-feed/target"
)

func main() {
	// 実行時のコマンドライン引数から参照先ファイル名を入力（指定されない場合）
	path := flag.String("path", ".example_data/feed_list.json", "for read feed target list")
	bytes, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	// ファイルの内容を読み取り、構造体にマッピング
	var InfoList []target.Info
	if err := json.Unmarshal(bytes, &InfoList); err != nil {
		panic(err)
	}

	// 取得したい範囲の開始時刻、終了時刻を設定
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day()-1, 00, 00, 00, 00, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 00, time.UTC)

	// 情報取得対象の情報、取得対象範囲をもとに、新しい情報を取得
	newsList, err := feed.GetNewInfo(&InfoList, &start, &end)
	if err != nil {
		panic(err)
	}

	// 通知用に適した形式で文章を作成
	msg := makeMessage(newsList)

	fmt.Println(msg)
}

func makeMessage(newsList *[]feed.News) string {
	var msg string
	if len(*newsList) == 0 {
		return "今日の更新情報はありません"
	}

	for _, out := range *newsList {
		msg += fmt.Sprintf("【%s】\n", out.SiteTitle)
		for i, article := range out.Articles {
			msg += fmt.Sprintf("%d. %s\n", i, article.Title)
			msg += fmt.Sprintf("URL: %s\n", article.Link)
			msg += "-----------------------------------------\n"
		}
	}

	return msg
}
