//Package util - used to generate shorten url string
package util

//Base62URL in order to generate shorten url string,it's consider to use base 62 approach(6yte)
//@param id : unix time
func Base62URL(id int64) string {
	if id < 0 {
		return ""
	}
	//[a-z,A-Z,0-9]
	urlStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //total 62 character
	url := ""

	for {
		if id <= 0 {
			break
		}
		str := urlStr[int(id%62)]
		url += string(str)
		id = id / 62
	}

	return url
}
