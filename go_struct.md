# Goの構造体など

## 構造体の定義、インスタンス化
```go
package main

// 構造体を定義
type Person struct {
    Name string // 先頭が大文字の場合、外部からアクセスできる
    age int // 先頭が小文字の場合、外部からアクセスできない
}

func main() {
    // Person型のmikeを作成(1)
    var mike Person
    mike.Name = "Mike"
    mike.age = 20

    // Person型のlizaを作成(2)
    liza := Person{"liza", 18}

    // Person型のrobertを作成(3)
	robert := Person{Name: "robert", age: 28}
}
```

## クラスメソッド的な
```go
package main

import "fmt"

type Person struct {
	Name string
	age  int
}

// func (*レシーバ引数*) *関数名*(*引数*) *戻り値* {
//     ***処理***
// }
func (p Person) set_name(name string) Person {
    p.Name = name
	return p
}

func (p Person) set_age(age int) Person {
    p.age = age
	return p
}

func (p Person) introduction() {
    fmt.Println("My name is ", p.Name, "I'm ", p.age, "yars old")
}

func main() {
	var liza Person
	liza.set_name("liza").set_age(20)
	liza.introduction()
}
```

## 継承的な
```go
package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) introduction() string {
	return p.Name + string(p.age)
}

// Personを継承したUserを作成
type User struct {
	Person
	user_id string
}


func main() {
	liza := Person{Name: "liza", age: 20}

    // {Person: {Name: "liza", age: 20}, user_id: "u111"}が生成できる
	liza_account := User{user_id: "u111", Person: liza}

    // Personで定義したメソッドをUserでも利用可能
	liza_account.introduction()
}
```


## ポインタ(* &)
```go
package main

import "fmt"

func main(){
	var var1 int = 10
	var2 var2 *int = &var1

	fmt.Println(var1)  // 10
	fmt.Println(var2)  // 0x10414020
	fmt.Println(*var2) // 10
}
```
