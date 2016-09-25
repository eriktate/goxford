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

	// TESTING PERSONS
	/*
		log.Println("Testing CreatePerson...")
		p, err := client.CreatePerson("testgroup", "Erik Tate", "")

		if err != nil {
			log.Println(err)
		}

		log.Printf("Created person with personID: %s", p.PersonID)
	*/

	log.Println("Testing GetPerson...")
	p, err := client.GetPerson("testGroup", "f4e60dcf-c32f-4609-ad1c-ffa8ea92cbf7")

	if err != nil {
		log.Println(err)
	}

	data, _ = json.Marshal(p)

	log.Println(string(data))
}
