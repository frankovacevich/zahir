package server

import (
	"fmt"
	"os"
	"testing"
	"zahir/data"
	"zahir/utils"
)

func setUp() func() {
	err := data.LoadAs("../fixtures.json", "temp.json")
	if err != nil {
		panic(err)
	}

	if Router == nil {
		createRouter()
	}

	return func() {
		os.Remove("temp.json")
	}
}

func TestGetSources(t *testing.T) {
	defer setUp()()

	r := utils.MakeGetRequest("/v1/sources", Router, t)
	utils.AssertIntEqual(200, r.Code, t)

	body := utils.GetResponseBodyAsJsonList(r, t)
	utils.AssertIntEqual(3, len(body), t)
}

func TestCreateSource(t *testing.T) {
	defer setUp()()

	sources, err := data.GetSources()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(3, len(sources), t)

	r := utils.MakePostRequest("/v1/sources", Router, t)
	utils.AssertIntEqual(200, r.Code, t)

	sources, err = data.GetSources()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(4, len(sources), t)
	fmt.Println(sources)
}

func TestGetSequences(t *testing.T) {
	defer setUp()()

	r := utils.MakeGetRequest("/v1/sequences", Router, t)
	utils.AssertIntEqual(200, r.Code, t)

	body := utils.GetResponseBodyAsJsonList(r, t)
	utils.AssertIntEqual(3, len(body), t)
}

func TestGetSequenceData(t *testing.T) {
	defer setUp()()

	r := utils.MakeGetRequest("/v1/sequences/1", Router, t)
	utils.AssertIntEqual(200, r.Code, t)

	body := utils.GetResponseBodyAsJsonMap(r, t)
	utils.AssertIntEqual(2, len(body), t)

	sources := body["sources"].([]interface{})
	utils.AssertIntEqual(2, len(sources), t)
}
