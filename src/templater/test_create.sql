create or replace `my-project.transactions.coffee`
as
select * from `{{ params.project }}.transactions.coffee` c
where date(insertionTimestamp) >= '{{ ds_nodash }}'
left join `{{ params.web_project }}.unified_segment.tracks` t
on c.userId = t.userId
group by insertionTimestamp desc