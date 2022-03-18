# this is a sample project based on Go for getting URLs and printing md5 of the response 


## Introduction

This Application perform the requests in parallel(it accepts a flag to indicate this limit and the default is 10), and the URLs to download the URL and calculate Md5 of the result on the screen as it was defined in the task.
It also can get urls through a restful Api whit Post method on path /md5urls and data as json formate such as {"urls":["","",...]}

* ### *Description:*
    * This project stand-alone backend-server and just print the results on the screen and respond Ok.

## Getting Started
Clone the git repository in your system and then cd into the project root directory

```bash
$ git clone https://github.com/mitrajaafari60/urlsMd5
$ cd urlsMd5
```
### Project Structure
It used some related and needed parts of one of the best practice Go-Project skeleton (which is recommended by GitHub is https://github.com/golang-standards/project-layout)
* Project skeleton is as below:
    * cmd
        * the main part of the project for setting a flag and getting urls and HTTP server is launched in this package.
    * internal  
        * controller
            * <sub>functions in  this folder controls the request formate and in cas of success, service is called and the appropriate response is generated and will be sent back to the client </sub>
        * DataModel
            * <sub>In this folder,data model for restApi request included.</sub> 
        * service
            * <sub>all related parts of the service and their unit tests are developed in service package.</sub>
    * pkg 
        * flagArgs
            * <sub>general functions which developed and used in the project is categorized in pkg folder</sub>

### how to lunch
  for easy use of the project for those who don't have a Go compiler in their system 
  they can just run simply as below, also lunch http server on port 8000 will be lunched.
  for running by go it can be run by main which is located in cmd folder with go run main.go ...
```bash
  $> ./myhttp http://google.com
  $> ./myhttp -parallel=5 google.com
  $> ./myhttp -parallel 3 google.com facebook.com yahoo.com yandex.com 
````
  Notice: the permission of the executable file myhttp may change based on settings, so in case it does not have x permission please give the permission through chmod.

  running arguments which is defined by the task is "parallel" and it can be given as an input integer flag to control limit the number of parallel requests.

  # urls
  BaseURL: /md5urls
  Routes:
    -
      URL: md5urls
      Description: download the urls and calculate md5 0f the results and prints in screen.
      Method: POST
      Data: json format in body which is `urls` array of string
       
      Response :
        just standard rest response is respond in the response which means statusCode 200 for Ok cases, otherwise appropriate error status code for example 400 for empty url array.
   
### how to check        
  it is so simple!you can run and test the method with the postman collection included in the base folder named as `Md5Urls.postman_collection.json`  
  
  or simply use curl such as:
  curl -s -X POST  --url http://localhost:8000/md5urls --data '{"urls":["youtube.com","google.com","yahoo.com","golang.org"]}'