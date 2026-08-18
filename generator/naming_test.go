package generator

import (
	"testing"
)

func TestToExportedName(t *testing.T) {
	tests := []struct {
		table          string
		prefix         string
		expectedModel  string
		expectedStruct string
		expectedMiddle string
	}{
		{
			table:          "safe_uts",
			prefix:         "",
			expectedModel:  "SafeUts",
			expectedStruct: "SafeUts",
			expectedMiddle: "SafeUtsRequest",
		},
		{
			table:          "sys_dictionaries",
			prefix:         "",
			expectedModel:  "SysDictionaries",
			expectedStruct: "SysDictionaries",
			expectedMiddle: "SysDictionariesRequest",
		},
		{
			table:          "sys_dictionary",
			prefix:         "",
			expectedModel:  "SysDictionary",
			expectedStruct: "SysDictionary",
			expectedMiddle: "SysDictionaryRequest",
		},
		{
			table:          "sys_dictionary",
			prefix:         "sys_",
			expectedModel:  "Dictionary",
			expectedStruct: "Dictionary",
			expectedMiddle: "DictionaryRequest",
		},
		{
			table:          "safe_uts_api",
			prefix:         "safe_",
			expectedModel:  "UtsApi",
			expectedStruct: "UtsApi",
			expectedMiddle: "UtsApiRequest",
		},
	}

	for _, tt := range tests {
		rawName := tt.table
		if tt.prefix != "" && len(rawName) > len(tt.prefix) {
			rawName = rawName[len(tt.prefix):]
		}
		gotModel := toExportedName(rawName)
		if gotModel != tt.expectedModel {
			t.Errorf("table %s (prefix %q): expected model %s, got %s", tt.table, tt.prefix, tt.expectedModel, gotModel)
		}

		gotStruct := validatorStructName(tt.table, tt.prefix)
		if gotStruct != tt.expectedStruct {
			t.Errorf("table %s (prefix %q): expected struct %s, got %s", tt.table, tt.prefix, tt.expectedStruct, gotStruct)
		}

		gotMiddle := validatorMiddlewareName(tt.table, tt.prefix)
		if gotMiddle != tt.expectedMiddle {
			t.Errorf("table %s (prefix %q): expected middleware %s, got %s", tt.table, tt.prefix, tt.expectedMiddle, gotMiddle)
		}
	}
}
