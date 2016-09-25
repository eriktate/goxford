package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/eriktate/goxford"
)

func main() {
	client := &goxford.Client{}

	client.InitFace("92b79212e6784feda169c146233a4545")

	log.Println("Testing DetectURL...")
	res, err := client.DetectURL("http://images.amcnetworks.com/bbcamerica.com/wp-content/uploads/2015/01/tenthdoctor.jpg", "age,gender,headPose,smile,facialHair,glasses", true, true)

	if err != nil {
		log.Println(err)
	}

	data, _ := json.Marshal(res)

	fmt.Println(string(data))

	log.Println("Testing DetectPath (and DetectBytes)...")
	res, err = client.DetectPath("./testImage.jpg", "age,gender,headPose,smile,facialHair,glasses", true, true)

	if err != nil {
		log.Println(err)
	}

	data, _ = json.Marshal(res)

	fmt.Println(string(data))

	// TESTING PERSON GROUPS
	/*log.Println("Testing CreatePersonGroup...")
	_, err = client.CreatePersonGroup("testGroup", "Group for testing", "")

	if err != nil {
		log.Println(err)
	}

	log.Println("Done!")
	*/
	log.Println("Testing GetPersonGroup...")
	pg, err := client.GetPersonGroup("testGroup")

	if err != nil {
		log.Println(err)
	}

	data, _ = json.Marshal(pg)

	log.Println(string(data))
}
