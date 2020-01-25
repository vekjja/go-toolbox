package net

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

// DownloadFromURL : download image from provided url and save to provided filelocation
func DownloadFromURL(url, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

// SendRequest : send http request to provided url
func SendRequest(req *http.Request) ([]byte, error) {
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// GetLocalIP : get local ip address
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}

// GetPubIP : get public ip address
func GetPubIP() (string, error) {
	// we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	// https://ifconfig.co
	// https://ifconfig.me
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
