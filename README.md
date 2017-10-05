# Func-Redshirts
Func-Redshirts explains the difference between `func (r *redshirt)`
and `func (r redshirt)`.

Simply run the following:

```bash
$ go run main.go
r1.data=	r2.data=
r1.data=	r2.data=data
```

The above program illustrates the difference between declaring a function
on a type and declaring a function on the address of a type. Please
see [`main.go`](https://github.com/akutz/func-redshirts/blob/master/main.go)
for additional information.
