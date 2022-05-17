package config

import (
	"reflect"
	"testing"
)

func TestLog_Empty(t *testing.T) {
	type fields struct {
		Dir      string
		FileName string
		empty    bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name:    "正常系_empty=trueの場合_trueとエラーを返す",
			fields:  fields{empty: true},
			want:    true,
			wantErr: true,
		},
		{
			name:    "正常系_empty=falseの場合_falseとnilを返す",
			fields:  fields{empty: false},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogArgs{
				Dir:      tt.fields.Dir,
				FileName: tt.fields.FileName,
				empty:    tt.fields.empty,
			}
			got, err := l.Empty()
			if (err != nil) != tt.wantErr {
				t.Errorf("Empty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Empty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyLog(t *testing.T) {
	tests := []struct {
		name string
		want *LogArgs
	}{
		{
			name: "正常系_emptyがtrueであること",
			want: &LogArgs{empty: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmptyLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmptyLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_ExistsDir(t *testing.T) {
	type fields struct {
		Dir      string
		FileName string
		empty    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "正常系_Dirの指定がある場合_trueを返す",
			fields: fields{
				Dir:      "log/purple/api",
				FileName: "application.log",
				empty:    false,
			},
			want: true,
		},
		{
			name: "正常系_Dirの指定が空の場合_falseを返す",
			fields: fields{
				Dir:      "",
				FileName: "application.log",
				empty:    false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogArgs{
				Dir:      tt.fields.Dir,
				FileName: tt.fields.FileName,
				empty:    tt.fields.empty,
			}
			if got := l.ExistsDir(); got != tt.want {
				t.Errorf("ExistsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_ExistsFile(t *testing.T) {
	type fields struct {
		Dir      string
		FileName string
		empty    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "正常系_FileNameの指定がある場合_trueを返す",
			fields: fields{
				Dir:      "log/purple/api",
				FileName: "application.log",
				empty:    false,
			},
			want: true,
		},
		{
			name: "正常系_FileNameの指定が空の場合_falseを返す",
			fields: fields{
				Dir:      "log/purple/api",
				FileName: "",
				empty:    false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogArgs{
				Dir:      tt.fields.Dir,
				FileName: tt.fields.FileName,
				empty:    tt.fields.empty,
			}
			if got := l.ExistsFile(); got != tt.want {
				t.Errorf("ExistsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
