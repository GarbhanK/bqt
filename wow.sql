
select *
  from `{{ params.project }}.transactions.coffee`
 where date(insertionTimestamp) >= '2023-08-16'
 group by insertionTimestamp desc

