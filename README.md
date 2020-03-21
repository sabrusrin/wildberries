## Solutions for Wildberries practical tasks 
### Day 1  
#### Задание  
> Реализовать паттерн Фасад в соответствии с конвенцией.  
> Реализация паттерна должна быть следующая:  
> описать пакет facade в директории pkg/facade в директории cmd/facade создать файл main.go,  
> который фактически будет представлять собой интеграционное тестирование пакета.  

##### workspace structure
```.  
 
├── cmd  
│   └── facade 
│       └── main.go  
└── pkg  
    └── facade  
        ├── facade.go  
        ├── mentor  
        │   └── mentor.go  
        ├── practice  
        │   └── practice.go  
        └── theory  
            └── theory.go  
 ```  
