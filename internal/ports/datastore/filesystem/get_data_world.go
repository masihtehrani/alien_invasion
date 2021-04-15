package filesystem

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// GetDataWorld get data from sile system and return data.
func (d *DataStore) GetDataWorld(ctx context.Context) ([]string, error) {
	f, err := os.Open(d.dataDir)
	if err != nil {
		return nil, fmt.Errorf("GetDataWorld >> os.Open >> %w", err)
	}

	var result []string

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		a := scan.Text()
		result = append(result, a)
	}

	err = scan.Err()
	if err != nil {
		return nil, fmt.Errorf("GetDataWorld >> scan.Err() >> %w", err)
	}

	if err = f.Close(); err != nil {
		return nil, fmt.Errorf("GetDataWorld >> f.Close() >> %w", err)
	}

	return result, nil
}
