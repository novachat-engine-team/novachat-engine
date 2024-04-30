/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/10 18:26
 * @File : report.go
 */

package constants

import (
	"fmt"
	"novachat_engine/mtproto"
)

//inputReportReasonSpam#58dbcab8 = ReportReason;
//inputReportReasonViolence#1e22c78d = ReportReason;
//inputReportReasonPornography#2e59d922 = ReportReason;
//inputReportReasonChildAbuse#adf44ee3 = ReportReason;
//inputReportReasonOther#e1746d0a text:string = ReportReason;
//inputReportReasonCopyright#9b89f93a = ReportReason;
//inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;

type ReportType int32

const (
	ReportTypeNone          ReportType = 0
	ReportTypeSpam          ReportType = 1
	ReportTypeViolence      ReportType = 2
	ReportTypePornography   ReportType = 3
	ReportTypeChildAbuse    ReportType = 4
	ReportTypeOther         ReportType = 5
	ReportTypeCopyright     ReportType = 6
	ReportTypeGeoIrrelevant ReportType = 7
)

func (m ReportType) ToInt32() int32 {
	return int32(m)
}

func ToReportType(reason *mtproto.ReportReason) (ReportType, string) {
	switch reason.ClassName {
	case mtproto.ClassInputReportReasonSpam:
		return ReportTypeSpam, ""
	case mtproto.ClassInputReportReasonViolence:
		return ReportTypeViolence, ""
	case mtproto.ClassInputReportReasonPornography:
		return ReportTypePornography, ""
	case mtproto.ClassInputReportReasonChildAbuse:
		return ReportTypeChildAbuse, ""
	case mtproto.ClassInputReportReasonOther:
		return ReportTypeOther, reason.Text
	case mtproto.ClassInputReportReasonCopyright:
		return ReportTypeCopyright, ""
	case mtproto.ClassInputReportReasonGeoIrrelevant:
		return ReportTypeGeoIrrelevant, ""
	default:
		panic(fmt.Sprintf("ToReportType typeName:%s", reason.ClassName))
	}
}
