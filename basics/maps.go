// Maps are unordered collection of Key Value Pairs
package basics

import (
	"fmt"
	"maps"
)

func main() {

	//Ways to Create Map
	// var m map[keyType]valueType

	// makeVariable = make(map[keyType]valueType)

	//using a map literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key"]) // In this case the key did not exist so it resulted in a 0, if it would have been a string the result would be a blank string ""
	
	delete(myMap, "key1") // delete a key-value pair in a map
	fmt.Println(myMap)	
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, unknownValue := myMap["key1"] // unknownValue is a boolean which tells whether the key exists or not.
	// fmt.Println(value)
	fmt.Println("Is any value associated with 'key1':", unknownValue)
  
	myMap2 := map[string]int{"a":1, "b":2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a":1, "b":2}
	
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	// Zero Value of the map
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil.")
	} else {
		fmt.Println("The map is not initialized to nil.")
	}

	val := myMap4["Key"]
	fmt.Println(val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4) //This does not work when initializing map like myMap4 it contains nil, to make key value pairs in this map you need to use make 

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)
	fmt.Println("myMap4 length is", len(myMap))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)


}
// By Convention, we need to use "ok" variable to check whether the key exists in the map or not.
