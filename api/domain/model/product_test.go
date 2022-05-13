package model

import (
	"reflect"
	"testing"
)

func TestProduct_Empty(t *testing.T) {
	type fields struct {
		ID      int64
		Name    string
		BrandID int64
		empty   bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name:    "正常系_empty=trueの場合_trueとエラーを返す",
			fields:  fields{empty: true},
			want:    true,
			wantErr: true,
		},
		{
			name:    "正常系_empty=falseの場合_falseとnilを返す",
			fields:  fields{empty: false},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				BrandID: tt.fields.BrandID,
				empty:   tt.fields.empty,
			}
			got, err := p.Empty()
			if (err != nil) != tt.wantErr {
				t.Errorf("Empty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Empty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyProduct(t *testing.T) {
	tests := []struct {
		name string
		want *Product
	}{
		{
			name: "正常系_emptyがtrueであること",
			want: &Product{empty: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmptyProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmptyProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
