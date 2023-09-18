# urlcleaner
It sorts all the urls for more uniqueness. Considering a list of urls can have same parameters with different values considered unique, using it with tools like nuclei generates a lot of unnecessary traffic and IP restrictions. 

The main goal is to reduce the effective number of urls to lower the traffic and being nice to servers ultimately. 

For example:

The list of urls below.

https://www.google.com
https://www.google.com?u=100
https://www.google.com?u=199


When passed as an stdin to sort -u , the output remains the same. 

When passed through urlcleaner as stdin, the output comes out to 

https://www.google.com
https://www.google.com?u=100
