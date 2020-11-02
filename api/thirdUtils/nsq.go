package thirdUtils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/wangcong0918/sunrise/log"
	"os"
	"strings"
	"time"
)

func DedupeIoIDs(data []map[string]string, peopleID string) (result []string) {
	tempMap := make(map[string]string,0)
	for _,v := range data{
		tmpPeopleID := v["peopleID"]
		var tmpIoIDs []string
		if strings.ContainsAny(v["ioIDs"], "[") {
			json.Unmarshal([]byte(v["ioIDs"]), &tmpIoIDs)
		} else {
			tmpIoIDs = append(tmpIoIDs, v["ioIDs"])
		}
		if tmpPeopleID == peopleID {
			for index, val := range tmpIoIDs {
				log.Logger.Info(val)
				if tempMap[tmpIoIDs[index]]  == "" {
					result = append(result, val)
					tempMap[tmpIoIDs[index]] = val
				}
			}
		}
	}
	return result
}

func GerProjectCode(villageCode string) string {
	projectCodeStr := os.Getenv("PROJECT_CODE")
	projectCodeArr := strings.Split(projectCodeStr, ";")

	for _, v := range projectCodeArr {
		tmpVillageCode := v[0:14]
		if tmpVillageCode == villageCode {
			return strings.Split(v,"-")[1]
		}
	}

	return ""
}

func SendNsq(body []byte) (err error){
	tcpNsqdAddrr := os.Getenv("NSQ_CENTER_SERVER")
	//初始化配置
	config := nsq.NewConfig()
	//config.TlsConfig = newTLSConfig()
	config.LookupdPollInterval = 1 * time.Second // 重连时间
	config.AuthSecret = "lkMEzkefKFuilrJKsi5cbk9B69yRhnFh"
	//config.TlsV1 = true

	// 验证配置文件
	err = config.Validate()
	if err != nil {
		log.Logger.Error("config err --> ", err.Error())
		return err
	}

	tPro, err := nsq.NewProducer(tcpNsqdAddrr, config)
	if err != nil {
		log.Logger.Error("new producer error --> ", err.Error())
		return err
	}

	//发布消息
	err = tPro.Publish("client_proj", body)

	if err != nil {
		log.Logger.Error("nsq publish error --> ", err.Error())
		return err
	}

	return nil
}

func newTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("/Users/chenye/Desktop/chen/nsq/ca.crt", "/Users/chenye/Desktop/chen/nsq/ca.key")
	if err != nil {
		panic(err)
	}

	// Just to print out the client certificate..
	//cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	//if err != nil {
	//	panic(err)
	//}

	//log.Logger.Info(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}