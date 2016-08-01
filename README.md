# caddy-ace
add [Ace](https://github.com/yosssi/ace) template engine directive &amp; plugin for caddy server



Middleware for [Caddy](https://caddyserver.com).

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/pschlump/Go-FTL/master/LICENSE)


### Usage

```
gzip
log ../access.log
ace  {
    path /example
}
```
* **path** whose value is the relative path where you store your ace sourcefile 


Then just type 
```
caddy 
```
and visit [http://localhost:2015/example/](http://localhost:2015/example/)