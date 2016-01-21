# webimagecheck
i run a web(http://www.zhiliaoyuan.com). it was built using wordpress; some times i find some piture being not show. so i want to write a tool to check my web images. Find out it can not be show.

in the project, use goquery to parse web structure and get the image path. 

use http.get to read image data. via the image header data judge the image can be show or not.

at the same time, the check record will be record, some times to be write to stdout or file.
