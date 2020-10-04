-- variable and use variable

DECLARE @var int               
select @var = (select top 1 months * salary from employee order by months * salary desc)
 
select @var, count(*) from employee where months * salary = @var ;


-- alternate


select (salary * months)as earnings ,count(*) from employee group by 1 order by earnings desc limit 1 ;


--

SELECT * FROM (SELECT MAX(MONTHS*SALARY), COUNT(EMPLOYEE_ID) FROM EMPLOYEE GROUP BY MONTHS*SALARY ORDER BY MONTHS*SALARY DESC) WHERE ROWNUM = 1;

