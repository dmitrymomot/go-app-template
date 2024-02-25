package emailx_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dmitrymomot/go-app-template/pkg/emailx"
)

func TestValidateEmail(t *testing.T) {
	t.Parallel()

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid email",
			args:    args{email: "test@gmail.com"},
			wantErr: false,
		},
		{
			name:    "valid test email",
			args:    args{email: "john+test@gmail.com"},
			wantErr: false,
		},
		{
			name:    "valid email: dot in user",
			args:    args{email: "john.test@gmail.com"},
			wantErr: false,
		},
		{
			name:    "invalid email",
			args:    args{email: "invalid_email"},
			wantErr: true,
		},
		{
			name:    "empty email",
			args:    args{email: ""},
			wantErr: true,
		},
		{
			name:    "invalid email: no domain",
			args:    args{email: "invalid_email@localhost"},
			wantErr: true,
		},
		{
			name:    "invalid email: no user",
			args:    args{email: "@test.com"},
			wantErr: true,
		},
		{
			name:    "invalid email: no @",
			args:    args{email: "invalid_email.test.com"},
			wantErr: true,
		},
		{
			name:    "invalid email: no dot",
			args:    args{email: "invalid_email@testcom"},
			wantErr: true,
		},
		{
			name:    "invalid email: no tld",
			args:    args{email: "invalid_email@test."},
			wantErr: true,
		},
		{
			name:    "invalid email: short tld",
			args:    args{email: "invalid_email@test.c"},
			wantErr: true,
		},
		{
			name:    "invalid email: invalid tld",
			args:    args{email: "invalid_email@test.business-test"},
			wantErr: true,
		},
		{
			name:    "invalid email: too long tld",
			args:    args{email: "invalid_email@test.businesstest"},
			wantErr: true,
		},
		{
			name:    "invalid email: special chars in user",
			args:    args{email: "test#" + gofakeit.Email()},
			wantErr: true,
		},
		{
			name:    "invalid email: host is not valid",
			args:    args{email: "test@gmail.test"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := emailx.ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSanitizeEmail(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid email",
			args:    args{s: "test@mail.dev"},
			want:    "test@mail.dev",
			wantErr: false,
		},
		{
			name:    "valid email: dot in user",
			args:    args{s: "test.test@mail.dev"},
			want:    "testtest@mail.dev",
			wantErr: false,
		},
		{
			name:    "valid email: dash in user",
			args:    args{s: "test-test@mail.dev"},
			want:    "testtest@mail.dev",
			wantErr: false,
		},
		{
			name:    "valid email: dash in user",
			args:    args{s: "test+test@mail.dev"},
			want:    "test@mail.dev",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := emailx.SanitizeEmail(tt.args.s, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("sanitizeEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sanitizeEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
