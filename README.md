# GoDataStructure
Go Data Structure

> Go is not an object-oriented language

## Type & Struct
* 封裝性：引入初始化函式和方法（如 NewXXX 和 DoXXX）。
* 驗證性：透過初始化函式檢查屬性的有效性。
* 測試性：撰寫單元測試以確保功能的正確性。
* 擴展性：整理資料夾結構，方便日後擴充功能。
* Go doesn't support classes or inheritance in the complete sense

## Anonymous Structs
sometimes use anonymous structs when I know I won't ever need to use a struct again. For example, sometimes I'll use one to create the shape of some JSON data in HTTP handlers.

## Embedded Structs
### Embedded vs Nested
|特性	|Nested Structs	|Embedded Structs|
|----|----|----|
|字段訪問方式	|必須通過結構名稱訪問（如 p.Address.City）	|可以直接訪問內嵌結構的字段（如 p.City）|
|結構命名	|必須為嵌套結構的字段命名	|可以是匿名字段|
|字段重名處理	|不會衝突	|重名時需要通過結構名稱訪問|
|語義用途	|用於清晰表達層次結構	|用於結構體的字段重用|
|設計理念	|更加強調組織和分層	|更加強調簡化訪問和重用|

## Reciver
* Value Reciver
    *  值的複製：傳入方法的是接收器的副本，而非原始值。
    * 不修改原始值：在方法中對接收器進行的操作，不會影響外部的結構體。
* Pointer Reciver
    * 傳遞指標：方法接收的是接收器的指標，操作的是原始值，而非副本。
    *  可以修改原始值：方法內的修改會影響外部的結構體。

|特性	|Value Receiver	|Pointer Receiver|
|----|----|----|
|傳遞的內容	|值的複製（副本）	|指標（記憶體位址）|
|修改原始值的能力	|無法修改	|可以修改|
|記憶體效率	|開銷較大（結構體大時複製成本高）	|開銷較小（僅傳遞指標）|
|方法適用性	|只能在值接收器上調用	|可以在值或指標接收器上調用|