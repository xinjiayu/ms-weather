package units

//空气质量描述转换
/*
AQI数值	空气质量等级
0 ~50	优
50~100	良
100~150	轻度污染
150~200	中度污染
>200	重度污染
*/
func AQIDescribe(aqi int) string {
	if aqi < 50 {
		return "优"
	} else if aqi < 100 && aqi >= 50 {
		return "良"
	} else if aqi < 150 && aqi >= 100 {
		return "轻度污染"
	} else if aqi < 200 && aqi >= 150 {
		return "中度污染"
	} else if aqi >= 200 {
		return "重度污染"
	}
	return ""
}
