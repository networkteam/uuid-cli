package uuidcli

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestGenerateUUid(t *testing.T) {
	type args struct {
		version   int
		domainStr string
		ns        string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "v1",
			args: args{
				version: 1,
			},
		},
		{
			name: "v2 org",
			args: args{
				version:   2,
				domainStr: "org",
			},
		},
		{
			name: "v2 group",
			args: args{
				version:   2,
				domainStr: "group",
			},
		},
		{
			name: "v2 person",
			args: args{
				version:   2,
				domainStr: "person",
			},
		},
		{
			name: "v2 unknown domain",
			args: args{
				version:   2,
				domainStr: "other",
			},
			wantErr: true,
		},
		{
			name: "v3 with namespace and name",
			args: args{
				version: 3,
				ns:      uuid.Must(uuid.NewV1()).String(),
				name:    "baby yoda",
			},
		},
		{
			name: "v3 without namespace",
			args: args{
				version: 3,
			},
			wantErr: true,
		},
		{
			name: "v3 with invalid namespace",
			args: args{
				version: 3,
				ns:      "notauuid",
			},
			wantErr: true,
		},
		{
			name: "v3 without name",
			args: args{
				version: 3,
				ns:      uuid.Must(uuid.NewV1()).String(),
			},
			wantErr: true,
		},
		{
			name: "v4",
			args: args{
				version: 4,
			},
		},
		{
			name: "v5 with namespace and name",
			args: args{
				version: 5,
				ns:      uuid.Must(uuid.NewV1()).String(),
				name:    "baby yoda",
			},
		},
		{
			name: "v5 without namespace",
			args: args{
				version: 5,
			},
			wantErr: true,
		},
		{
			name: "v5 with invalid namespace",
			args: args{
				version: 5,
				ns:      "notauuid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := GenerateUUid(tt.args.version, tt.args.domainStr, tt.args.ns, tt.args.name)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but error was nil")
				}
				return
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if gotU == uuid.Nil {
				t.Error("Expected non-nil UUID")
				return
			}
		})
	}
}
