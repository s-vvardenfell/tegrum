package telegram

import (
	"path/filepath"
	"testing"
)

func TestTelegram_DownLoadFile(t *testing.T) {
	path := "W:/Golang/src/Backuper/resources"

	type args struct {
		fileId string
		dst    string
	}
	tests := []struct {
		name string
		tr   *Telegram
		args args
	}{
		{
			name: "base test",
			tr:   NewTelegram(filepath.Join(path, "telegram.json")),
			args: args{
				fileId: "BQACAgIAAxkDAAJYHGIzUnxyLIkpxsmfeyIywP7wdt_VAAKMGAACEzqZSfw6u4tyzo_pIwQ",
				dst:    "resources/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.DownLoadFile(tt.args.fileId, tt.args.dst)
		})
	}
}
