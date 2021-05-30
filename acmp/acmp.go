package acmp

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	my_cookie string = "Cookie: ASPSESSIONIDSWRSDQBD=GEEHMCECHPIPHIFJPJPGMJNJ; ASPSESSIONIDSSSTAQBD=FAMJEODCDFHHGOMODECMOJDI; fid=c1358474-510e-4c1b-b796-a3ea4ba52964; ASPSESSIONIDQUQRRDST=LLMBCILANHAALIMIOGNEDIIC; _ym_uid=1622323415281367458; _ym_d=1622323415; _ym_visorc=w; _ym_isad=2; __gads=ID=317f2f64ed660c2b-224ccfed2ec800a5:T=1622323415:RT=1622323415:S=ALNI_MahzeCSXFK5U94zRf9NtpgxpstVWQ; login=qwertyasdfgh123; Banners%5FDown=No; English=1"
	pattern   string = `(Difficulty: \d+%\))`
)

func getPercentFromWord(str string) float64 {
	ans := 0
	for _, letter := range str {
		if 48 <= int(letter) && int(letter) <= 57 {
			ans += int(letter) - 48
			ans *= 10
		}
	}
	return float64(ans / 10)
}

func Difficulty(url string) float64 {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", my_cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(pattern)
	response := re.FindString(string(body))
	if len(response) == 0 {
		return -1
	}
	return getPercentFromWord(response)
}
