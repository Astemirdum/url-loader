package loader

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_loadURL(t *testing.T) {
	t.Parallel()
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		args    args
		want    loadURLResult
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				u: "https://google.com",
			},
			want: loadURLResult{
				u:    "https://google.com",
				size: 10_000,
				ltv:  786641875,
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				u: "http:/dsmflsf.ru",
			},
			want:    loadURLResult{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := loadURL(context.Background(), tt.args.u, &http.Client{Timeout: time.Second * 2})
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want.u, got.u)
				require.GreaterOrEqual(t, got.size, tt.want.size)
			}
		})
	}
}
