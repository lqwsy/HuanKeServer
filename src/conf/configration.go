package configration

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"os"
)

type xmlGlobal struct {
	Server xmlServer `xml:"server"`
	Mysql  xmlMysql  `xml:"mysql"`
}

type xmlServer struct {
	Bindaddr string `xml:"bindaddr"`
	Hsot     string `xml:"hsot"`
}

type xmlMysql struct {
	AdminDsn    string `xml:"admin_dsn"`
	Maxidleconn int    `xml:"maxidleconn"`
}

var Global xmlGlobal

func InitData(filename string) bool {
	_, err := LoadXmlConfig(filename, &Global)
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return false
	}
	return true
}

func LoadXmlConfig(filename string, xmlStruct interface{}) (contents []byte, err error) {
	fd, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("LoadConfig: Error: Counld not open %q for reading: %s\n ", filename, err)
		return
	}
	defer fd.Close()

	contents, err = ioutil.ReadAll(fd)
	if err != nil {
		err = fmt.Errorf("LoadConfig: Error: Could not open %q: %s \n", filename, err)
		return
	}

	if err = xml.Unmarshal(contents, xmlStruct); err != nil {
		err = fmt.Errorf("LoadConfig: Error: Could not parse XML configuration in %q: %s\n", filename, err)
		return
	}
	err = nil
	return
}
