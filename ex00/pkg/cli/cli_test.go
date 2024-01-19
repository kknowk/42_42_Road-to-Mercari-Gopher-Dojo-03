// package cli
package cli

import "testing"

func TestParseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "no port", args: args{args: []string{"./omikuji"}}, want: "", wantErr: true},
		{name: "valid port", args: args{args: []string{"./omikuji", "8080"}}, want: "8080", wantErr: false},
		{name: "invalid port mix non integer", args: args{args: []string{"./omikuji", "80a0"}}, want: "", wantErr: true},
		{name: "invalid port over 65535", args: args{args: []string{"./omikuji", "80800"}}, want: "", wantErr: true},
		{name: "invalid port less 1024", args: args{args: []string{"./omikuji", "80"}}, want: "", wantErr: true},
		{name: "invalid port no integer", args: args{args: []string{"./omikuji", "abc"}}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		// テストケースのローカルコピーを作成
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ParseArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
