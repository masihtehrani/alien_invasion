package filesystem

import (
	"context"
	"reflect"
	"testing"
)

func TestDataStore_GetDataWorld(t *testing.T) {
	tests := []struct {
		name    string
		dataDir string
		want    []string
		wantErr bool
	}{
		{name: "success", dataDir: "../../../../testdata/world.txt",
			want: []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"}, wantErr: false},
		{name: "failed", dataDir: "notfound.txt", want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataStore{
				dataDir: tt.dataDir,
			}
			got, err := d.GetDataWorld(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataWorld() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataWorld() got = %v, want %v", got, tt.want)
			}
		})
	}
}
