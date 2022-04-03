package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"time"

	"github.com/hi-bridge-9/news-feed/feed"
)

func main() {
	// 実行時のコマンドライン引数から参照先ファイル名を入力（指定されない場合）
	path := flag.String("path", "data/input/feed_list.json", "For read target list")
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
	dir := "data/output/"
	fn := feed.MakeFileName(&start, &end)
	if err := feed.ExportFile(newsList, dir+fn); err != nil {
		panic(err)
	}
}
