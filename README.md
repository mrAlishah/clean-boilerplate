# Clean-Boiler-Plate
CleanBoilerPlate


### features
environment variables <br>
fully dockerized <br>
advanced debugger<br>
sever reloader <br>

advanced test and mocking<br>
responses and validators<br>


### architecture and  rules
![clean architecture](https://github.com/mahdimehrabi/clean-boilerplate/blob/main/architecture.png?raw=true)

Errors log must handle in service<br>
Errors response must handle controller<br>
Never pass router framework instances like gin.Content to services and repositories<br>
Its so important that order be like Controller -> Service -> Repository<br>
controller => router framework stuff like validation , response<br>
