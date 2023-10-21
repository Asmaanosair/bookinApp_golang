package main
import (
	"fmt"
	"strings"
)

func main(){
	for {
	var fristName string
	var lastName  string
	var email     string
	var phone     string
	var remainingTicket int
	var conferancetecket int
	var bookings  []string
	 fmt.Println("please Enter your fristName  ")
	 fmt.Scan(&fristName)
	 fmt.Println("please Enter your lastName  ")
	 fmt.Scan(&lastName)
	 fmt.Println("please Enter your email  ")
	 fmt.Scan(&email)
	 fmt.Println("please Enter your phone  ")
	 fmt.Scan(&phone)
	 fmt.Println("pleaseinter your conferancetecket")
     fmt.Scan(&conferancetecket)
	 remainingTicket=50
	 if remainingTicket <= conferancetecket  {
		 fmt.Printf("you cant book \n") 
		 break
	 }
	 remainingTicket=remainingTicket-conferancetecket

	 fmt.Printf("thank you for your  fristName : %v & \n lastName : %v  \n email : %v \n  Phone : %v \n" ,fristName, lastName ,email ,phone)
     fmt.Printf(" your conferanceticket = %v" , remainingTicket)
      bookings=append(bookings,fristName + " " + lastName)
	  bookings=append(bookings,fristName + " " + lastName)
	  fristNames :=[]string{}
	  for _,booking := range bookings{
		var names = strings.Fields(booking)
           fristNames = append(fristNames,names[0])
	  } 
	    fmt.Printf("booking fristNames  %v \n",fristNames)
	  fmt.Printf("booking Type %T \n",bookings)
	  fmt.Printf("booking Value %v \n",bookings)
	  fmt.Printf("booking Value %v \n",bookings[0])
	  fmt.Printf("booking length %v \n",len(bookings))
	 fmt.Printf(" your conferanceticket = %v \n" , remainingTicket)
	}
	}