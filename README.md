# HTTP/2 - How could it break you

Unlike HTTP/1.1, HTTP/2 sends all the requests over a single TCP connection. Sort of like bundling up all the requests from client and delivering it to the server.

What could possibly go wrong !!

Well, bundling the request causes a spike in the server !?  
This project tries to compare request fetching over HTTP/1.1 and HTTP/2 protocols

## How to use
Run the following command from the root of the project  

`docker compose up`

Open the dev tools of the browser and throttle the network speed. Then in the browser type in [http://localhost/http1](http://localhost/http1)  
You'll see all the 49 images loading up albeit slowly

Now, with same network throttling as before type, this in the browser [https://localhost/http2](https://localhost/http2). Ignore the SSL warning.
You'll see the images loading very fast but only 26 images have been received by the browser. This is because of the spike in the request received by the server.

Now in the nginx.conf file, update the burst attribute to 50. See the below image for reference.

<img width="516" height="353" alt="image" src="https://github.com/user-attachments/assets/f5450491-e0e8-4059-be98-d24060e0004e" />












