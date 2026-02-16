package ops

import (
	"context"
	"fmt"
	"io"
	"os"
)

type StandardOperator struct{}

func NewStandardOperator() *StandardOperator {
	return &StandardOperator{}
}

func (o *StandardOperator) Copy(ctx context.Context, src, dest string) (<-chan Progress, error) {
	ch := make(chan Progress, 1)
	go func() {
		defer close(ch)
		err := copyFile(src, dest, ch)
		if err != nil {
			ch <- Progress{Error: err}
		}
	}()
	return ch, nil
}

func (o *StandardOperator) Move(ctx context.Context, src, dest string) (<-chan Progress, error) {
	// Simple rename if on same FS, otherwise copy+delete
	ch := make(chan Progress, 1)
	go func() {
		defer close(ch)
		err := os.Rename(src, dest)
		if err != nil {
			ch <- Progress{Error: err}
		} else {
			ch <- Progress{CurrentFile: dest, FilesDone: 1}
		}
	}()
	return ch, nil
}

func (o *StandardOperator) Delete(ctx context.Context, paths []string) error {
	for _, p := range paths {
		if err := os.RemoveAll(p); err != nil {
			return err
		}
	}
	return nil
}

func (o *StandardOperator) Trash(ctx context.Context, paths []string) error {
	// Implement Freedesktop Trash spec here
	// For now, simple delete
	return o.Delete(ctx, paths)
}

func copyFile(src, dst string, progress chan<- Progress) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}
	return nil
}
