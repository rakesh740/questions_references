SELECT  city, max(LENGTH(city)) FROM Station group by city order by LENGTH(city) desc, city  LIMIT 1 ;
SELECT  city, min(LENGTH(city)) FROM Station group by city order by LENGTH(city) asc, city LIMIT 1 ;

