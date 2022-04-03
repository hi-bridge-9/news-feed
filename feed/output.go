package feed

import (
	"fmt"
	"io/ioutil"
	"time"
)

func ExportFile(newsList *[]News, fp string) error {
	msg := convertToMessage(newsList)
	return ioutil.WriteFile(fp, []byte(msg), 0664)

}

func MakeFileName(start, end *time.Time) string {
	dateRange := start.Format("20060102150405")
	dateRange += "-"
	dateRange += end.Format("20060102150405")
	return fmt.Sprintf("%s.md", dateRange)
}

func convertToMessage(newsList *[]News) (msg string) {
	if newsList == nil {
		return "更新情報はありません"
	}

	msg += fmt.Sprintf("# %s\n", "更新情報")
	for _, news := range *newsList {
		msg += fmt.Sprintf("## **%s**\n", news.SiteTitle)
		for i, article := range news.Articles {
			msg += fmt.Sprintf("### %d. %s\n", i+1, article.Title)
			msg += fmt.Sprintf("- 時刻: %s\n", article.PublishedParsed)
			msg += fmt.Sprintf("- URL : %s\n \n", article.Link)
		}
	}

	return
}
