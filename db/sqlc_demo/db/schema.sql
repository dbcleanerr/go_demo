create table emp
(
    empno       serial primary key,
    ename       text not null,
    job         text,
    mgr         int,
    hiredate    date,
    sal         numeric(7,2),
    comm        numeric(7,2),
    deptno      int not null
);

create table dept
(
    deptno      serial primary key,
    dname       text not null,
    loc         text not null
);