# pathserve
A server that returns a JSON encoded recursive list of files in a directory

After building run: ./pathserve 

To use provide a directory path in the url:

http://127.0.0.1:8080/path/home/testing/

Example response:
```javascript
{
Files: [
{
Name: "/go",
Size: 4096,
ModTime: "2014-12-30T09:59:21.9505971-05:00",
IsDir: true
},
{
Name: "\go\AUTHORS",
Size: 17575,
ModTime: "2014-12-11T01:45:12-05:00",
IsDir: false
},
{
Name: "\go\CONTRIBUTORS",
Size: 24564,
ModTime: "2014-12-11T01:45:12-05:00",
IsDir: false
},
{
Name: "\go\LICENSE",
Size: 1479,
ModTime: "2014-12-11T01:45:12-05:00",
IsDir: false
}
}
```
