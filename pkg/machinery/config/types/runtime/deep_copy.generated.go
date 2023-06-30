// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type EventSinkV1Alpha1 -type KmsgLogV1Alpha1 -pointer-receiver -header-file ../../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package runtime

import (
	"net/url"
)

// DeepCopy generates a deep copy of *EventSinkV1Alpha1.
func (o *EventSinkV1Alpha1) DeepCopy() *EventSinkV1Alpha1 {
	var cp EventSinkV1Alpha1 = *o
	return &cp
}

// DeepCopy generates a deep copy of *KmsgLogV1Alpha1.
func (o *KmsgLogV1Alpha1) DeepCopy() *KmsgLogV1Alpha1 {
	var cp KmsgLogV1Alpha1 = *o
	if o.KmsgLogURL.URL != nil {
		cp.KmsgLogURL.URL = new(url.URL)
		*cp.KmsgLogURL.URL = *o.KmsgLogURL.URL
		if o.KmsgLogURL.URL.User != nil {
			cp.KmsgLogURL.URL.User = new(url.Userinfo)
			*cp.KmsgLogURL.URL.User = *o.KmsgLogURL.URL.User
		}
	}
	return &cp
}
