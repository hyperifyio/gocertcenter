// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package hashutils

import (
	"fmt"
	"hash"
)

func ToUint64(
	str string,
	h hash.Hash64,
) (uint64, error) {
	_, err := h.Write([]byte(str))
	if err != nil {
		return 0, fmt.Errorf("[hashutils] error writing hash: %v", err)
	}
	return h.Sum64(), nil
}
