# Vkontakte API server client

[![Build Status](https://travis-ci.org/chapsuk/govk.svg)](https://travis-ci.org/chapsuk/govk)

Vk API version: `5.53`.

## API Methods

* [authorization](https://new.vk.com/dev/secure_how_to)
* [database.getCountries](https://new.vk.com/dev/database.getCountries)
* [database.getRegions](https://new.vk.com/dev/database.getRegions)
* [database.getCities](https://new.vk.com/dev/database.getCities)
* [database.getCitiesById](https://new.vk.com/dev/database.getCitiesById)
* [orders.get](https://new.vk.com/dev/orders.get)
* [users.isAppUser](https://new.vk.com/dev/users.isAppUser)

## Run

```
go install github.com/chapsuk/govk/cmd/govk
```

## Example

orders.get

```
 $ govk -c CLIENT_ID -s CLIENT_SECRET -cmd orders.get -count 20 -ogt 1
2016/04/28 14:58:15
Gotten access_token: ACCESS_TOKEN
2016/04/28 14:58:15
Result orders.get method
{ID:957030 AppOrderID:0 Status:charged UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:5 Date:1448969223}
{ID:957012 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_three#45189849 Amount:0 Date:1448967342}
{ID:957007 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:0 Date:1448967017}
{ID:957006 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_two#45189849 Amount:0 Date:1448966974}
{ID:957005 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:0 Date:1448966971}
{ID:948881 AppOrderID:0 Status:charged UserID:10872 ReceiverID:10872 Item:coins_30#9256941 Amount:3 Date:1447853021}
{ID:838963 AppOrderID:0 Status:charged UserID:87715133 ReceiverID:87715133 Item:premium_7#80706626 Amount:25 Date:1434620746}
{ID:838959 AppOrderID:0 Status:charged UserID:87715133 ReceiverID:87715133 Item:premium_7#80706626 Amount:25 Date:1434620662}
{ID:767480 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1426177226}
{ID:761985 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:like_1#149855 Amount:5 Date:1425558347}
{ID:761089 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1425470860}
{ID:760405 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1425394741}
{ID:378144 AppOrderID:0 Status:charged UserID:8182 ReceiverID:8182 Item:premium_1#4528709 Amount:10 Date:1386677299}
{ID:41867 AppOrderID:0 Status:charged UserID:5499135 ReceiverID:5499135 Item:premium_7d Amount:10 Date:1351602715}
```

database.getCountries

```
$ govk -cmd database.getCountries -count 20
2016/08/09 18:47:54
Without auth
2016/08/09 18:47:55
Result database.getCountries method
{ID:1 Title:Россия}
{ID:2 Title:Украина}
{ID:3 Title:Беларусь}
{ID:4 Title:Казахстан}
{ID:5 Title:Азербайджан}
{ID:6 Title:Армения}
{ID:7 Title:Грузия}
{ID:8 Title:Израиль}
{ID:9 Title:США}
{ID:65 Title:Германия}
{ID:11 Title:Кыргызстан}
{ID:12 Title:Латвия}
{ID:13 Title:Литва}
{ID:14 Title:Эстония}
{ID:15 Title:Молдова}
{ID:16 Title:Таджикистан}
{ID:17 Title:Туркменистан}
{ID:18 Title:Узбекистан}
```

database.getCities

```
$ govk -cmd database.getCities -count 10 -offset 0 -country 1
2016/08/09 19:34:55
Without auth
2016/08/09 19:34:56
Result database.getCities method
{ID:1 Title:Москва Important:1}
{ID:2 Title:Санкт-Петербург Important:1}
{ID:10 Title:Волгоград Important:0}
{ID:37 Title:Владивосток Important:0}
{ID:153 Title:Хабаровск Important:0}
{ID:49 Title:Екатеринбург Important:0}
{ID:60 Title:Казань Important:0}
{ID:61 Title:Калининград Important:0}
{ID:72 Title:Краснодар Important:0}
{ID:73 Title:Красноярск Important:0}
```

database.getCitiesById

```
$ go run cmd/govk/main.go  -cmd database.getCitiesById -lang 3 -city-ids "398,1,2"
2016/08/10 16:29:31
Without server auth
2016/08/10 16:29:31
Result database.getCitiesById method
{ID:398 Title:Alexandrov}
{ID:1 Title:Moscow}
{ID:2 Title:Saint Petersburg}
```

database.getRegions

```
$ govk -cmd database.getRegions -count 10 -offset 0 -country 1
2016/08/10 17:16:38
Without server auth
2016/08/10 17:16:38
Result database.getRegions method
{ID:1000001 Title:Адыгея}
{ID:1121540 Title:Алтай}
{ID:1121829 Title:Алтайский край}
{ID:1123488 Title:Амурская область}
{ID:1000236 Title:Архангельская область}
{ID:1004118 Title:Астраханская область}
{ID:1004565 Title:Башкортостан}
{ID:1009404 Title:Белгородская область}
{ID:1011109 Title:Брянская область}
{ID:1124157 Title:Бурятия}
```
