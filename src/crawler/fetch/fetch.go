package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {
	if strings.Index(url, "https") == -1 {
		url = strings.Replace(url, "http", "https", 1)
	}
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	request.Header.Add("cookie", "enable_FSSBBIl1UgzbN7N=true; sid=3094d9a2-1dd8-4346-a6ba-b056742359ca; ec=rPr4ZFqN-1631082127157-a9662a3bbd7b3-666224740; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1631082406; FSSBBIl1UgzbN7NO=5FHG1M6S7YWLlu7aU0Ly3SflucBte6PkqgdU.88tMmbLO6DyvYlq23Q4LF9dHgWAPO5EL2sfj.a4Mj4CGC3ZQSA; _exid=wu%2FvI2kg9r7XSINN%2FclCPwpg%2Fzj6uIF7WKNiYPGHYtTT3H%2FIt7kROAQl3OMbBVW7rcaTyOM13c%2BW%2BrdVRx10%2FQ%3D%3D; _efmdata=dyY8qPJvZQjO4CreY%2FTLLk42f4COphOi34kNpyobU12eSzK62Musuq0WGynAhTPLa1fsl3GH3RbCPnyC8ipu%2B8KYMzns%2BeiqFMrvwFLe1%2FY%3D; FSSBBIl1UgzbN7NP=53VDzBCmr.MLqqqm4_E_uvAOMU7iDe2G.wQxmee10394_gHtpkVnL7_ISBNKuLSS313WH613CbYuX.u6jqE6rWNc1TCP7fAXN1bokEgSxYsh2qNgI3RtRzz0SZTzCKT0nfpTgo6Bly_sV5ObnzBe_k2RVTe1TC1s.kQcEvZ6OzEJsKTFeDRs3wwP.Uq4sros4NLOyOULA7Lie5HVZdrjvxGxst2xpogjVMBwRyTJdiEqHstQZxNTvPobHAtPEK8yG0; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1631694268")
	client := http.Client{}
	get, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer get.Body.Close()
	if get.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", get.StatusCode)
	}
	newReader := bufio.NewReader(get.Body)
	e := determineEncoding(newReader)
	reader := transform.NewReader(newReader, e.NewDecoder())
	return ioutil.ReadAll(reader)
	//return ioutil.ReadAll(get.Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//bytes, err := r.Peek(1024)
	//if err != nil {
	//	log.Printf("Fethc error:%v", err)
	//return unicode.UTF8
	//return nil
	//}
	//e, _, _ := charset.DetermineEncoding(bytes, "")
	//return e
	return nil
}
