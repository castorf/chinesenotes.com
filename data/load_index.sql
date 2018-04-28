USE cse_dict;

LOAD DATA LOCAL INFILE 'corpus/collections.csv' INTO TABLE collection CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/erya.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/gongyang.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/guliang.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/laoshe.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/liji.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/lunyu.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/modern_articles.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/shiji.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/shangshu.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/shijing.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/shuowen.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/sishuzhangju.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/tang_poetry.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/tea_classic.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/yeshengtao.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/yili.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/zhuangzi.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/zhouli.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;

LOAD DATA LOCAL INFILE 'corpus/zuozhuan.csv' INTO TABLE document CHARACTER SET utf8 LINES TERMINATED BY '\n' IGNORE 2 LINES;
SHOW WARNINGS;
