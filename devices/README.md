# Devices
A golang server to manage devices.

> Part of [goHome](https://www.github.com/CabraKill/goHome/) project

![Device](/devices/readme/device.png)
### Console
![Console](/devices/readme/console.PNG)
### Features

Like a REST API, it has operations like GET, POST and DELETE using the corresponding  HTTP method.

* GET: Returns the list of devices.

* POST: Adds a device to the database. 

	* > If a device with same Ip and Name existings the operation overrides it.

* DELETE: Deletes a device with same Ip and Name if exists.

### Infrastruct
```
📦devices
 ┣ 📂driver
 ┣ 📂model
 ┣ 📂readme
 ┣ 📂repository
 ┣ 📜devices.go
```

### Database

The current need of the database is managing the devices while the server is operating. Since the data infrastruct is horizontal with all its properties, a json list is implemented in memory to work as the database.



### Device Types

| type | name       | icon                                            |
| ---- | ---------- | ----------------------------------------------- |
| 0    | Desktop    | ![desktop](/devices/readme/icons/desktop.png)           |
| 1    | SmartPhone | ![smartphone](/devices/readme/icons/smartphone.png)     |
| 2    | IOT        | ![scatter_plot](/devices/readme/icons/scatter_plot.png) |

### Action Types

| type | name  | icon                                      |
| ---- | ----- | ----------------------------------------- |
| 0    | OnOff | ![onoff](/devices/readme/icons/onoff.png)         |
| 1    | On    | ![power_on](/devices/readme/icons/power_on.png)   |
| 2    | Off   | ![power_off](/devices/readme/icons/power_off.png) |
| 3    | Play  | ![play](/devices/readme/icons/play.png)           |
| 4    | Pause | ![pause](/devices/readme/icons/pause.png)         |
| 5    | Stop  | ![stop](/devices/readme/icons/stop.png)           |





### Updates

Next incoming updates for this server.

* Concurrency: 
  * It is one of the most important features to have added in a server.