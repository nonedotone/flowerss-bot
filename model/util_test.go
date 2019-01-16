package model

import "testing"

func Test_genHashID(t *testing.T) {
	type args struct {
		sLink string
		id    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"http://www.ruanyifeng.com/blog/atom.xml", "tag:www.ruanyifeng.com,2019:/blog//1.2054"}, "96b2e254"},
		{"2", args{"https://rsshub.app/guokr/scientific", "https://www.guokr.com/article/445877/"}, "770fff44"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genHashID(tt.args.sLink, tt.args.id); got != tt.want {
				t.Errorf("genHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}
