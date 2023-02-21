package persistor

import (
	"os"
	"reflect"
	"testing"
)

func TestHarDrivePersistor(t *testing.T) {
	dname, err := os.MkdirTemp("", "persistor")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)

	t.Run("Persist", func(t *testing.T) {
		var (
			wantName = "name"
			wantData = []byte("data")
			wantPath = dname + string(os.PathSeparator) + wantName
		)
		p := NewHarDrivePersistor()
		err := p.Persist(wantName, wantData, dname)
		if err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(wantPath)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(data, wantData) {
			t.Fatalf("unexpected Data, want '%v' actual '%v'", wantData, data)
		}
	})

	t.Run("Persist data with undefined name", func(t *testing.T) {
		p := NewHarDrivePersistor()
		err := p.Persist("", []byte{}, "")
		if err != ErrUndefinedName {
			t.Fatal(err)
		}
	})

}
