# Goの基本的なとこ

## package ** について
Goではかならず一番上に`package **`の記述をする必要がある。

`**`の部分は自由に決めていいが、ディレクトリごとに統一するのが良い気がする

```
|
|- dir_1
|   |- file_1.go
|   |- file_2.go
|  
|- dir_2
    |- file_3.go
    |- file_4.go
```

上の例の場合、file_1とfile_2は`package module1`、file_3とfile_4は`package module2`にするイメージ

## 変数・定数について
```go
package main
import (
    "fmt"
)

func main(){
    // 変数
    var x int // varで変数名、型を宣言
    x = 1
    fmt.Println(x) // 1

    // 定数
    const y int = 8 // const 変数名 型 = * で宣言
    fmt.Println(y)
}
```

## データ型について軽く
```go
package main
import (
    "fmt"
)

func main(){
    var a int = 10
    var b string = "hello"
    var c float32 = 3.14
    var d bool = true

    a = c // コンパイルエラーとなる
    a = int(c) // キャストすることでOK -> 小数点以下切り捨てられて3になる
}
```

## 変数の初期化
```go
package main
import (
    "fmt"
)

func main(){
    var x int = 10
    var y = 10
    z := 10 // これよく使うらしい
}
```

## 制御構文(if, for, switch)など
```go
package main
import (
    "fmt"
)

func main(){

    // if文
    x := 10
    if x%10 == 0 {
        fmt.Println("10の倍数")
    } else if  x%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }


    // for文
    sum := 0
    for i := 0; i < 10; i++ {
        fmt.Println(i) // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
        sum += i
    }
    fmt.Println(sum) // 45

    for sum < 1000 {
        sum += sum
        fmt.Println(sum) // 90, 180, 360, 720, 1440
    }


    // switch文
    r := 2
    switch r {
        case 2:
            fmt.Println(i)
        case 3:
           fmt.Println(i)
        default:
           fmt.Println(i)
    }


    // 配列1
    list_1 := [6]int{1, 2, 3, 4, 5, 6} // [個数]型{要素, ...}
    fmt.Println(list_1[2]) // 3

    // 配列2.0
    list_2 := []int{1, 2, 3, 4, 5, 6} // []型{要素, ...}
    fmt.Println(len(list_2)) // 6 長さを取得

    // 配列2.1
    list_2 = append(list_2, 7, 8) // [1, 2...7, 8] 追加もできる

    // 配列スライス
    fmt.Println(list_2[1:3]) // 2, 3
    fmt.Println(list_2[:2]) // 1, 2, 3
    fmt.Println(list_2[4:]) // 5, 6, 7, 8

    // 配列 + for <-これよく使う
    list_3 := []int{1, 2, 3, 4, 5, 6}
    for i, v := range list_3 {
        fmt.Println(i, v) // 0 1, 1 2, 2 3, 3 4, 4 5, 5 6
    }

    for i := range list_3 {
        fmt.Println(list_3[i]) // 1, 2, 3, 4, 5, 6
    }


    // マップ
    scores := map[string]int{
        "佐藤": 100,
        "鈴木": 90,
        "伊藤": 80,
    }
    fmt.Println(scores["佐藤"]) // 100

    delete(scores, "佐藤") // {"鈴木": 90, "伊藤": 80,}
    score, exist := scores["佐藤"] // 存在するかどうかのboolを第二引数に取れる

    // マップ + for
    for k, v := range scores{
        fmt.Println(k, v) // 佐藤 100, 鈴木 90, 伊藤 80
    }

}
```