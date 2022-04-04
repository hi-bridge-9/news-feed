package feed

import (
	"reflect"
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
)

func TestGetNewInfo(t *testing.T) {
	type args struct {
		ts    *[]Tartget
		start *time.Time
		end   *time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *[]News
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNewInfo(tt.args.ts, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNewInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNewInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUpdate(t *testing.T) {
	type args struct {
		site  *Site
		start *time.Time
		end   *time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *News
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkUpdate(tt.args.site, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractNewArticle(t *testing.T) {
	type args struct {
		f     *gofeed.Feed
		start *time.Time
		end   *time.Time
	}

	start := time.Date(2022, 01, 01, 00, 00, 00, 00, time.UTC)
	end := time.Date(2022, 01, 02, 00, 00, 00, 00, time.UTC)

	beforeStart := time.Date(2021, 12, 31, 00, 00, 00, 00, time.UTC)
	afterEnd := time.Date(2022, 01, 03, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name string
		args args
		want *News
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &start,
					Items: []*gofeed.Item{
						{
							Title:           "article 1",
							Link:            "https://test.com/articles/test.html",
							UpdatedParsed:   &start,
							PublishedParsed: &start,
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: &News{
				SiteTitle: "site 1",
				SiteURL:   "https://test.com",
				Articles: []gofeed.Item{
					{
						Title:           "article 1",
						Link:            "https://test.com/articles/test.html",
						UpdatedParsed:   &start,
						PublishedParsed: &start,
					},
				},
			},
		},
		{
			name: "正常系(記事の投稿日なし)",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &start,
					Items: []*gofeed.Item{
						{
							Title:         "article 1",
							Link:          "https://test.com/articles/test.html",
							UpdatedParsed: &start,
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: &News{
				SiteTitle: "site 1",
				SiteURL:   "https://test.com",
				Articles: []gofeed.Item{
					{
						Title:           "article 1",
						Link:            "https://test.com/articles/test.html",
						UpdatedParsed:   &start,
						PublishedParsed: &start,
					},
				},
			},
		},
		{
			name: "正常系(開始範囲よりも古い記事のみ)",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &start,
					Items: []*gofeed.Item{
						{
							Title:           "article 1",
							Link:            "https://test.com/articles/test.html",
							UpdatedParsed:   &beforeStart,
							PublishedParsed: &beforeStart,
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: nil,
		},
		{
			name: "正常系(フィード用ファイルの更新が対象開始時刻よりも前)",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &beforeStart,
					Items: []*gofeed.Item{
						{
							Title:           "article 1",
							Link:            "https://test.com/articles/test.html",
							UpdatedParsed:   &beforeStart,
							PublishedParsed: &beforeStart,
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: nil,
		},

		{
			name: "正常系(フィード用ファイルの更新が対象終了時刻よりも後)",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &afterEnd,
					Items: []*gofeed.Item{
						{
							Title:           "article 1",
							Link:            "https://test.com/articles/test.html",
							UpdatedParsed:   &afterEnd,
							PublishedParsed: &afterEnd,
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: nil,
		},

		{
			name: "異常系(記事の投稿日＆更新日なし)",
			args: args{
				f: &gofeed.Feed{
					Title:         "site 1",
					Link:          "https://test.com",
					UpdatedParsed: &start,
					Items: []*gofeed.Item{
						{
							Title: "article 1",
							Link:  "https://test.com/articles/test.html",
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractNewArticle(tt.args.f, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nget %v\n want %v", got, tt.want)
			}
		})
	}
}
