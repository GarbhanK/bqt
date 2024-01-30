select * from `${params.project}.transactions.coffee` c
where date(insertionTimestamp) >= '${ds_nodash}'
left join `${params.web_project}.user_data.signup` t
on c.userId = t.userId
group by insertionTimestamp desc