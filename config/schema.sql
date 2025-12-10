CREATE TABLE IF NOT EXISTS chsi_student_infos (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    gender VARCHAR(10),
    birthday VARCHAR(20),
    nation VARCHAR(50),
    school VARCHAR(100),
    level VARCHAR(50),
    major VARCHAR(100),
    duration VARCHAR(50),
    degree_type VARCHAR(50),
    study_mode VARCHAR(50),
    college VARCHAR(100),
    department VARCHAR(100),
    entrance_date VARCHAR(20),
    status VARCHAR(50),
    expected_grad VARCHAR(20),
    vcode VARCHAR(50),
    student_id VARCHAR(50),
    email VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS student_users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    school_id INT NOT NULL,
    student_id VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255) DEFAULT NULL,
    UNIQUE KEY uk_email (email)
);

CREATE TABLE IF NOT EXISTS company_users (
   id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',

   name VARCHAR(50) COMMENT '注册人姓名',
   email VARCHAR(100) COMMENT '注册人邮箱（公司邮箱）',
   company VARCHAR(100) COMMENT '公司名称（需与营业执照一致）',

   license_url VARCHAR(255) COMMENT '营业执照相对URL',
   social_code VARCHAR(32) COMMENT '统一社会信用代码',

   password VARCHAR(255) DEFAULT NULL COMMENT '密码（初始为空，可后续修改）',
   avatar_url VARCHAR(255) DEFAULT NULL COMMENT '头像URL（默认头像）',

   verify_status VARCHAR(20) NOT NULL DEFAULT '验证中'
       COMMENT '验证状态：验证中/验证成功/验证失败',

   fail_info VARCHAR(255) DEFAULT NULL COMMENT '上次验证失败信息',

   UNIQUE KEY uk_email (email),
   UNIQUE KEY uk_social_code (social_code)
) COMMENT='企业用户注册与认证表';


