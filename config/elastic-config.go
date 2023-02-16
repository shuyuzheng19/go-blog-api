package config

import "github.com/olivere/elastic/v7"

var ES *elastic.Client

func connectElastic() *elastic.Client {

	client, err := elastic.NewClient(

		elastic.SetSniff(false),

		elastic.SetURL("http://"+HostName+":9200"),

		elastic.SetBasicAuth("elastic", "xiaoyu2528959216"),

		elastic.SetGzip(true),
	)

	if err != nil {
		panic(err)
	}

	return client
}

func InitElastic() {
	ES = connectElastic()
}
