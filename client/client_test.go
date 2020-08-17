package client

import (
	"github.com/fanjindong/dcs/utils"
	"reflect"
	"testing"
)

func TestTcpClient_Command(t *testing.T) {
	tests := []struct {
		name    string
		op      utils.Operation
		kv      []string
		want    string
		wantErr bool
	}{
		{name: "Set", op: utils.Set, kv: []string{"say", "hello world!!!"}, wantErr: false, want: "ok"},
		{name: "Get", op: utils.Get, kv: []string{"say"}, wantErr: false, want: "hello world!!!"},
		{name: "Del", op: utils.Del, kv: []string{"say"}, wantErr: false, want: "ok"},
		{name: "Get", op: utils.Get, kv: []string{"say"}, wantErr: true, want: ""},
	}
	for _, tt := range tests {
		c := NewTcpClient()
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Command(tt.op, tt.kv...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Command() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTcpClient_Command(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewTcpClient()
		_, _ = c.Command(utils.Set, "k", "v")
	}
}
