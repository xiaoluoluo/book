DROP TABLE IF EXISTS `question`;
create table question(
	question_id int(20) NOT NULL AUTO_INCREMENT COMMENT '题目id',
    user_id int(20) NOT NULL DEFAULT 0 COMMENT '用户id',
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
    insert_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    ts timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(user_id),
    INDEX (user_wid)
) ENGINE=InnoDB  AUTO_INCREMENT=100  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';






