# golang-oracle-demo

### Update database date format
```shell
sqlplus / as sysdba

SQL> alter system set nls_date_format='YYYY-MM-DD HH24:MI:SS' scope=spfile;
SQL> alter system set nls_time_format='HH24:MI:SS' scope=spfile;
SQL> alter system set nls_timestamp_format='YYYY-MM-DD HH24:MI:SS.FF' scope=spfile;

SQL> shutdown immediate;
SQL> startup;
```

### Reference:
- https://getbootstrap.com/docs/5.2/getting-started/introduction/
- https://pkg.go.dev/crypto/md5
- https://xorm.io/
- https://pkg.go.dev/github.com/godror/godror
- https://pkg.go.dev/github.com/gin-gonic/gin
- https://pkg.go.dev/github.com/spf13/viper
- https://pkg.go.dev/github.com/sirupsen/logrus
- https://pkg.go.dev/github.com/swaggo/gin-swagger
- https://pkg.go.dev/github.com/asaskevich/govalidator
- https://pkg.go.dev/github.com/gomodule/redigo