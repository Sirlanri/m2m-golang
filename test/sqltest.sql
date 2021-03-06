/* 选择每个日期有人的次数 */
select date(itime),COUNT(date(itime)) from bodysensor 
    WHERE status=1;

/* 计算一天中 每小时的次数*/
SELECT date(itime),hour(itime)
    from bodysensor GROUP BY minute(itime)
     ORDER BY date(itime);

/* 计算一天中 每小时的次数*/
SELECT day(itime),hour(itime)
    from bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day);
    GROUP BY day(itime) ORDER BY hour(itime);



SELECT itime from bodysensor where date(itime)="2020-11-01" ORDER BY itime;

/* 计算一天中 每小时的次数*/
SELECT rday,rhour,count(rhour) FROM 
(SELECT itime, date(itime) as rday,hour(itime) as rhour
    from bodysensor GROUP BY itime ORDER BY itime) T
GROUP BY itime;

/* 计算一天中 每小时的次数*/
SELECT itime, count(*) from bodysensor
GROUP BY hour(itime) ORDER BY itime;

SELECT COUNT(*) as cnt
from bodysensor
WHERE itime>=DATE_SUB(now(),interval 1 hour);

SELECT itime, count(*)
FROM bodysensor
GROUP BY itime ORDER BY itime;

SELECT day(itime), hour(itime)
FROM bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day)
GROUP BY itime ORDER BY itime;

SELECT count(day(itime)), hour(itime)
		FROM bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY itime ORDER BY itime;

/*当天 有人和没人的次数*/
SELECT COUNT(*) FROM bodysensor
    WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=1 
    UNION 
SELECT COUNT(*) FROM bodysensor
    WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=0;

/*计算一周每天温度的平均值*/
SELECT round(AVG(num),2) FROM tempsensor 
WHERE itime>=DATE_SUB(now(),interval 7 day)
GROUP BY day(itime) ORDER BY day(itime);

/*能跑了 获取近七天的每小时次数*/
SELECT day(itime), hour(itime)
FROM bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day)
GROUP BY itime ORDER BY itime;

SELECT day(itime), hour(itime)
		FROM bodysensor WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY itime ORDER BY itime;

DELETE from mym2m.tempsensor;
DELETE FROM mym2m.humisensor;

DELETE FROM mym2m.lightsensor;
DELETE FROM mym2m.bodysensor;