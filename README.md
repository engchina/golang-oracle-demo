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
- https://xorm.io/