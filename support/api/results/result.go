package results

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/wsrf16/swiss/sugar/base/timekit"
	"net/http"
	"time"
)

const (
	SUCCEED = 0
	FAILED  = 1
)

type Result struct {
	SpanId    string    `json:"spanId"`
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	Data      any       `json:"data"`
	DateTime  time.Time `json:"dateTime"`
	Timestamp int64     `json:"timestamp"`
}

func New(code int, msg string, data any) Result {
	var u uuid.UUID
	if _u, err := uuid.NewUUID(); err != nil {
		panic(err)
	} else {
		u = _u
	}

	spanId := u.String()
	now := time.Now()
	timestamp := timekit.ToMilliSecond(now)
	return Result{SpanId: spanId, Code: code, Msg: msg, Data: data, Timestamp: timestamp, DateTime: now}
}

func Success(data any) Result {
	return New(SUCCEED, "操作成功", data)
}

func Failed(err error) Result {
	return New(FAILED, err.Error(), nil)
}

func FailedText(err string) Result {
	return New(FAILED, err, nil)
}

func WriteHttp(w http.ResponseWriter, result Result, code int) error {
	if b, err1 := json.Marshal(result); err1 != nil {
		http.Error(w, err1.Error(), code)
		return nil
	} else {
		_, err2 := w.Write(b)
		return err2
	}
}

func WriteHttpFailed(w http.ResponseWriter, err string) error {
	return WriteHttp(w, FailedText(err), http.StatusOK)
}

func WriteHttpFailedData(w http.ResponseWriter, err string, data any) error {
	result := New(FAILED, err, data)
	return WriteHttp(w, result, http.StatusOK)
}

func WriteHttpSucceed(w http.ResponseWriter, data any) error {
	return WriteHttp(w, Success(data), http.StatusOK)
}
