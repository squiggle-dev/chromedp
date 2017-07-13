// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package emulation

import (
	json "encoding/json"
	cdp "github.com/knq/chromedp/cdp"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation(in *jlexer.Lexer, out *SetVirtualTimePolicyParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "policy":
			(out.Policy).UnmarshalEasyJSON(in)
		case "budget":
			out.Budget = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation(out *jwriter.Writer, in SetVirtualTimePolicyParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"policy\":")
	(in.Policy).MarshalEasyJSON(out)
	if in.Budget != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"budget\":")
		out.Int64(int64(in.Budget))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetVirtualTimePolicyParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetVirtualTimePolicyParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetVirtualTimePolicyParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetVirtualTimePolicyParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation1(in *jlexer.Lexer, out *SetTouchEmulationEnabledParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "enabled":
			out.Enabled = bool(in.Bool())
		case "configuration":
			(out.Configuration).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation1(out *jwriter.Writer, in SetTouchEmulationEnabledParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"enabled\":")
	out.Bool(bool(in.Enabled))
	if in.Configuration != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"configuration\":")
		(in.Configuration).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetTouchEmulationEnabledParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetTouchEmulationEnabledParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetTouchEmulationEnabledParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetTouchEmulationEnabledParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation2(in *jlexer.Lexer, out *SetScriptExecutionDisabledParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation2(out *jwriter.Writer, in SetScriptExecutionDisabledParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"value\":")
	out.Bool(bool(in.Value))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetScriptExecutionDisabledParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetScriptExecutionDisabledParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetScriptExecutionDisabledParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetScriptExecutionDisabledParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation3(in *jlexer.Lexer, out *SetPageScaleFactorParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "pageScaleFactor":
			out.PageScaleFactor = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation3(out *jwriter.Writer, in SetPageScaleFactorParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"pageScaleFactor\":")
	out.Float64(float64(in.PageScaleFactor))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetPageScaleFactorParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetPageScaleFactorParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetPageScaleFactorParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetPageScaleFactorParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation4(in *jlexer.Lexer, out *SetGeolocationOverrideParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "accuracy":
			out.Accuracy = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation4(out *jwriter.Writer, in SetGeolocationOverrideParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Latitude != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"latitude\":")
		out.Float64(float64(in.Latitude))
	}
	if in.Longitude != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"longitude\":")
		out.Float64(float64(in.Longitude))
	}
	if in.Accuracy != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accuracy\":")
		out.Float64(float64(in.Accuracy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetGeolocationOverrideParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetGeolocationOverrideParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetGeolocationOverrideParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetGeolocationOverrideParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation5(in *jlexer.Lexer, out *SetEmulatedMediaParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "media":
			out.Media = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation5(out *jwriter.Writer, in SetEmulatedMediaParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"media\":")
	out.String(string(in.Media))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetEmulatedMediaParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetEmulatedMediaParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetEmulatedMediaParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetEmulatedMediaParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation6(in *jlexer.Lexer, out *SetDeviceMetricsOverrideParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "width":
			out.Width = int64(in.Int64())
		case "height":
			out.Height = int64(in.Int64())
		case "deviceScaleFactor":
			out.DeviceScaleFactor = float64(in.Float64())
		case "mobile":
			out.Mobile = bool(in.Bool())
		case "fitWindow":
			out.FitWindow = bool(in.Bool())
		case "scale":
			out.Scale = float64(in.Float64())
		case "screenWidth":
			out.ScreenWidth = int64(in.Int64())
		case "screenHeight":
			out.ScreenHeight = int64(in.Int64())
		case "positionX":
			out.PositionX = int64(in.Int64())
		case "positionY":
			out.PositionY = int64(in.Int64())
		case "screenOrientation":
			if in.IsNull() {
				in.Skip()
				out.ScreenOrientation = nil
			} else {
				if out.ScreenOrientation == nil {
					out.ScreenOrientation = new(ScreenOrientation)
				}
				(*out.ScreenOrientation).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation6(out *jwriter.Writer, in SetDeviceMetricsOverrideParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"width\":")
	out.Int64(int64(in.Width))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"height\":")
	out.Int64(int64(in.Height))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"deviceScaleFactor\":")
	out.Float64(float64(in.DeviceScaleFactor))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mobile\":")
	out.Bool(bool(in.Mobile))
	if in.FitWindow {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fitWindow\":")
		out.Bool(bool(in.FitWindow))
	}
	if in.Scale != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scale\":")
		out.Float64(float64(in.Scale))
	}
	if in.ScreenWidth != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"screenWidth\":")
		out.Int64(int64(in.ScreenWidth))
	}
	if in.ScreenHeight != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"screenHeight\":")
		out.Int64(int64(in.ScreenHeight))
	}
	if in.PositionX != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"positionX\":")
		out.Int64(int64(in.PositionX))
	}
	if in.PositionY != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"positionY\":")
		out.Int64(int64(in.PositionY))
	}
	if in.ScreenOrientation != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"screenOrientation\":")
		if in.ScreenOrientation == nil {
			out.RawString("null")
		} else {
			(*in.ScreenOrientation).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetDeviceMetricsOverrideParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetDeviceMetricsOverrideParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetDeviceMetricsOverrideParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetDeviceMetricsOverrideParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation7(in *jlexer.Lexer, out *SetDefaultBackgroundColorOverrideParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "color":
			if in.IsNull() {
				in.Skip()
				out.Color = nil
			} else {
				if out.Color == nil {
					out.Color = new(cdp.RGBA)
				}
				(*out.Color).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation7(out *jwriter.Writer, in SetDefaultBackgroundColorOverrideParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Color != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"color\":")
		if in.Color == nil {
			out.RawString("null")
		} else {
			(*in.Color).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetDefaultBackgroundColorOverrideParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetDefaultBackgroundColorOverrideParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetDefaultBackgroundColorOverrideParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetDefaultBackgroundColorOverrideParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation8(in *jlexer.Lexer, out *SetCPUThrottlingRateParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "rate":
			out.Rate = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation8(out *jwriter.Writer, in SetCPUThrottlingRateParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"rate\":")
	out.Float64(float64(in.Rate))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetCPUThrottlingRateParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetCPUThrottlingRateParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetCPUThrottlingRateParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetCPUThrottlingRateParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation8(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation9(in *jlexer.Lexer, out *ScreenOrientation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "angle":
			out.Angle = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation9(out *jwriter.Writer, in ScreenOrientation) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	(in.Type).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"angle\":")
	out.Int64(int64(in.Angle))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScreenOrientation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScreenOrientation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScreenOrientation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScreenOrientation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation9(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation10(in *jlexer.Lexer, out *ResetPageScaleFactorParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation10(out *jwriter.Writer, in ResetPageScaleFactorParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResetPageScaleFactorParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResetPageScaleFactorParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResetPageScaleFactorParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResetPageScaleFactorParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation10(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation11(in *jlexer.Lexer, out *EventVirtualTimeBudgetExpired) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation11(out *jwriter.Writer, in EventVirtualTimeBudgetExpired) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventVirtualTimeBudgetExpired) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventVirtualTimeBudgetExpired) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventVirtualTimeBudgetExpired) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventVirtualTimeBudgetExpired) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation11(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation12(in *jlexer.Lexer, out *ClearGeolocationOverrideParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation12(out *jwriter.Writer, in ClearGeolocationOverrideParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearGeolocationOverrideParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearGeolocationOverrideParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearGeolocationOverrideParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearGeolocationOverrideParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation12(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation13(in *jlexer.Lexer, out *ClearDeviceMetricsOverrideParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation13(out *jwriter.Writer, in ClearDeviceMetricsOverrideParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearDeviceMetricsOverrideParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearDeviceMetricsOverrideParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearDeviceMetricsOverrideParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearDeviceMetricsOverrideParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation13(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation14(in *jlexer.Lexer, out *CanEmulateReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			out.Result = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation14(out *jwriter.Writer, in CanEmulateReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Result {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"result\":")
		out.Bool(bool(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CanEmulateReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CanEmulateReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CanEmulateReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CanEmulateReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation14(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation15(in *jlexer.Lexer, out *CanEmulateParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation15(out *jwriter.Writer, in CanEmulateParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CanEmulateParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CanEmulateParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpEmulation15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CanEmulateParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CanEmulateParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpEmulation15(l, v)
}
