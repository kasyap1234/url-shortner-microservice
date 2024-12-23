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
![Screenshot from 2024-12-18 12-43-08](https://github.com/user-attachments/assets/ae436560-2126-4209-bd9b-e9146ff308f9)
![Screenshot from 2024-12-18 12-42-54](https://github.com/user-attachments/assets/530dc64e-f46d-43bc-8f58-53aa3a91f7bf)
