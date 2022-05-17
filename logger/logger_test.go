package logger

import (
	"testing"

	"github.com/haradayoshitsugucz/purple-server/config"
)

func Test_getFilePath(t *testing.T) {
	type args struct {
		setting  *config.LoggerSetting
		logModel *config.LogArgs
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "正常系_起動引数の指定の場合",
			args:    args{
				setting:  &config.LoggerSetting{},
				logModel: &config.LogArgs{
					Dir:      "../log/purple/tmp",
					FileName: "20220517.log",
				},
			},
			want:    "../log/purple/tmp/20220517.log",
			wantErr: false,
		},
		{
			name:    "正常系_起動引数でファイルのみを指定した場合",
			args:    args{
				setting:  &config.LoggerSetting{},
				logModel: &config.LogArgs{
					Dir:      "",
					FileName: "20220517.log",
				},
			},
			want:    "20220517.log",
			wantErr: false,
		},
		{
			name:    "異常系_起動引数でディレクトリのみを指定した場合",
			args:    args{
				setting:  &config.LoggerSetting{},
				logModel: &config.LogArgs{
					Dir:      "../log/purple/tmp",
					FileName: "",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "正常系_設定ファイルの場合",
			args:    args{
				setting:  &config.LoggerSetting{
					LogDir:        "../log/purple/api",
					FileName:      "application.log",
				},
				logModel: config.EmptyLog(),
			},
			want:    "../log/purple/api/application.log",
			wantErr: false,
		},
		{
			name:    "正常系_設定ファイルでファイルのみを指定した場合",
			args:    args{
				setting:  &config.LoggerSetting{
					LogDir:        "",
					FileName:      "application.log",
				},
				logModel: config.EmptyLog(),
			},
			want:    "application.log",
			wantErr: false,
		},
		{
			name:    "異常系_設定ファイルでディレクトリのみを指定した場合",
			args:    args{
				setting:  &config.LoggerSetting{
					LogDir:        "../log/purple/api",
					FileName:      "",
				},
				logModel: config.EmptyLog(),
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "正常系_設定ファイルと起動引数の指定が両方存在する場合、起動引数の指定が取得できること",
			args:    args{
				setting:  &config.LoggerSetting{
					LogDir:        "../log/purple/api",
					FileName:      "application.log",
				},
				logModel: &config.LogArgs{
					Dir:      "../log/purple/tmp",
					FileName: "20220517.log",
				},
			},
			want:    "../log/purple/tmp/20220517.log",
			wantErr: false,
		},
		{
			name:    "正常系_設定ファイルと起動引数の指定がどちらも存在しない場合",
			args:    args{
				setting:  &config.LoggerSetting{},
				logModel: config.EmptyLog(),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFilePath(tt.args.setting, tt.args.logModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFilePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
