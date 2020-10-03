--
with count_table as(
    select count(city) as num from station group by city 
) 
SELECT
    (SELECT COUNT(city) FROM station )
  - (select count(num) from count_table where num >= 1) AS Difference

-- or use this 

    select (count(city) - count(distinct city)) from station;   