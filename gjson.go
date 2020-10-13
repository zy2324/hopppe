package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
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
	value := gjson.Get(json, "name.last")
	println(value.String())

	result := gjson.Get(json1, "programmers.#.lastName")
	for _, name := range result.Array() {
		println(name.String())
	}
}
