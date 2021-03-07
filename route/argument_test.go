package route

import (
	"reflect"
	"testing"
)

func Test_stringType_Parse(t *testing.T) {

	x := "hello world"
	xq := `"hello world"`
	xc := `"hello world`

	type args struct {
		content *string
	}
	tests := []struct {
		name    string
		s       stringType
		args    args
		want    interface{}
		wantErr bool
	}{
		{s: stringType{}, args: args{content: &x}, want: "hello", wantErr: false},
		{s: stringType{}, args: args{content: &xq}, want: "hello world", wantErr: false},
		{s: stringType{}, args: args{content: &xc}, wantErr: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := stringType{}
			got, err := s.Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringType.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringType.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
