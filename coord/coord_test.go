/*
Copyright 2015 Tamás Gulácsi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package coord

import (
	"testing"

	"golang.org/x/net/context"
)

func TestGetCoord(t *testing.T) {
	for i, tc := range []struct {
		Address string
		WantErr bool
	}{
		{"Telepy utca 24, Budapest, HUNGARY", false},
		{"XXXXXXX utca", true},
	} {

		loc, err := GetCoord(context.Background(), tc.Address)
		t.Logf("%#v [error=%v]", loc, err)
		if tc.WantErr && err == nil {
			t.Errorf("%d. want error for %q.", i, tc.Address)
		} else if err != nil && !tc.WantErr {
			t.Errorf("%d. got error for %q: %v.", i, tc.Address, err)
		}
	}
}