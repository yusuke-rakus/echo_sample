# ポインタについて

## 参照渡し
```go
package main

import "fmt"

func main() {
    name := "Liza"

    namePoint := &name // ポインタ変数を渡す

    fmt.Println(namePoint)  // 0xc000014070
    fmt.Println(*namePoint) // Liza

    *namePoint = "Emma"
	/*
    namePoint = "Emma" // 'namePoint'はポインタ変数(１６進数のアドレス)のため、そこに文字列"Emma"を入れようとしてエラー

    namePoint := name  // 値渡し
    namePoint = "Emma" // からの値変更ならOK
	*/

    fmt.Println(name)       // Emma
    fmt.Println(namePoint)  // Emma
}
```


## 構造体とポインタ
```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	Liza := Person{"Liza", 18}

    Emma := Liza      // 値渡し(deep copy)
    Emma.name = "Emma"
	fmt.Println(Liza) // {Liza, 18}
	fmt.Println(Emma) // {Emma, 18}

	Maria := &Liza // 参照渡し(shallow copy)
	Maria.name = "Maria" // (*Maria).name でもOK
	fmt.Println(Liza)    // {Maria, 18}
	fmt.Println(*Maria)  // {Maria, 18}
	// fmt.Println(Maria) の場合、 &{Maria, 18} となる
}
```

## ポインタのまとめ
ポインタとは、むやみやたらに変数を作りまくらないための機能

`&変数名A`でポインタを渡し、`*変数名B`でポインタを変数として受け取っている。

つまり、`*変数名`となっているものを辿っていくと`&変数名`となっている
