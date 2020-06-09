package tools

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common/hexutil"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/rlp"
)

func DecodeLogData(logData []byte) (interface{}, error) {
	logMap := make(map[string]interface{})
	if err := json.Unmarshal(logData, &logMap); err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	if basestr, ok := logMap["Base"].(string); ok {
		if data, err := base64.StdEncoding.DecodeString(basestr); err == nil {
			buyTicketParam := common.BuyTicketParam{}
			if err = rlp.DecodeBytes(data, &buyTicketParam); err == nil {
				logMap["StartTime"] = buyTicketParam.Start
				logMap["ExpireTime"] = buyTicketParam.End
			}
		}
	}
	delete(logMap, "Base")
	return logMap, nil
}

func DecodeTxInput(input []byte) (interface{}, error) {
	res, err := common.DecodeTxInput(input)
	if err == nil {
		return res, err
	}
	fsnCall, ok := res.(common.FSNCallParam)
	if !ok {
		return res, err
	}
	switch fsnCall.Func {
	case common.ReportIllegalFunc:
		h1, h2, err := DecodeReport(fsnCall.Data)
		if err != nil {
			return nil, fmt.Errorf("DecodeReport err %v", err)
		}
		reportContent := &struct {
			Header1 *types.Header
			Header2 *types.Header
		}{
			Header1: h1,
			Header2: h2,
		}
		fsnCall.Data = nil
		return common.DecodeFsnCallParam(&fsnCall, reportContent)
	}
	return nil, fmt.Errorf("Unknown FuncType %v", fsnCall.Func)
}

func DecodeReport(report []byte) (*types.Header, *types.Header, error) {
	if len(report) < 4 {
		return nil, nil, fmt.Errorf("wrong report length")
	}
	data1len := common.BytesToInt(report[:4])
	if len(report) < 4+data1len {
		return nil, nil, fmt.Errorf("wrong report length")
	}
	data1 := report[4 : data1len+4]
	data2 := report[data1len+4:]

	if bytes.Compare(data1, data2) >= 0 {
		return nil, nil, fmt.Errorf("wrong report sequence")
	}

	header1 := &types.Header{}
	header2 := &types.Header{}

	if err := rlp.DecodeBytes(data1, header1); err != nil {
		return nil, nil, fmt.Errorf("can not decode header1, err=%v", err)
	}
	if err := rlp.DecodeBytes(data2, header2); err != nil {
		return nil, nil, fmt.Errorf("can not decode header2, err=%v", err)
	}
	return header1, header2, nil
}

func DecodePunishTickets(delTicketsData string) ([]common.Hash, error) {
	bs, err := hexutil.Decode(delTicketsData)
	if err != nil {
		return nil, fmt.Errorf("decode hex data error: %v", err)
	}
	delTickets := []common.Hash{}
	if err := rlp.DecodeBytes(bs, &delTickets); err != nil {
		return nil, fmt.Errorf("decode report log error: %v", err)
	}

	return delTickets, nil
}

// MustPrintJSON prints the JSON encoding of the given object and
// exits the program with an error message when the marshaling fails.
func MustPrintJSON(jsonObject interface{}) {
	str, err := json.MarshalIndent(jsonObject, "", "  ")
	if err != nil {
		utils.Fatalf("Failed to marshal JSON object: %v", err)
	}
	fmt.Println(string(str))
}
