# dts-task-planner
Task Planner app built with go fullstack web application using echo, tailwindcss, and mysql gorm.

## Getting started
- clone repository ini
```
https://github.com/wildanie12/dts-task-planner.git
```
### ENV Setup
- Copy file `.env.example`, lalu rename menjadi `.env`
- Siapkan database mysql baru
atur konfigurasi pada file `.env`
```
APP_HOST=localhost
APP_PORT=8000
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=root
DB_NAME=dts_task_planner
```

### Run the Application

**Development mode**
- Lakukan run
```
go run .
```


**Build from source**
- Lakukan build
```
go build -o ./app
```
- Jalankan app
```
./app
```

### Screenshot
![image](https://user-images.githubusercontent.com/13761315/184459222-06c413b8-2b25-4492-8731-d700d75871aa.png)


### Contributor

> - Badar Wildani | [Github](https://github.com/wildanie12)
