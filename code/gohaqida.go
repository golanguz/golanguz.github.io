/*
Avvalan bor Go dasturlash tili haqida bir oz ma'lumot
Go dasturlash tili Google korporatsiyasida 2009 yil yaratilgan 
uni yaratgan insonlar Robert griesemer, Rob pike, Ken thompson.
https://en.wikipedia.org/wiki/Go_(programming_language)
nima sababdan Google bu dasturlash tilini yaratdi? chunki, ular allaqachon 
ischlatiyotgan dasturlash tillari C++, Java, Python, va boshqa". Google 
korporatsiyasiga kerak bolgan performans bera ololmaganligi uchun yaratilga 
xozirgi zamonda Google'ning 90foyizga yaqin backend'i Golang dasturlash 
tilini ishlatadi, agar bilsangiz Google.com saytiga kuniga 1'miliard'dan ko'p 
inson tashrif etadi, Va bullarning hammasi Golang yordami bilan eplanadi.

Go bu statik-yozilim va garbage-collected dasturlash tili u ko'p
tillardan ALGOL60, CSP, Squeak, Newsqueak, Alef, Pascal, Modula-2, Oberon,
Object Oberon, Oberon2, Va ko'pincha C dan  g'oya olgan. Statik-yozilim,
Bu degani siz kod yoziyotganingizda Go To'plamchisi(Compiler) xatolaringizni
tekshiradi misol atom editorida go-plus paketini o'rnatgan bo'lsangiz
kod yoziyotganingizda var ni o'rniga car deb yozsangiz u sizga car so'zini
tagida qizli chiziq yozadi(xato degani),
Go to'plamchisi xam xuddi shunday ishlaydi.
garbage-collceted(axlat-teruvchi)? Sistemingizdagi xotiralarni Go programmasi 
Qullanmiyotgan bo'lsa Go to'plamchisi o'sha xotiralarni Bo'shatadi,
Qanaqa xotira? Misol birinchi sahifamizda ko'rganingizdek ("salom gopherbek")
programmasini yozdik, uning xotirasi (mening sitemimda)"0xc82000a310" adresida, 
siszning sistemingizda farqli bo'lishi mumkin. Ammo katta proyektlarda
biz ("salom gopherbek") yozmaymiz, Bizda yuzlab o'zgaruvchilar,funksiyalar
va structlar bo'ladi, Va ularning xammasi o'zlariga sistemingizda joy aniqlashadi.
Misol sizda mobil o'yin dasturi bor, va qullanuvchi u o'yinni o'yniyapti
farqli tugmalarni bosiyapti, lekin sizning dasturingizsa tugmalar juda xam ko'p
ammo qullanchi faqat 2ta tugmani bosiyapti boshqa tugmalar ishlatilmasdan turipti
lekin mobil telefonida xotira sarfliyapti, shu yerda (axlat-teruvchi) ishga
tushadi bosilmiyotgan tugmalarni xotirasini bo'shatadi(toki qullanchi ularni
bosguncha) va shunday qilib sizning programmingiz yaxshi va tez ishlaydi.
Buni (garbage-collection "axlat-teruvchi") deydilar.
*/
