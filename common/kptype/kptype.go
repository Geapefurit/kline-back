package kptype

import (
	basetype "github.com/Geapefurit/kline-back/proto/kline/basetype/v1"
)

var KPTypeSecond = map[basetype.KPointType]uint32{
	basetype.KPointType_FiveSecond: 5,
	basetype.KPointType_OneMinute:  60,
	basetype.KPointType_TenMinute:  60 * 10,
	basetype.KPointType_OneHour:    60 * 60,
	basetype.KPointType_OneDay:     60 * 60 * 24,
	basetype.KPointType_OneWeek:    60 * 60 * 24 * 7,
	basetype.KPointType_OneMonth:   60 * 60 * 24 * 30,
}

var KPTypeSampleSecond = map[basetype.KPointType]uint32{
	basetype.KPointType_FiveSecond: 5,
	basetype.KPointType_OneMinute:  10,
	basetype.KPointType_TenMinute:  60,
	basetype.KPointType_OneHour:    60,
	basetype.KPointType_OneDay:     60,
	basetype.KPointType_OneWeek:    60,
	basetype.KPointType_OneMonth:   60,
}
