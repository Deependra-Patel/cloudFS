# cloudFS
This allows following operations:

### Create
Create a text file with some contents stored in a given path.
```console
curl -d 'file=/tmp/dir/yyy&content=abc skfj fff' http://localhost:8080/create  
```

### Retrieve
Retrieve the contents of a text file under the given path.
```console
curl -d 'file=/tmp/dir/yyy' http://localhost:8080/retrieve
abc skfj fff     
```

### Replace
Replace the contents of a text file.
```console
curl -d 'file=/tmp/dir/yyy&content=abc sk d' http://localhost:8080/replace
```
### Stats
```console
curl -d 'file=/tmp/dir/' http://localhost:8080/stats

Total number of files: 1
Average number of alphanumeric characters per text file: 6.000000, StdDev: 0.000000
Average word length in folder: 2.000000, StdDev: 1.414214
Total number of bytes stored in folder: 8
```

### Delete
Delete the resource that is stored under a given path.
```console
curl -d 'file=/tmp/dir/' http://localhost:8080/stats
```