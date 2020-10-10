-- 
--
select CAST(ROUND(sum(lat_n), 2) AS DECIMAL(8,2)), 
CAST(round(sum(long_w), 2) AS DECIMAL(8,2))  from station;