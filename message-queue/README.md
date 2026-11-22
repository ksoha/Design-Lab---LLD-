#Building a RabbitMQ style message broker 


* Message : Represents a message and its rounting key.
* Queue : Stores messages until consumed.
* Exchange : Routes message to queue using binding. 
* Broker : Manages message and queues. 
* Producer and Consumer : publish and consume messages. 



#Architecture : 

Order Service
     |
     v
   Broker
     |
   Exchange
     |
     |---- email_queue ----> Email Worker
     |
     |---- sms_queue ------> SMS Worker
     |
     |---- inventory_queue -> Inventory Worker
