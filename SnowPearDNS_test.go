package main

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_byteString(t *testing.T) {
	Convey("byetString_test", t, func() {
		var retbyte []byte = []byte("8.8.8.8;3.3.3.3")
		t.Logf("start test bytes")
		fmt.Printf(byteString(retbyte))
		//t.Log(byteString(retbyte))
		So(byteString(retbyte), ShouldNotBeBlank)
	})
}

func Test_cx(t *testing.T) {

	var testres string = "{\"Status\":0,\"TC\":false,\"RD\":true,\"RA\":true,\"AD\":false,\"CD\":false,\"Question\":[{\"name\":\"github.com.\",\"type\":1}],\"Answer\":[{\"name\":\"github.com.\",\"type\":1,\"TTL\":180,\"Expires\":\"Tue, 14 Dec 2021 23:45:41 UTC\",\"data\":\"20.205.243.166\"}],\"edns_client_subnet\":\"116.227.244.37/22\"}"
	var cnameres string = "{\"Status\":0,\"TC\":false,\"RD\":true,\"RA\":true,\"AD\":false,\"CD\":false,\"Question\":[{\"name\":\"www.baidu.com.\",\"type\":5}],\"Answer\":[{\"name\":\"www.baidu.com.\",\"type\":5,\"TTL\":1200,\"Expires\":\"Wed, 15 Dec 2021 00:18:23 UTC\",\"data\":\"www.a.shifen.com.\"}],\"edns_client_subnet\":\"116.227.244.37/24\"}"
	t.Log(testres)
	t.Log(cnameres)
	var ans DOH_Response
	json.Unmarshal([]byte(testres), &ans)
	fmt.Println("here we go")
	fmt.Println(ans.Status)
	fmt.Println(string(ans.Answer[0].Data))
	PrettyPrint(ans)
	//fmt.Printf("%+v\n", ans)

	/*in := []int{2, 5, 6}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	fmt.Println(pick)*/
	//https://go.dev/play/p/bVovbZHNGRQ.go?download=true
	/*package main

	import (
		"encoding/json"
		"fmt"
	)

	// The same json tags will be used to encode data into JSON
	type Bird struct {
		Species     string `json:"birdType"`
		Description string `json:"what it does"`
	}

	func main() {
		pigeon := &Bird{
			Species:     "Pigeon",
			Description: "likes to eat seed",
		}

		// we can use the json.Marhal function to
		// encode the pigeon variable to a JSON string
		data, _ := json.Marshal(pigeon)
		// data is the JSON string represented as bytes
		// the second parameter here is the error, which we
		// are ignoring for now, but which you should ideally handle
		// in production grade code

		// to print the data, we can typecast it to a string
		fmt.Println(string(data))
	}*/
}
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func Test_get_a(t *testing.T) {
	Convey("get_a_test", t, func() {
		var domain string = "www.github.com"
		t.Logf("start test bytes")
		//fmt.Printf(string(get_a(domain)))
		for _, vl := range get_a(domain) {
			fmt.Println(vl)
		}
		fmt.Println("get_a")
		//t.Log(byteString(retbyte))
		So(get_a(domain)[0], ShouldNotBeBlank)
	})
}
func Test_get_cname(t *testing.T) {
	Convey("get_cname_test", t, func() {
		var domain string = "www.baidu.com"
		t.Logf("start test bytes")
		//fmt.Printf(string(get_a(domain)))
		for _, vl := range get_cname(domain) {
			fmt.Println(vl)
		}
		fmt.Println("get_cname")
		//t.Log(byteString(retbyte))
		So(get_cname(domain)[0], ShouldNotBeBlank)
	})
}
