# url-shortner-microservice

using sync mutex to prevent race condition 

eg : 
imagine two requests for the same url 
example.com  via 2 goroutines , 
assume count of url =0; 

first goroutine will increment the counter by 1 ; 
updated count =1; 
second goroutine will also increment the counter by 1 ; 
updated count=1; 
without mutext , we will get udpdated count=1; 
because the second goroutine will also read the same initial value of count=0;


defer statements are eecuted when the function returns ; 