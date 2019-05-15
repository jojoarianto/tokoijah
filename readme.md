# TokoIjah

![](http://careers.salestock.io/images/logo-sorabel.png)

this is a case study which called "Toko Ijah" main domain of this case is Inventory. TokoIjah want to replace her spreadsheet by creating an application. so, goal of this project is to provide REST API for TokoIjah inventory application.

## Documentation
[Api Endpoint Documentation](https://documenter.getpostman.com/view/7437617/S1Lzwmf7?version=latest)

### Data Model
![TokoIjah Data Model](https://user-images.githubusercontent.com/5858756/57764381-87609c00-772d-11e9-86e5-90f32346528f.png)

### Architechture Design
![DDD Diagram](https://user-images.githubusercontent.com/5858756/57762643-0fdd3d80-772a-11e9-8fac-80ad6bf49b1e.png)

### quick start
clone this repository
```bash
git clone https://github.com/jojoarianto/tokoijah.git
```

### run to play
run download depedencies, compile & unit testing
```bash
bin/setup
```

run the app
```bash
bin/tokoijah
```

### run to development
```bash
make run
```

### migrate db schema
```bash
make migrate-schema
```

## libraries
project library : 
* github.com/julienschmidt/httprouter
* github.com/jinzhu/gorm
* gopkg.in/go-playground/validator.v9

