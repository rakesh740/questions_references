
-- city name does not start with vowel 

SELECT distinct City FROM STATION WHERE City NOT LIKE '[A,E,I,O,U]%';

-- city name start and end with vowel
SELECT distinct City FROM STATION WHERE City LIKE '[A,E,I,O,U]%[A,E,I,O,U]';

-- do not start or end with vowel 
select distinct city from station where (left(city,1) not in ('a','e','i','o','u') and  right(city,1) not in ('a','e','i','o','u'));

-- 