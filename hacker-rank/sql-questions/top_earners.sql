-- variable and use variable

DECLARE @var int               
select @var = (select top 1 months * salary from employee order by months * salary desc)
 
select @var, count(*) from employee where months * salary = @var ;