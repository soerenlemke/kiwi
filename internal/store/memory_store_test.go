package store

import (
	"io"
	"kiwi/internal/domain"
	"kiwi/internal/recovery"
	"log/slog"
	"reflect"
	"testing"
)

func TestNewInMemoryStore(t *testing.T) {
	type args[T domain.AllowedTypes] struct {
		logger    slog.Logger
		recoverer recovery.Recoverer[T]
	}
	type testCase[T domain.AllowedTypes] struct {
		name string
		args args[T]
		want *InMemoryStore[T]
	}
	tests := []testCase[string]{
		{
			name: "creates_empty_store_and_get_unknown_key_returns_false",
			args: args[string]{
				logger:    *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
			},
			want: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInMemoryStore[string](tt.args.logger, tt.args.recoverer)

			if got == nil {
				t.Fatalf("NewInMemoryStore() returned nil")
			}
			if got.data == nil || len(got.data) != 0 {
				t.Fatalf("expected empty data map, got: %#v", got.data)
			}

			val, ok, err := got.Get("unknown")
			if err != nil {
				t.Fatalf("Get on unknown key returned error: %v", err)
			}
			if ok {
				t.Fatalf("Get on unknown key returned ok=true, want false")
			}
			if !reflect.ValueOf(val).IsZero() {
				t.Fatalf("Get on unknown key returned non-zero value: %#v", val)
			}
		})
	}
}

func TestInMemoryStore_Set(t *testing.T) {
	type args[T any] struct {
		key   string
		value T
	}
	type testCase[T domain.AllowedTypes] struct {
		name    string
		s       *InMemoryStore[T]
		args    args[T]
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "sets_and_gets_value_from_store",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{},
			},
			args: args[string]{
				key:   "key",
				value: "value",
			},
			wantErr: false,
		},
	}

	for i := range tests {
		tt := &tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Fatalf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, ok, err := tt.s.Get(tt.args.key)
			if err != nil {
				t.Fatalf("Get() unexpected error: %v", err)
			}
			if !ok {
				t.Fatalf("Get() ok=false, expected true")
			}
			if got != tt.args.value {
				t.Fatalf("Get() got=%q, want=%q", got, tt.args.value)
			}
		})
	}
}

func TestInMemoryStore_Get(t *testing.T) {
	type args struct {
		key string
	}
	type testCase[T domain.AllowedTypes] struct {
		name    string
		s       *InMemoryStore[T]
		args    args
		want    T
		want1   bool
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "gets_value_from_store",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{"key": "value"},
			},
			args: args{
				key: "key",
			},
			want:  "value",
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInMemoryStore_Delete(t *testing.T) {
	type args struct {
		key string
	}
	type testCase[T domain.AllowedTypes] struct {
		name    string
		s       *InMemoryStore[T]
		args    args
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "delete_existing_key_returns_no_error_and_removes_value",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{"key": "value"},
			},
			args:    args{key: "key"},
			wantErr: false,
		},
		{
			name: "delete_missing_key_returns_error_and_nop",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{},
			},
			args:    args{key: "absent"},
			wantErr: true,
		},
	}

	for i := range tests {
		tt := &tests[i]
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Delete(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			_, ok, getErr := tt.s.Get(tt.args.key)
			if getErr != nil {
				t.Fatalf("Get() unexpected error: %v", getErr)
			}
			if ok {
				t.Fatalf("expected key %q to be absent after Delete", tt.args.key)
			}
		})
	}
}

func TestInMemoryStore_keyInStore(t *testing.T) {
	type args struct {
		key string
	}
	type testCase[T domain.AllowedTypes] struct {
		name string
		s    *InMemoryStore[T]
		args args
		want bool
	}
	tests := []testCase[string]{
		{
			name: "key_absent_returns_false",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{},
			},
			args: args{key: "missing"},
			want: false,
		},
		{
			name: "key_present_returns_true",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{"a": "x"},
			},
			args: args{key: "a"},
			want: true,
		},
		{
			name: "key_absent_after_prior_delete_is_equivalent_to_absent_case",
			s: &InMemoryStore[string]{
				log:       *slog.New(slog.NewTextHandler(io.Discard, nil)),
				recoverer: recovery.NopRlw[string]{},
				data:      map[string]string{},
			},
			args: args{key: "tbd"},
			want: false,
		},
	}

	for i := range tests {
		tt := &tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.keyInStore(tt.args.key); got != tt.want {
				t.Fatalf("keyInStore(%q) = %v, want %v", tt.args.key, got, tt.want)
			}
		})
	}
}
