package parameter

import (
	"reflect"
	"testing"
)

func TestProduct_Valid(t *testing.T) {
	type fields struct {
		ProductID int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				ProductID: 101,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "異常系_ProductID=0の場合",
			fields: fields{
				ProductID: 0,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Product{
				ProductID: tt.fields.ProductID,
			}
			got, err := f.Valid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Valid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsByName_Valid(t *testing.T) {
	type fields struct {
		Name   string
		Offset int
		Limit  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				Name:   "商品",
				Offset: 0,
				Limit:  1,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "異常系_Name>20の場合",
			fields: fields{
				Name:   "123456789012345678901",
				Offset: 0,
				Limit:  1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "異常系_Offset<0の場合",
			fields: fields{
				Name:   "商品",
				Offset: -1,
				Limit:  1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "異常系_Limit>50の場合",
			fields: fields{
				Name:   "商品",
				Offset: 0,
				Limit:  51,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &ProductsByName{
				Name:   tt.fields.Name,
				Offset: tt.fields.Offset,
				Limit:  tt.fields.Limit,
			}
			got, err := f.Valid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Valid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddProduct_Valid(t *testing.T) {
	type fields struct {
		Name    string
		BrandID int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				Name:    "商品名1",
				BrandID: 201,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "異常系_Nameが空の場合",
			fields: fields{
				Name:    "",
				BrandID: 201,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "異常系_BrandID=0の場合",
			fields: fields{
				Name:    "商品名1",
				BrandID: 0,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &AddProduct{
				Name:    tt.fields.Name,
				BrandID: tt.fields.BrandID,
			}
			got, err := f.Valid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Valid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEditProduct_Valid(t *testing.T) {
	type fields struct {
		ProductID int64
		Name      string
		BrandID   int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				ProductID: 101,
				Name:      "商品名1",
				BrandID:   201,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "異常系_ProductID=0の場合",
			fields: fields{
				ProductID: 0,
				Name:      "",
				BrandID:   201,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "異常系_Nameが空の場合",
			fields: fields{
				ProductID: 101,
				Name:      "",
				BrandID:   201,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "異常系_BrandID=0の場合",
			fields: fields{
				ProductID: 101,
				Name:      "商品名1",
				BrandID:   0,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &EditProduct{
				ProductID: tt.fields.ProductID,
				Name:      tt.fields.Name,
				BrandID:   tt.fields.BrandID,
			}
			got, err := f.Valid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Valid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProductsByName(t *testing.T) {
	type args struct {
		nameValue   string
		offsetValue string
		limitValue  string
	}
	tests := []struct {
		name string
		args args
		want *ProductsByName
	}{
		{
			name: "正常系",
			args: args{
				nameValue:   "商品名1",
				offsetValue: "0",
				limitValue:  "10",
			},
			want: &ProductsByName{
				Name:   "商品名1",
				Offset: 0,
				Limit:  10,
			},
		},
		{
			name: "正常系_パラメータを指定しない場合_デフォルト値がセットされること",
			args: args{},
			want: &ProductsByName{
				Name:   "",
				Offset: 0,
				Limit:  20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductsByName(tt.args.nameValue, tt.args.offsetValue, tt.args.limitValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductsByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
