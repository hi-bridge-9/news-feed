package feed

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

func ExportFile(newsList *[]News, fp string) error {
	msg := convertToMessage(newsList)
	return ioutil.WriteFile(fp, []byte(msg), 0664)
}

func MakeFileName(start, end *time.Time) (string, error) {
	if start == nil || end == nil {
		return "", errors.New("start or end date is not exits\nstart: %v\nend: %v")
	}
	dateRange := start.Format("2006-01-02")
	dateRange += "_"
	dateRange += end.Format("2006-01-02")
	return fmt.Sprintf("%s.md", dateRange), nil
}

func convertToMessage(newsList *[]News) (msg string) {
	if newsList == nil {
		return "更新情報はありません"
	}

	msg += fmt.Sprintln("# 更新情報")
	for _, news := range *newsList {
		msg += fmt.Sprintf("## **%s**\n", news.SiteTitle)
		if news.errMessage != "" {
			msg += fmt.Sprintf("- **Error: %s**\n", news.errMessage)
			msg += fmt.Sprintf("- URL  : %v\n \n", news.SiteURL)
		} else {
			for i, article := range news.Articles {
				msg += fmt.Sprintf("### %d. %s\n", i+1, article.Title)
				msg += fmt.Sprintf("- 時刻: %s\n", article.PublishedParsed)
				msg += fmt.Sprintf("- URL : %s\n \n", article.Link)
			}
		}
	}
	return
}
