select *
  from `{{ params.project }}.transactions.coffee` c
  left join `{{ params.web_project }}.unified_segment.tracks` t
    on c.userId = t.userId
 where date(insertionTimestamp) >= '{{ ds_nodash }}'
 group by insertionTimestamp desc
