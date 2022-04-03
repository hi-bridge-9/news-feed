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
	tests := []struct {
		name string
		args args
		want *News
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractNewArticle(tt.args.f, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractNewArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
