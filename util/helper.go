package util

import "encoding/json"

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int32) *int32 {
	return &i
}

func toStringPtrSlice(s []string) []*string {
	var res []*string
	for _, str := range s {
		strCopy := str // capture new instance to avoid referencing loop var
		res = append(res, &strCopy)
	}
	return res
}

func mapToJSONStringPtr(m map[string]string) *string {
	if m == nil {
		return nil
	}
	b, err := json.Marshal(m)
	if err != nil {
		s := "{}"
		return &s
	}
	s := string(b)
	return &s
}
