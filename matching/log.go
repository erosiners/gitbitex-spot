// Copyright 2019 GitBitEx.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package matching

import (
	"github.com/gitbitex/gitbitex-spot/models"
	"github.com/shopspring/decimal"
	"time"
)

type LogType string

const (
	LogTypeMatch = LogType("match")
	LogTypeOpen  = LogType("open")
	LogTypeDone  = LogType("done")
)

type Log interface {
	GetSeq() int64
}

type Base struct {
	Type      LogType
	Sequence  int64
	ProductId string
	Time      time.Time
}

type ReceivedLog struct {
	Base
	OrderId   int64
	Size      decimal.Decimal
	Price     decimal.Decimal
	Side      models.Side
	OrderType models.OrderType
}

func (l *ReceivedLog) GetSeq() int64 {
	return l.Sequence
}

type OpenLog struct {
	Base
	OrderId       int64
	RemainingSize decimal.Decimal
	Price         decimal.Decimal
	Side          models.Side
}

func (l *OpenLog) GetSeq() int64 {
	return l.Sequence
}

type DoneLog struct {
	Base
	OrderId       int64
	Price         decimal.Decimal
	RemainingSize decimal.Decimal
	Reason        models.DoneReason
	Side          models.Side
}

func (l *DoneLog) GetSeq() int64 {
	return l.Sequence
}

type MatchLog struct {
	Base
	TradeId      int64
	TakerOrderId int64
	MakerOrderId int64
	Side         models.Side
	Price        decimal.Decimal
	Size         decimal.Decimal
}

func (l *MatchLog) GetSeq() int64 {
	return l.Sequence
}
