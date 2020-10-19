package main

// 不错的文章：http://www.360doc.com/content/20/0325/08/33093582_901513606.shtml

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const area = `{
		"武清staging->IDC": ["staging", "ksywq"],
			"IDC->C3金服专用-捷付区":["金服专用-捷付区", "c3"],
				"IDC->C3金服专用-NFC区":["金服专用-NFC区", "c3"],
					"IDC->C3金服专用-金融区":["金服专用-金融区", "c3"],
						"IDC->C4金服专用-捷付区":["金服专用-捷付区", "c4"],
							"IDC->C4金服专用-NFC区":["金服专用-NFC区", "c4"],
								"IDC->C4金服专用-金融区":["金服专用-金融区", "c4"]
									}`

const json = `{"name":["Janet","Prichard"],"age":47}`
const json1 = `{
	  "programmers": [
	      {
		            "firstName": "Janet", 
			          "lastName": "McLaughlin", 
				      }, {
					            "firstName": "Elliotte", 
						          "lastName": "Hunter", 
							      }, {
								            "firstName": "Jason", 
									          "lastName": "Harold", 
										      }
										        ]
										}`

func main() {
	value := gjson.Get(area, "IDC->C3金服专用-NFC区")
	fmt.Println("site:" + value.Array()[1].String())

	valuses := gjson.Get(json1, "programmers.0.lastName")
	fmt.Println(valuses)

}
