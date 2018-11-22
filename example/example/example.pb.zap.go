// Code generated by protoc-gen-defaults. DO NOT EDIT.

package example

import (
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap/zapcore"
)

func (u *User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if u == nil {
		return nil
	}

	enc.AddString("Id:", u.Id)

	enc.AddString("PhysicalDesk:", u.PhysicalDesk)

	enc.AddObject("Service:", u.Service)

	BlockedArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range u.Blocked {
			enc.AppendString(v)

		}
		return nil
	}
	enc.AddArray("Blocked:", zapcore.ArrayMarshalerFunc(BlockedArrMarshaller))

	enc.AddReflected("Extra:", u.Extra)

	timeHireDate, _ := ptypes.Timestamp(u.HireDate)
	enc.AddString("HireDate:", timeHireDate.Format("Mon Jan 2 2006 15:04:05 -0700 MST "))

	ListArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range u.List {
			enc.AppendObject(v)

		}
		return nil
	}
	enc.AddArray("List:", zapcore.ArrayMarshalerFunc(ListArrMarshaller))

	return nil
}

func (u *ServiceMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if u == nil {
		return nil
	}

	enc.AddString("Id:", u.Id)

	enc.AddString("Name:", u.Name)

	return nil
}
