package gorestful

import "testing"

func TestIsDebugging(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDebugging(); got != tt.want {
				t.Errorf("IsDebugging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMode(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_debugPrint(t *testing.T) {
	type args struct {
		format string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_debugPrintMessage(t *testing.T) {
	type args struct {
		httpMethod   string
		absolutePath string
		h            Handle
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_debugPrintWARNINGNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_nameOfFunction(t *testing.T) {
	type args struct {
		f interface{}
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
			if got := nameOfFunction(tt.args.f); got != tt.want {
				t.Errorf("nameOfFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}