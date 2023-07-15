package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	types "github.com/LintaoAmons/undercontrol/src/persistence/account/single-file/types"
	"github.com/davecgh/go-spew/spew"
)

func Test_AccountJson(t *testing.T) {
	bb := &types.AccountJson{
		Name:   "Lintao",
		Amount: money.NewFromFloat(2.3, "CNY"),
		Audit:  common.DefaultAudit(),
	}
	result, _ := json.Marshal(bb)
	fmt.Println(string(result))
}

func Test_AccountUnmarshal(t *testing.T) {
	str := `{"Name":"Lintao","Amount":{"amount":"2.3","currency":{"code":"CNY","numeric_code":"156","fraction":2,"grapheme":"元","template":"1 $","decimal":".","thousand":","}},"CreatedAt":"2023-07-02T18:24:52.956197+08:00","CreatedBy":"system","UpdatedAt":"2023-07-02T18:24:52.956197+08:00","UpdatedBy":"system"}`
	var accjson types.AccountJson
	err := json.Unmarshal([]byte(str), &accjson)
	if err != nil {
		panic("Unmarshal error")
	}
	spew.Dump(accjson)
}
