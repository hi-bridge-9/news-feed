package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hi-bridge-9/news-feed/feed"
)

func main() {
	// 実行時のコマンドライン引数から参照先ファイル名を入力（指定されない場合）
	path := flag.String("path", "input_data/feed_list.json", "For read target list")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	// ファイルの内容を読み取り、構造体にマッピング
	var targetList []feed.Tartget
	if err := json.Unmarshal(bytes, &targetList); err != nil {
		panic(err)
	}

	// 取得したい範囲の開始時刻、終了時刻を設定
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day()-20, 00, 00, 00, 00, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 00, time.UTC)

	// 情報取得対象の情報、取得対象範囲をもとに、新しい情報を取得
	newsList, err := feed.GetNewInfo(&targetList, &start, &end)
	if err != nil {
		panic(err)
	}

	// データ活用例： 新しく投稿された情報をマークダウン形式のログファイルとして出力
	if err := exportOutpuFile(newsList, &start, &end); err != nil {
		panic(err)
	}
}


func exportOutpuFile(newsList *[]feed.News, start, end *time.Time) error {
	fn := makeFileName(start, end)
	dir := "export_log/"
	msg := convertToMessage(newsList)
	return ioutil.WriteFile(dir+fn, []byte(msg), 0664)

}

func makeFileName(start, end *time.Time) string {
	dateRange := start.Format("20060102150405")
	dateRange += "-"
	dateRange += end.Format("20060102150405")
	return fmt.Sprintf("%s.md", dateRange)
}

func convertToMessage(newsList *[]feed.News) (msg string) {
	if newsList == nil {
		return "今日の更新情報はありません"
	}

	msg += fmt.Sprintf("# %s\n", "更新情報")
	for _, news := range *newsList {
		msg += fmt.Sprintf("## **%s**\n", news.SiteTitle)
		for i, article := range news.Articles {
			msg += fmt.Sprintf("### %d. %s\n", i+1, article.Title)
			msg += fmt.Sprintf("   - 時刻: %s\n", article.PublishedParsed)
			msg += fmt.Sprintf("   - URL : %s\n \n", article.Link)
		}
	}

	return
}
