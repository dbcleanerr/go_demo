copy (select url, url_cn, pct from qq_out_rst_url order by url_cn ) to '/tmp/com_tag0';

-- 备份
create table qq_output_20200509
as
select * from qq_output;

create table qq_out_rst_url_20200509
as
select * from qq_out_rst_url;

-- qq_out_rst_url 迁移 bundle_id
alter table qq_out_rst_url add bundle_id text;

update qq_out_rst_url
set bundle_id = qq_output.bundle_id
from qq_output
where qq_out_rst_url.cn=qq_output.cn;

select * from qq_out_rst_url where cn in (484,485,486);

select * from qq_output where cn in (484,485,486);


select max(url_cn) from qq_out_rst_url;

-- 新增 qq_out_rst_url
-- create sequence seq_qq_out_rst_url_cn start 62902;
-- create sequence seq_qq_out_rst_url_url_cn start 211923;

do $$
declare
    _line record;
    _cnt int;
    _cn  int;
begin
  for _line in select *
                    from dim_lt_migration
                    where (app, domain) not in(select app, url from qq_out_rst_url)
                      and yyyymmdd = '20200429'::date
    loop
      -- 检查app是否存在
      select count(*)
        into _cnt
        from qq_out_rst_url
       where app = _line.app;

      if _cnt > 0 then
        raise notice 'find';
        select cn
          into _cn
          from qq_out_rst_url
         where app = _line.app;

        insert into qq_out_rst_url(cn, app, url, url_cn, pct, bundle_id, yyyymmdd)
          values (_cn, _line.app, _line.domain, nextval('seq_qq_out_rst_url_url_cn'), 100, '', current_date);
      else
          raise notice 'not find';
          insert into qq_out_rst_url(cn, app, url, url_cn, pct, bundle_id, yyyymmdd)
          values (nextval('seq_qq_out_rst_url_cn'), _line.app, _line.domain, nextval('seq_qq_out_rst_url_url_cn'), 100, '', current_date);
      end if;
    end loop;
end;
$$;


-------------------------------------------------------------------------------------
-- 20200520 强哥新增标签
-- 备份
create table qq_out_rst_url_20200520
as
select * from qq_out_rst_url;

-- 导入数据
copy adhoc_qg_20200520 from '/home/yy/20200520/lt_dk-518.txt';

-- 新增
insert into qq_out_rst_url
select null, null, col01, nextval('seq_qq_out_rst_url_url_cn'), 100, null, current_date
  from adhoc_qg_20200520
 where col01 not in (select url from qq_out_rst_url)
;


insert into qq_out_rst_url
select null, null, col01, nextval('seq_qq_out_rst_url_url_cn'), 100, null, current_date
  from (select distinct col01 as col01 from adhoc_qg_20200520_2) v01
 where col01 not in (select url from qq_out_rst_url)
;

insert into qq_out_rst_url
select null, null, col01, nextval('seq_qq_out_rst_url_url_cn'), 100, null, current_date
  from (select distinct col01 as col01 from adhoc_qg_20200520_3) v01
 where col01 not in (select url from qq_out_rst_url)
;

insert into qq_out_rst_url
select null, null, col01, nextval('seq_qq_out_rst_url_url_cn'), 100, null, current_date
  from (select distinct col01 as col01 from adhoc_qg_20200520_4) v01
 where col01 not in (select url from qq_out_rst_url)
;

copy (select url, url_cn, pct from qq_out_rst_url order by url_cn ) to '/tmp/com_tag0';

select max(url_cn) from qq_out_rst_url;

--
select * from adhoc_qg_20200520
where col01 in (
select url from qq_out_rst_url )

select * from qq_out_rst_url where url = 'infocollect.weidai.com.cn'

delete from qq_out_rst_url where yyyymmdd = current_date and url = '';

select * from qq_out_rst_url where yyyymmdd = current_date and url = '';

select * from qq_out_rst_url where yyyymmdd = current_date;