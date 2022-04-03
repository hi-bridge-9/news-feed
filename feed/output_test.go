package feed

import (
	"testing"
	"time"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExportFile(tt.args.newsList, tt.args.fp); (err != nil) != tt.wantErr {
				t.Errorf("ExportFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMakeFileName(t *testing.T) {
	type args struct {
		start *time.Time
		end   *time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFileName(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MakeFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToMessage(t *testing.T) {
	type args struct {
		newsList *[]News
	}
	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := convertToMessage(tt.args.newsList); gotMsg != tt.wantMsg {
				t.Errorf("convertToMessage() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
