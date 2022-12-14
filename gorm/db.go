package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type tb_1 struct {
	// ID   int    `gorm:"primaryKey"`
	ID   int    `gorm:"primaryKey;autoIncrement:true;column:id"`
	Uid  int    `gorm:"column:uid"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
	Addr string `gorm:"column:addr"`
	// created_at time.Time
	// updated_at time.Time
	// deleted_at time.Time
	// `gorm:"type:varchar(255);"`
}
type marksheet struct {
	Student_id int `gorm:"foreignKey:id"`
	Marks      int `gorm:"column:marks"`
}

// CompanyRefer int
//   Company      Company `gorm:"foreignKey:CompanyRefer"`

// type users struct {
// 	Uid  int
// 	Name string
// 	Age  int
// 	Addr string
// }

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1213"
		dbname   = "db_1"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Db Connected...")
	}

	db.AutoMigrate(&tb_1{})
	// db.Create(&tb_1{Uid: 6, Name: "Nirmal", Age: 25, Addr: "Kolkata"})
	// db.Create(&tb_1{Uid: 9, Name: "Bolai", Age: 21, Addr: "Gopalgonj"})
	// db.Create(&tb_1{Uid: 66, Name: "Madhobda", Age: 22, Addr: "Kolkata"})
	// db.Create(&tb_1{Uid: 55, Name: "Bolai", Age: 27, Addr: "Jhargram"})

	fmt.Printf("DB %T \n", db)
	db.AutoMigrate(&marksheet{})

	// u := tb_1{Uid: 33, Name: "nimai", Age: 20, Addr: "Gokuldham"}
	// db.Create(&u)

	// db.First(&user)

	// er := db.Migrator().AddColumn(&tb_1{}, "uid")
	// if er != nil {
	// 	fmt.Println(er)
	// }

	// person := tb_1{name: "Jinzhu", age: 18, addr: "jgm"}

	// fmt.Println("res:", result)

	// result := db.Create(&person)
	// fmt.Println("res:", result)

	// Read
	// var p Data
	// db.First(&p, 1)             // find product with integer primary key
	// db.First(&p, "uid =?", "1") // find product with code D42

	// Update - update
	// db.Model(&p).Update("uid", 1)
	// // Update - update multiple fields
	// db.Model(&p).Updates(Data{uid: 2, name: "arijit"}) // non-zero fields
	// db.Model(&p).Updates(map[string]interface{}{"name": "laltu", "uid": 1})

	// // Delete - delete product
	// db.Delete(&p, 1)
	// var us tb_1
	// db.Where("ID = ?", 4).Find(&us)
	// us.Name = "Bisakha"
	// us.Age = 100
	// us.Uid = 24
	// us.Addr = "jhargram"
	// db.Save(&us)
	fmt.Println("")
	db.Model(&tb_1{}).Where("id = ?", 10).Update("name", "Rikisi")
	db.Model(&tb_1{}).Where("id = ?", 10).Update("age", 200)
	fmt.Println("Value Updated...")
	fmt.Println("")

	var dt tb_1
	er := db.Model(&tb_1{}).Where("id = ?", 9).Delete(dt)
	if er != nil {
		fmt.Println("Er on delete->", er)
	}

	var ud []tb_1
	db.Find(&ud)
	fmt.Println("DATAS FROM DATBASE ARE ......")
	for _, v := range ud {
		// if v.ID == 2 {
		// 	fmt.Printf("ID : %v, Uid : %v, Name : %v, Age : %v, Addr: %v \n", v.ID, v.Uid, v.Name, v.Age, v.Addr)
		// }
		fmt.Printf("ID : %v, Uid : %v, Name : %v, Age : %v, Addr: %v \n", v.ID, v.Uid, v.Name, v.Age, v.Addr)

	}
	// var dt tb_1
	// db.Model(&dt).Updates(tb_1{Uid: 5, Name: "Malay", Age: 18, Addr: "jhargram"})
	// db.Model(&dt).Select("Name", "Age").Where("id = ?", 4).Updates(tb_1{Name: "mitali", Age: 20})
	// fmt.Println("Updated..")

	// db.Where("name = ?", "Bisakha").Delete(&dt)

	// drop colum not tested yet
	//db.Model(&tb_1{}).Migrator().DropColumn(&tb_1{},"uid")
	// db.Table()
	ul := []tb_1{}
	db.Where(&tb_1{ID: 7}).Find(&ul)
	for _, v := range ul {
		fmt.Println("")
		fmt.Printf(" Data is %v  where id is %v\n", v, v.ID)
		fmt.Println("")
	}

	var e bool = true

	for e == true {
		fmt.Println("")
		fmt.Println("Enter 1 to serach with name..")

		fmt.Println("Enter 2 to insert into markssheet database.")

		fmt.Println("Enter 3 to Delete data from markssheet database. ")

		fmt.Println("Enter 4 to exit")

		fmt.Println("")

		var ch int

		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter Name to fetcch details from database...")
			var s string
			fmt.Scan(&s)

			var arr []string

			ul := []tb_1{}
			// Name: s
			db.Where(&tb_1{}).Find(&ul)
			for _, v := range ul {
				arr = append(arr, v.Name)
			}
			var flag bool = false

			for i := 0; i < len(arr); i++ {
				if arr[i] == s {
					flag = true
				}
			}
			if flag == true {

				dl := []tb_1{}
				// Name: s
				db.Where(&tb_1{Name: s}).Find(&dl)

				for _, v := range dl {
					fmt.Printf("ID: '%v' Uid: '%v' Name: '%v' Age: '%v' Addr: '%v' \n",
						v.ID, v.Uid, v.Name, v.Age, v.Addr)
				}

			} else {
				//fmt.Println("flag not true..")
				var arr2 []string
				var arr3 []int
				uh := []tb_1{}
				res := strings.ToLower(s)
				//fmt.Println("res:", res)

				db.Where("lower(name) LIKE ?  ", "%"+res+"%").Order("name").Find(&uh)
				for _, v := range uh {

					arr2 = append(arr2, v.Name)
					arr3 = append(arr3, v.ID)

				}
				if len(arr2) == 0 {
					fmt.Println("NO match found..")
				}

				// fmt.Println("arr2: ", arr2)
				for i := 0; i < len(arr2); i++ {
					fmt.Printf("Do You Mean %v with id %v ? \n", arr2[i], arr3[i])

					fmt.Printf("Enter enter y/n : ")
					var choice string
					fmt.Scan(&choice)
					// strings.ToLower(choice)
					if choice == "y" || choice == "Y" {

						dl := []tb_1{}
						// Name: s
						db.Where(&tb_1{Name: arr2[i], ID: arr3[i]}).Find(&dl)

						for _, v := range dl {
							fmt.Printf("ID: '%v' Uid: '%v' Name: '%v' Age: '%v' Addr: '%v' \n",
								v.ID, v.Uid, v.Name, v.Age, v.Addr)
						}

					}

				}

			}

			// }

			// var arr []string

			// //
			// db.Where("name LIKE ?", "%"+s+"%").Order("name").Find(&ul)
			// // db.Where(&tb_1{Name: s}).Find(&ul)
			// for _, v := range ul {

			// 	arr = append(arr, v.Name)

			// }
			// var cho string
			// fmt.Scan(&cho)
			// cho1 := strings.ToLower(cho)

			// var arrs []string
			// for i := 0; i < len(arr); i++ {
			// 	fmt.Printf("Do you mean '%v' ? (if yes Enter Y / if no enter N \n", arr[i])
			// 	arrs = append(arrs, arr[i])
			// }

			// if cho1 == "y" {
			// 	for _, v := range arrs {

			// 		db.Where(&tb_1{Name: v}).Find(&ul)
			// 		for _, v := range ul {

			// 	fmt.Printf("DETAIS ARE ID: %v Uid: %v Name: %v Age: %v Addr: %v \n",
			// 		v.ID, v.Uid, v.Name, v.Age, v.Addr)

			// }

			// }
			// }

		case 2:
			var a bool = true
			for a == true {
				fmt.Println("Enter Student Id :")
				var id int
				fmt.Scan(&id)
				fmt.Println("Enter Marks :")
				var mrk int
				fmt.Scan(&mrk)

				u := marksheet{Student_id: id, Marks: mrk}
				db.Create(&u)
				break

			}

		case 3:
			fmt.Println("enter id to be deleted...")
			var id1 int
			fmt.Scan(&id1)

			var dt marksheet
			db.Model(&marksheet{}).Where("student_id = ?", id1).Delete(dt)
			fmt.Println("deleted Student id", id1)

		case 4:
			e = false
		default:
			fmt.Println("Enter Write Choice...")

		}
	}
}
