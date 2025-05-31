package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Kopeks uint

func (kops *Kopeks) UnmarshalJSON(from []byte) error {
	var s string
	if err := json.Unmarshal(from, &s); err != nil {
		return err
	}

	u64, err := strconv.ParseUint(strings.ReplaceAll(s, ".", ""), 10, 64)
	if err != nil {
		return err
	}

	*kops = Kopeks(u64)

	return nil
}

func KopeksFromRub(rub uint) Kopeks {
	return Kopeks(rub * 100)
}

func (kops Kopeks) String() string {
	rem := kops % 100
	return fmt.Sprintf("%d.%02d", kops.Rub(), rem)
}

func (kops Kopeks) Rub() uint {
	return uint(kops) / 100
}

func (kops Kopeks) MarshalJSON() ([]byte, error) {
	return json.Marshal(kops.String())
}
