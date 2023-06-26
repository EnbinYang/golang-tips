# gin

GET Request
```bash
go test -v ./ -run=GetName -count=1  # -count=1 can clear cache
```

POST Request
```bash
go test -v ./ -run=GetAge -count=1
```

POST Request (data type is JSON format)
```bash
go test -v ./ -run=GetHeight -count=1
```