package myoop

//import ("fmt")

//type User struct {
//	id int
//	name string
//}
//type Manager struct {
//	User
//	title string
//}
//
//func (xxx *User) ToString() string{
//	return fmt.Sprintf("User: %p, %v", xxx, xxx)
//}
//func (xxx *Data) PointerTest() {
//	// func PointerTest(self *Data);
//	fmt.Printf("Pointer: %p\n", xxx)
//}

func Myoop()  {
	//m:=Manager{User{1,"Tom"},"Adminisrator"}
	//var u User = m.User
	//fmt.Println()
	//d := Data{}
	//p := &d
	//fmt.Printf("Data: %p\n", p)
	//d.ValueTest() // ValueTest(d)
	//d.PointerTest() // PointerTest(&d)
	//p.ValueTest() // ValueTest(*p)
	//p.PointerTest() // PointerTest(p)
	//Pointer指向的是原对象
	//从 1.4 开始,不再支支持多级指针查找方方法成员。
	//通过匿名字段,可获得和继承类似的复用用能力力。依据编译器查找次序,只需在外层定义同
	//名方方法,就可以实现 "override"。
	/*
	• 类型 T 方方法集包含全部 receiver T 方方法。
	• 类型 *T 方方法集包含全部 receiver T + *T 方方法。
	• 如类型 S 包含匿名字段 T,则 S 方方法集包含 T 方方法。
	• 如类型 S 包含匿名字段 *T,则 S 方方法集包含 T + *T 方方法。
	• 不管嵌入入 T 或 *T,*S 方方法集总是包含 T + *T 方方法。
	*/


	//u := User{1, "Tom"}
	//u.Test()
	//mValue := u.Test
	//mValue()
	//// 隐式传递 receiver
	//mExpression := (*User).Test
	//mExpression(&u)
	//显式传递receiver
	//u:=User{1,"Tom"}
	//mValue := u.Test
	// 立即复制 receiver,因为不是指针类型,不受后续修改影响。
	//u.id,u.name = 2,"Jack"
	//u.Test()
	//进行了复制
	//mValue()

	//u := User{1, "Tom"}
	//fmt.Printf("User: %p, %v\n", &u, u)
	//mv := User.TestValue
	//mv(u)
	//mp := (*User).TestPointer
	//mp(&u)
	//mp2 := (*User).TestValue // *User 方方法集包含 TestValue。
	//mp2(&u) // 签名变为 func TestValue(self *User)。




	}

//type User struct {
//	id int
//	name string
//	}
//func (self *User) TestPointer() {
//	fmt.Printf("TestPointer: %p, %v\n", self, self)
//}
//func (self User) TestValue() {
//	fmt.Printf("TestValue: %p, %v\n", &self, self)
//}
//func (self User)Test()  {
//	fmt.Println(self)
//}
	//func(self *User) Test(){
	//	fmt.Printf("%p, %v\n", self, self)
	//}