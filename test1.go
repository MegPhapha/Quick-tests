/*Quick test create a struct name Admission with the following attributes 
- Fullname
- Age
- Created at
- updated at
- High school name
- date of birth 
- place of birth
- Hometown 

1. Create an array of the struct with dummy values and print them all out.

2. Loop through the structs and print the students information such that the output would be:

Fullname —— Age —- High school name
************
Fullname —— Age —- High school name

3. Add the total ages in your array and display the answer. */


package main
import ("fmt"
"time"
)

type Admission struct {
	Fullname           string
	Age                int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	HighSchoolName     string
	DateOfBirth        string
	PlaceOfBirth       string
	Hometown           string
}

// 1. Array of structs with dummy values and print them all out //
func main() {
	var admissions = []Admission { 
		{ 
	Fullname:   "Margaret Amanfu", 
	Age:        24, 
	CreatedAt:   time.Now() ,
	UpdatedAt:  time.Now() ,
	HighSchoolName: "OLA High School" ,  
	DateOfBirth:   "31/05/1999" ,    
	PlaceOfBirth:  "Kumasi" ,    
	Hometown: "Ho",
} , 
{
	Fullname:   "Meg Phapha", 
	Age:        20, 
	CreatedAt:   time.Now() ,
	UpdatedAt:  time.Now() ,
	HighSchoolName: "Wesley Girls High School" ,  
	DateOfBirth:   "01/12/2003" ,    
	PlaceOfBirth:  "Accra" ,    
	Hometown: "Sunyani",
} ,
		}

	// print all structs
	//for _, admissions := range admissions { 
		// fmt.Println("Fullname:", admissions.Fullname)
		//  fmt.Println("Age:" , admissions.Age)
		 //fmt.Println("Created At:" ,admissions.CreatedAt)
		 //fmt.Println("Updated At:" , admissions.UpdatedAt)
		 //fmt.Println("High School Name:", admissions.HighSchoolName)
		 //fmt.Println("Date Of Birth:" ,admissions.DateOfBirth)
		// fmt.Println("Place Of Birt:" , admissions.PlaceOfBirth)
		 //fmt.Println("Hometown:" , admissions.Hometown)
	



		 // 2. loop through the struct and print the students information
		 for _, admissions:=range admissions {
			fmt.Println(
			"Fullname——" , admissions.Fullname,
			"Age—-" , admissions.Age,
			"High School Name:" ,admissions.HighSchoolName, )
			fmt.Println("***************")

			}
		 

// 3. Add total ages in array and display answer
		 totalAges := 0
for _, admission := range admissions {
	totalAges += admission.Age
}
	fmt.Println("Total ages:",totalAges)

} 
		
