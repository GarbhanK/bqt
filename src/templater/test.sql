select *
  from `{{ params.project }}.transactions.coffee` c
  left join `{{ params.web_project }}.user_data.signup` t
    on c.userId = t.userId
 where date(insertionTimestamp) >= '{{ ds_nodash }}'
 group by insertionTimestamp desc
