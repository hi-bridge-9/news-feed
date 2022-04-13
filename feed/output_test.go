package feed

import (
	"fmt"
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
)

func TestExportFile(t *testing.T) {
	type args struct {
		newsList *[]News
		fp       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			args: args{
				newsList: &[]News{},
				fp:       "../data/test/test.txt",
			},
			wantErr: false,
		},
		{
			name: "異常系(ファイルパスなし)",
			args: args{
				newsList: &[]News{},
				fp:       "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExportFile(tt.args.newsList, tt.args.fp); (err != nil) != tt.wantErr {
				t.Errorf("getErr %v\n wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMakeFileName(t *testing.T) {
	type args struct {
		start *time.Time
		end   *time.Time
	}

	start := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)
	end := time.Date(2999, 01, 02, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			args: args{
				start: &start,
				end:   &end,
			},
			want:    "2999-01-01_2999-01-02.md",
			wantErr: false,
		},
		{
			name: "異常系(開始時刻なし)",
			args: args{
				end: &end,
			},
			want:    "",
			wantErr: true,
		},

		{
			name: "異常系(終了時刻なし)",
			args: args{
				start: &start,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeFileName(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("getErr %v\n wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("get %v\n want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToMessage(t *testing.T) {
	type args struct {
		newsList *[]News
	}

	pub := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			args: args{
				newsList: &[]News{
					{
						SiteTitle: "site 1",
						SiteURL:   "https://test.com",
						Articles: []gofeed.Item{
							{
								Title:           "article 1",
								PublishedParsed: &pub,
								Link:            "https://test.com/articles/test.xml",
							},
						},
					},
				},
			},
			wantMsg: fmt.Sprintln("# 更新情報\n## **site 1**\n### 1. article 1\n- 時刻: 2999-01-01 00:00:00 +0000 UTC\n- URL : https://test.com/articles/test.xml\n "),
		},
		{
			name: "正常系(新しい情報なし)",
			args: args{
				newsList: nil,
			},
			wantMsg: "更新情報はありません",
		},
		{
			name: "正常系(リクエスト時のエラーメッセージあり)",
			args: args{
				newsList: &[]News{
					{
						SiteTitle:  "site 1",
						SiteURL:    "https://test.com/articles/test.xml",
						ErrMessage: "http error: 404 Not Found",
					},
				},
			},
			wantMsg: fmt.Sprintln("# 更新情報\n## **site 1**\n- **Error: http error: 404 Not Found**\n- URL  : https://test.com/articles/test.xml\n "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := convertToMessage(tt.args.newsList); gotMsg != tt.wantMsg {
				t.Errorf("get %v\n want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
