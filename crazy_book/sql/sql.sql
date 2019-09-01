DROP TABLE IF EXISTS `question`;
create table question(
	question_id int(20) NOT NULL AUTO_INCREMENT COMMENT '题目id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '用户id',
    user_grade int(20) NOT NULL DEFAULT 0 COMMENT '用户所在年级',
	question_title varchar(255) NOT NULL DEFAULT '' COMMENT '题目标题',
	answer_pic varchar(255) NOT NULL DEFAULT '' COMMENT '题目图片',
    subject_code int(10) NOT NULL DEFAULT 0 COMMENT '科目代码 1数学2语文3英语',
    true_title varchar(255) NOT NULL DEFAULT '' COMMENT '正解title',
    true_pic varchar(255) NOT NULL DEFAULT '' COMMENT '正解图片',
    false_title varchar(255) NOT NULL DEFAULT '' COMMENT '错解title',
    false_pic varchar(255) NOT NULL DEFAULT '' COMMENT '错解图片',
	insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
	ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY(question_id),
    INDEX (user_id)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='题目表';


DROP TABLE IF EXISTS `user`;
create table user(
    user_id int(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    user_wid varchar(100) NOT NULL DEFAULT '' COMMENT '用户微信id',
    user_name varchar(100) NOT NULL DEFAULT '' COMMENT '用户名字',
    user_head_pic varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    user_grade int(20) NOT NULL DEFAULT 0 COMMENT '用户所在年级',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(user_id),
    INDEX (user_wid)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';




DROP TABLE IF EXISTS `comment`;
create table comment(
    comment_id int(20) NOT NULL AUTO_INCREMENT COMMENT '评论递增id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '评论用户id',
    question_id int(20) NOT NULL DEFAULT 0 COMMENT '题目id',
    comment_intro varchar(1024) NOT NULL DEFAULT '' COMMENT '评论内容',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(comment_id),
    INDEX (user_id),
    INDEX (question_id)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='评论表';



DROP TABLE IF EXISTS `collection`;
create table collection(
    collection_id int(20) NOT NULL AUTO_INCREMENT COMMENT '收藏递增id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '收藏用户id',
    question_id int(20) NOT NULL DEFAULT 0 COMMENT '题目id',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(collection_id),
    INDEX (user_id),
    UNIQUE INDEX (user_id,question_id)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='收藏表';


DROP TABLE IF EXISTS `liked`;
create table liked(
    liked_id int(20) NOT NULL AUTO_INCREMENT COMMENT '点赞递增id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '点赞用户id',
    question_id int(20) NOT NULL DEFAULT 0 COMMENT '题目id',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(liked_id),
    INDEX (user_id),
    UNIQUE INDEX (user_id,question_id)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='点赞表';


DROP TABLE IF EXISTS `label`;
create table label(
    label_id int(20) NOT NULL AUTO_INCREMENT COMMENT '标签递增id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '标签用户id',
    subject_code int(10) NOT NULL DEFAULT 0 COMMENT '课程代码',
    label varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(label_id),
    INDEX (user_id)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='标签表';


# 知识点初始化表
DROP TABLE IF EXISTS `initLabel`;
create table initLabel(
    grade  int(10) NOT NULL DEFAULT 0 COMMENT '年级',
    term   int(10) NOT NULL DEFAULT 0 COMMENT '学期：1上学期 2下学期',
    subject_code int(10) NOT NULL DEFAULT 0 COMMENT '课程代码',
    label varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB  AUTO_INCREMENT=1  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='初始标签表';

INSERT INTO `initLabel`  (grade, term,subject_code,label)  VALUES
(1,2,2,'电学'),
(1,2,2,'力学'),
(1,2,2,'电磁学'),
(1,2,2,'空气动力学'),
(1,2,2,'基因'),
(1,2,2,'多次函数');