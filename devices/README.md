# Devices
A golang server to manage devices.

> Part of [goHome](https://www.github.com/CabraKill/goHome/) project

![Device](/readme/device.png)
### Console
![Console](/readme/console.PNG)
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
| 1    | Desktop    | ![desktop](/readme/icons/desktop.png)           |
| 2    | SmartPhone | ![smartphone](/readme/icons/smartphone.png)     |
| 3    | IOT        | ![scatter_plot](/readme/icons/scatter_plot.png) |

### Action Types

| type | name  | icon                                      |
| ---- | ----- | ----------------------------------------- |
| 1    | OnOff | ![onoff](/readme/icons/onoff.png)         |
| 2    | On    | ![power_on](/readme/icons/power_on.png)   |
| 3    | Off   | ![power_off](/readme/icons/power_off.png) |
| 4    | Play  | ![play](/readme/icons/play.png)           |
| 5    | Pause | ![pause](/readme/icons/pause.png)         |
| 6    | Stop  | ![stop](/readme/icons/stop.png)           |





### Updates

Next incoming updates for this server.

* Concurrency: 
  * It is one of the most important features to have added in a server.