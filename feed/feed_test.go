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

	start := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)
	end := time.Date(2999, 01, 02, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name    string
		args    args
		want    *[]News
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "正常系",
		// 	args: args{
		// 		ts: &[]Tartget{
		// 			{
		// 				Sites: []Site{
		// 					{
		// 						Name:    "site 1",
		// 						TopURL:  "https://developers-jp.googleblog.com/",
		// 						FeedURL: "https://developers-jp.googleblog.com/atom.xml",
		// 					},
		// 				},
		// 			},
		// 		},
		// 		start: &min,
		// 		end:   &now,
		// 	},
		// 	want:    &[]News{},
		// 	wantErr: false,
		// },
		{
			name: "正常系(更新記事なし)",
			args: args{
				ts: &[]Tartget{
					{
						Sites: []Site{
							{
								Name:    "site 1",
								TopURL:  "https://developers-jp.googleblog.com/",
								FeedURL: "https://developers-jp.googleblog.com/atom.xml",
							},
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want:    &[]News{},
			wantErr: false,
		},
		{
			name: "正常系(終了時刻なし)",
			args: args{
				ts: &[]Tartget{
					{
						Sites: []Site{
							{
								Name:    "site 1",
								TopURL:  "https://developers-jp.googleblog.com/",
								FeedURL: "https://developers-jp.googleblog.com/atom.xml",
							},
						},
					},
				},
				start: &start,
			},
			want:    &[]News{},
			wantErr: false,
		},
		{
			name: "異常系(開始時刻なし)",
			args: args{
				ts: &[]Tartget{
					{
						Sites: []Site{
							{
								Name:    "site 1",
								TopURL:  "https://developers-jp.googleblog.com/",
								FeedURL: "https://developers-jp.googleblog.com/atom.xml",
							},
						},
					},
				},
				end: &end,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系(フィード用URLが不正)",
			args: args{
				ts: &[]Tartget{
					{
						Sites: []Site{
							{
								Name:    "site 1",
								TopURL:  "https://developers-jp.googleblog.com/",
								FeedURL: "https://",
							},
						},
					},
				},
				start: &start,
				end:   &end,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNewInfo(tt.args.ts, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("getErr %v\n wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get %v\n want %v", got, tt.want)
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

	// 記事を取得しない日付に設定（記事を取得できるかは本質ではない）
	start := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)
	end := time.Date(2999, 01, 02, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name    string
		args    args
		want    *News
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			args: args{
				site: &Site{
					Name:    "site 1",
					TopURL:  "https://developers-jp.googleblog.com/",
					FeedURL: "https://developers-jp.googleblog.com/atom.xml",
				},
				start: &start,
				end:   &end,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "異常系(フィード用URLが不正)",
			args: args{
				site: &Site{
					Name:    "site 1",
					TopURL:  "https://developers-jp.googleblog.com/",
					FeedURL: "https://",
				},
				start: &start,
				end:   &end,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkUpdate(tt.args.site, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("getErr %v\n wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get %v\n want %v", got, tt.want)
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
				t.Errorf("get %v\n want %v", got, tt.want)
			}
		})
	}
}

func TestByPublishedParsed_Len(t *testing.T) {
	tests := []struct {
		name string
		p    ByPublishedParsed
		want int
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			p:    ByPublishedParsed{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("get %v\n want %v", got, tt.want)
			}
		})
	}
}

func TestByPublishedParsed_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}

	pub1 := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)
	pub2 := time.Date(2999, 01, 02, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name string
		p    ByPublishedParsed
		args args
	}{
		// TODO: Add test cases.
		{
			name: "正常系",
			p: ByPublishedParsed{
				&gofeed.Item{
					Title:           "site 1",
					Link:            "https://test.com",
					PublishedParsed: &pub1,
				},
				&gofeed.Item{
					Title:           "site 2",
					Link:            "https://test2.com",
					PublishedParsed: &pub2,
				},
			},
			args: args{
				i: 0,
				j: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByPublishedParsed_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}

	pub1 := time.Date(2999, 01, 01, 00, 00, 00, 00, time.UTC)
	pub2 := time.Date(2999, 01, 02, 00, 00, 00, 00, time.UTC)

	tests := []struct {
		name string
		p    ByPublishedParsed
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系(前者の方が投稿日が新しい)",
			p: ByPublishedParsed{
				&gofeed.Item{
					Title:           "site 1",
					Link:            "https://test.com",
					PublishedParsed: &pub1,
				},
				&gofeed.Item{
					Title:           "site 2",
					Link:            "https://test2.com",
					PublishedParsed: &pub2,
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: false,
		},
		{
			name: "正常系(後者の方が投稿日が新しい)",
			p: ByPublishedParsed{
				&gofeed.Item{
					Title:           "site 1",
					Link:            "https://test.com",
					PublishedParsed: &pub2,
				},
				&gofeed.Item{
					Title:           "site 2",
					Link:            "https://test2.com",
					PublishedParsed: &pub1,
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("get %v\n want %v", got, tt.want)
			}
		})
	}
}
