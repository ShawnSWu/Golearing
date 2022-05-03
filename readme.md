
Golang學習筆記
===
![chat](https://img.shields.io/badge/github-GolangLearn-brightgreen.svg)

<!--ts-->
* [Golang學習筆記](#Golang學習筆記)
  * [GOROOT,GOPATH,GOMODULE差異](##GOROOT,GOPATH,GOMODULE差異)
    * [GO 1.11Version以前](#GO-1.11Version以前)
    * [GO 1.11Version以後](#GO-1.11Version以後)
  * [Go中較特別的資料型態差異](#Go中較特別的資料型態差異)
    * [Character](#Character)
  * [資料結構](#資料結構)
    * [Slice vs Array](#Slice-vs-Array)
    * [Slice Array兩者差異](#Slice-Array兩者差異)
    * [Slice示意圖](#Slice示意圖)
    * [Slice操作](#Slice操作)
    * [Map](#Map)
  * [Pointer指標](#Pointer指標)
  * [Structs](#Structs)
  * [Function](#Function)
    * [Call by Value or reference(pointer)](#Call-by-Value-or-reference(pointer))
    * [特殊情況](#特殊情況)
<!--te-->

## GOROOT,GOPATH,GOMODULE差異

* GOROOT：存放Go語言內建的程式庫的所在位置
* GOPATH：存放第三方套件的所在位置

### **GO 1.11Version以前**

當執行Golang程式碼：
先去GOROOT路徑下的src資料夾找，如果沒有就會去GOPATH路徑下的src找，如果還是沒有找到，就會錯誤。

> GOROOT就是放官方提供的標準庫，而GOPATH就是存放第三套件，所以預設將專案建立在GOPATH底下，就能引用到所有想引用的第三方套件，所以當執行 **go get** 來下載第三方套件時，就是存在GOPATH底下的src資料夾。

* **優點：統一、方便**

* **缺點：當專案很多時會很雜，甚至還會版本衝突**
  go get最常用在當我們想用別人公開在GitHub上的套件，可以幫我們從網路上clone到GOPATH/src裡面。雖然這樣很方便，但是你會發現GOPATH/src下的程式碼會很複雜，除了有你自己開發的專案資料夾，也包含其他第三方程式庫的專案資料夾。



### **GO 1.11Version以後**

出現了Go Modules這種方式來管理，將第三方程式庫儲存在本地的空間，並且給程式去引用，並且就是存在GOPATH裡面的pkg。
> Go Modules解決方式很像是Java看到Maven, Gradle的做法



---
## Go中較特別的資料型態差異

> 此處只紀錄與其他語言較有差異的地方

### Character
在Go中 沒有char型態可以宣告，Go使用 byte & rune 型態 來表達Character values，
Go的string其實也就是byte組成，
因為在電腦科學的角度來說，character value本質上就是整數型態，

例如： ASCII表中 97 = 'a', 100 = 'd'

範例code：
```
var ch rune = 'C'
fmt.Printf("%c", ch) // 打印出 C
 
var ch rune = 'a'
fmt.Printf("%d", ch) // 打印出 97
```

> 在Go中 儲存char的方式 是用byte & rune(整數型態)，差別只在 控制如何顯示它

---
## 資料結構
### Slice Array兩者差異

Slice 可以當成『在操作底層陣列』的『抽象』，透過操作這個『抽象』可以直接對底層真正對應的陣列去做操作，

而Slice有三個部分
1. **資料本身**
2. **容量（capacity）**
3. **長度(length)**

cap告訴你底層陣列總共有多少容量，而len告訴你陣列中現在有幾筆資料。

e.g.1
```
    //使用make()產生一個新的Slice，並設定底層Array總長度=5
    
    mySlice := make([]int, 0, 5) //產生一個len=0且cap=5的Slice
    fmt.Println(len(mySlice), cap(mySlice)) 
    
    //打印結果 =>  0 5
```

e.g. 2
```
    //產生新Slice，並設底層Array總長度=5且裡面放2筆資料(資料會給預設資料zero value)
    
    mySlice = make([]int, 2, 5)
fmt.Println(len(mySlice), cap(mySlice), mySlice) // 2 5
    
    //打印結果 =>  2 5 [0 0]
```
[zero value相關資訊](https://www.geeksforgeeks.org/zero-value-in-golang/)

### Slice Array兩者差異

> 1. Array宣告時要指定長度，Slice可以不用，
> 2. Array不行動態新增大小，所以若是今天需要存更多資料，我們必須自己處理把Array加大的相關算法，但如果透過Slice去操作的話，Go底層會幫你做這些事，所以當容量不夠時，底層會產生一個新的容量更大的Slice且把舊資料都複製過去，這也解釋了為什麼並不總是會產生新的Slice，更細節可以參閱[這裡](https://go.dev/blog/slices-intro)。

### Slice示意圖
![](https://i.imgur.com/FH720xB.jpg)


### Slice操作

```
    mySlice := make([]rune, 0, 5)

    //新增，透過append()
    mySlice = append(mySlice, 'A', 'K', 'G', 'L')
    printSlice(mySlice) // A,K,G,L

    //修改
    mySlice[0] = 'P'
    printSlice(mySlice) // P,K,G,L

    //*刪除 較特別
    toRemove := 2 //要刪除的元素index
    mySlice = append(mySlice[:toRemove], mySlice[toRemove+1:]...)
    printSlice(mySlice) // P,K,L
```
> *Slice沒有刪除元素的關鍵字，所以透過Slice取切片的方式，將不要的element排除掉後(如上述的code)，再組成新的Slice就完成了。
> 示意圖如下：
![](https://i.imgur.com/MgpqqsF.jpg)

<font size=3.5 color="#FF0000">注意！！
Slice在使用append()時的**可能會**產生新的Slice，並將舊Slice上的資料轉過去，即代表得到一個新的Slice，在使用時要注意這點。</font>


例子：

```
func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i = append(i, "4")
}

打印結果：
[1 2 3]

```

> 可以從結果看出，第0個值有被修改，但是新增的"4"卻沒被加入到Slice中，
> 這是因為，append可能會產生新的Slice
>
> Slice原文解釋：
> If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated.
>
>
>
> 導致其實"4"是被加到新的Slice當中，若此段程式想順利執行，解法就是直接return回去，如下所示

```
func main() {
    var s = []string{"1", "2", "3"}
    s = modifySlice2(s)
    fmt.Println(s)
}

func modifySlice2(i []string) []string {
    return append(i, "4")
}

打印結果：
[1 2 3 4]
```


### Map

> 在 Go 中，Map 也是 Key-Value pair 的組合，但是 Map 所有 Key 的資料型別都要一樣；所有 Value 的資料型別也要一樣。另外，在設值和取值需要使用 [] 而非 .。map 的 zero value 是 nil。

**基本宣告方式：**
![](https://i.imgur.com/u8liCgM.jpg)

更多例子
```
    // 可以直接帶值
    map := map[string]string{
	    "red":   "#ff0000",
	    "green": "#4bf745",
    }
    
    // 另種方式，使用 make 建立 Map。
    map := make(map[string]int)
```

#### Map操作
```
    //新增
    m := make(map[string]string)
    m["k1"] = "Shawn"
    m["k2"] = "Andy"
    fmt.Println(m) // map[k1:Shawn k2:Andy]

    //刪除
    delete(m, "k1")
    fmt.Println(m) // map[k2:Andy]
```


## Pointer指標

#### 指標概念：

程式語言中 有兩種變數：
* **變數 => 存放資料本身**
  宣告方式：
  `var x = 10`

* **指標變數 => 存放資料的『位址』**
  宣告方式：
```
    var x int = 10 //宣告變數
    var xPoint *int = &x //宣告指標變數，存放『位址』
    
    fmt.Println(x)
    fmt.Println(xPoint)
    fmt.Println(*xPoint)
    
    打印結果：
    10
    0x14000122008
    10
```
第一行宣告變數x，第二行宣告了指標變數xPoint，**由於是指標變數，所以只能存放『位址』，所以可以看到是存 &x 而不能存 x(&x 表示是取x這個資料的資料位址)**

較需要注意的是最後一行 打印出  *xPoint
在指標變數的前面加上星號＊，即表示<font size=3.5 color="#FF0000">反解指標</font>，即可透過指標指向的位置取得原始的資料

![](https://i.imgur.com/kl5nzZu.jpg)
<font size=2 color="#A9A9A9">image by [彭彭](https://www.youtube.com/watch?v=k_E9FCehyz4)</font>


## Structs

寫法類似於JavaScript object。

> 參數大小寫 決定此參數對外是public or protected, 非大寫開頭就只能在package內使用
```
type Person struct {
	Name string //對包外開放（public）
	age int // （protected）
}
```

待補...




## function 函數
Golang中的函數與其他程式語言相比並無太多特別之處，只需要記住語法即可

![](https://i.imgur.com/TbkacpF.jpg)
<font size=2 color="#A9A9A9">image by [GeeksforGeeks](https://www.geeksforgeeks.org/named-return-parameters-in-golang/)</font>

### Call by Value or reference(pointer)

> Everything in Go is passed by value.

<font size=3.5 color="#00BB00">在Golang中所有內容都是傳值的，因為Golang有指標，所以需要時call by reference時，可將指標變數傳進function中</font>

### 特殊情況

```
func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "9"
}

打印結果：
[9 2 3]
```
Go的slice其實也是一個Struct，此Struct包含了一個array指針以及len和cap兩個int類型的參數。

正如前面所說，Go是pass by value，將slice作為參數傳遞時，函數中確實創建一個slice參數的副本，這個副本同樣也包含array,len,cap這三個成員，
**只是副本中的array指針與原slice指向同一個地址，所以當修改副本slice的元素時，原slice的元素值也會被修改。（淺複製的既視感**
<font size=3 color="#00BB00">所以如果修改的是副本slice的len和cap時，原slice的len和cap仍會保持不變。</font>


<font size=4 color="#6F00D2">若想完全複製Slice到funcation中且不影響外部Slice方法</font>
> 在參數進到函數內後 先複製出一個全新的
```
func main() {
    var s = []string{"1", "2", "3"}
    modifySlice(s)
    fmt.Printf("原始的位址是：%p\n", &s)
}

func modifySlice(i []string) {
    copyNewOne := i[:len(i) - 1] //重新切片出一個
    copyNewOne = append(i, copyNewOne...) // 透過append複製出一個
    copyNewOne[0] = "9" // 此時修改的是新複製的Array複製出一個

    fmt.Printf("複製後的的位址是 %p\n", &copyNewOne)
}

打印結果：
複製後的的位址是 0x1400011e030
原始的位址是：0x1400011e018

```
