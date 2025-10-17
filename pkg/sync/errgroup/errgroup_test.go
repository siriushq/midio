package errgroup

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupWithNErrs(t *testing.T) {
	err1 := fmt.Errorf("errgroup_test: 1")
	err2 := fmt.Errorf("errgroup_test: 2")

	cases := []struct {
		errs []error
	}{
		{errs: []error{nil}},
		{errs: []error{err1}},
		{errs: []error{err1, nil}},
		{errs: []error{err1, nil, err2}},
	}

	for j, tc := range cases {
		t.Run(fmt.Sprintf("Test%d", j+1), func(t *testing.T) {
			g := WithNErrs(len(tc.errs))
			for i, err := range tc.errs {
				err := err
				g.Go(func() error { return err }, i)
			}

			gotErrs := g.Wait()
			if !reflect.DeepEqual(gotErrs, tc.errs) {
				t.Errorf("Expected %#v, got %#v", tc.errs, gotErrs)
			}
		})
	}
}
