# Tehlug add event

This is a program for add event to Tehlug website


Data:
```
CURRENT_DATE: 2024-02-10T00:00:00+03:30  (Auto generate)
EVENT_DATE format: 1402-11-26 (User input)
START_TIME: 16:00   (User input)
END_TIME: 20:30   (User input)
SPEAKERS: is array of speakers name like it: پیام کمر زرین,محمود اسکندری,محمد فاضلی,سید حمید مهدوی,سعید رسولی,سینا جعفری     (User input)
REGISTER_LINK: is link of event     (User input)
TOPICS: is array of topics like it: Go, Python, Linux    (User input)
ADDRESS: is address of event like it: هفت تیر، خیابان قائم مقام، میدان شعاع، خیابان خدری، بعد از چهارراه سنایی،  پلاک ۳۱، سالن همایش‌های یاس    (User input)
LAT: is latitude of event like it: 35.720612    (User input)
LNG: is longitude of event like it: 51.418212    (User input)
LOCATION: is name of location like it: کدانو    (User input)
DESCRIPTION: is description of event like it:    (User input)
جلسه ۲۷۲ گروه کاربران گنو/لینوکس تهران و جلسه ۸۲ گروه کاربران پایتون تهران و جلسه گروه گوفرکانف (GopherConf) به صورت مشترک در روز پنج‌شنبه مورخ ۱۴۰۲/۱۱/۲۶ با حمایت کدانو برگزار می‌گردد.

این جلسه از ساعت ۱۶:۰۰ الی ۲۰:۳۰ برگزار می‌شود و ۹۰ دقیقه ابتدایی جلسه مربوط به جلسه گوفرکانف و ۹۰ دقیقه میانی مربوط به جلسه تهران‌پاگ و ۹۰ دقیقه انتهایی مربوط به جلسه تهران‌لاگ خواهد بود.

زمان‌بندی برنامه‌ها و ارائه‌ها:

از ساعت ۱۶:۰۰ الی ۱۶:۳۰ : گوفرکانف: ارائه «Concurrency in Go» توسط پیام کمرزرین
از ساعت ۱۶:۳۰ الی ۱۷:۰۰ : گوفرکانف: ارائه «گولنگ: چرا و چطور؟» توسط محمود اسکندری
از ساعت ۱۷:۰۰ الی ۱۷:۳۰ : پذیرایی و نتورکینگ
از ساعت ۱۷:۳۰ الی ۱۸:۱۰ : پاگ: ارائه «کُد برای کیش مات: ساخت هوش مصنوعی شطرنج باز با پایتون» توسط محمد فاضلی
از ساعت ۱۸:۱۰ الی ۱۸:۵۰ : پاگ: ارائه «استفاده از پایتون به عنوان یک مهندس نرم‌افزار» توسط «سید حمید مهدوی»
از ساعت ۱۸:۵۰ الی ۱۹:۱۰ : پذیرایی و نتورکینگ
از ساعت ۱۹:۱۰ الی ۱۹:۴۰ : لاگ: ارائه «Elastic Stack» توسط سعید رسولی
از ساعت ۱۹:۴۰ الی ۲۰:۰۰ : لاگ: ارائه «معرفی افزونه‌های کاربردی گنوم» توسط سینا جعفری
از ساعت ۲۰:۰۰ الی ۲۰:۳۰ : پذیرایی و نتورکینگ
```

Template:
```
---
title: \"{{ .EVENT_TITLE}}\"
date: {{ .CURRENT_DATE }}
layout: event
type: event
eventNumber: {{ .EVENT_NUMBER }}
eventDate: {{ .EVENT_DATE }}
startTime: {{ .START_TIME }}
endTime: {{ .END_TIME }}
address: "{{ .ADDRESS }}"
registerLink: {{ .REGISTER_LINK }}
speakers: [{{ .SPEAKERS }}]
topics: [{{ .TOPICS }}]
lat: {{ .LAT }}
lng: {{ .LNG }}
locations: {{ .LOCATION }}
---

{{ .DESCRIPTION }}
```


