# TestCaseRESTY
a simple test case to debug RESTY call request being sent twice


# webserver_test
In the webserver.go, change the monitorEndPoint := "192.168.48.10:8050" with the IP address and port of your running monitor process.
>cd webserver_test

>go build

>webserver_test.exe

# monitor_test
In the monitor.go, change the "time.Sleep(60 * time.Second)" and set a value like 60 seconds, and it will fail with a "net::ERR_EMPTY_RESPONSE" and there will be 2 call request from RESTY.
If you set the value to less than 5 seconds, it will succeed and only one call request.
>cd monitor_test

>go build

>monitor_test

In the browser, click on "click" link.
