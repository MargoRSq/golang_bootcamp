package main

/*
#include "cow.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"unsafe"
	// "github.com/kubernetes/utils"
)

type Candy struct {
	FullName string
	Price    int
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err == nil {
		if candy, ok := candies[order.CandyType]; ok {
			toPay := int32(candy.Price) * order.CandyCount
			switch money := order.Money; {
			case money >= toPay:
				w.WriteHeader(http.StatusCreated)
				cs := C.CString("Thank you")
				defer C.free(unsafe.Pointer(cs))
				fooRes := C.ask_cow(cs)
				result := C.GoString(fooRes)
				success := InlineResponse201{Change: money - toPay, Thanks: result}
				json.NewEncoder(w).Encode(success)
			case money < toPay:
				w.WriteHeader(http.StatusBadRequest)
				fail := InlineResponse400{Error_: fmt.Sprintf("You need %d more money!", toPay-money)}
				json.NewEncoder(w).Encode(fail)
			}
		} else {
			fail := InlineResponse400{Error_: "We don't have these candies!"}
			json.NewEncoder(w).Encode(fail)
		}
	} else {
		fail := InlineResponse400{Error_: "Invalid input!"}
		json.NewEncoder(w).Encode(fail)
	}
}
