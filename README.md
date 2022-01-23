# Squeat Cassandra REST API

- We are using a Revel server to handle our HTTP request
- We are communicating with our Cassandra database using GoCQL package

## Requirements 

#### In order to implement this API (if you have to)

- You need to have the "cassandraRest.tar.gz" archive
- Unzip the archive
   
# Revel

### Start the API:

-> Double click on the "run.sh" to launch the API

## Routes

GET     /user_all                                   
GET     /user_by_id                                 
GET     /user_by_mail                               
GET     /user_by_username                          
GET     /add_user                                   

GET     /meal_all                                   
GET     /meal_by_id                                 
GET     /meal_by_group                              
GET     /meal_by_name                               
GET     /meal_by_member                             
GET     /add_meal                                   

GET     /group_all                                   
GET     /group_by_id                                 
GET     /group_by_leader                             
GET     /group_by_member                             
GET     /add_group                                   
GET     /add_member_group                            


