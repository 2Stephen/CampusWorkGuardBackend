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

   verify_status VARCHAR(20) NOT NULL DEFAULT 'pending'
       COMMENT '验证状态：pending/verified/unverified',

   fail_info VARCHAR(255) DEFAULT NULL COMMENT '上次验证失败信息',

   UNIQUE KEY uk_email (email),
   UNIQUE KEY uk_social_code (social_code)
) COMMENT='企业用户注册与认证表';

CREATE TABLE IF NOT EXISTS job_infos (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'id（主键）',

    name VARCHAR(100) NOT NULL COMMENT '岗位名称',

    type VARCHAR(20) NOT NULL COMMENT '岗位类型（part-time/intern/full-time）',

    salary INT NOT NULL COMMENT '薪资标准',

    salary_unit VARCHAR(20) NOT NULL COMMENT '薪资单位（hour/day/month）',

    salary_period VARCHAR(20) NOT NULL COMMENT '薪资发放周期（day/week/month）',

    content TEXT COMMENT '工作内容',

    headcount INT COMMENT '招聘人数',

    major VARCHAR(100) COMMENT '专业要求',

    region VARCHAR(100) COMMENT '工作地点（省/市/区）',

    address VARCHAR(255) COMMENT '详细地址',

    shift VARCHAR(20) COMMENT '工作时段（day/night/shift）',

    experience VARCHAR(20) COMMENT '经验要求（none/<1/1-3/>3）',

    picture_list TEXT COMMENT '岗位相关图片（/xx/xx.jpg;/xx/xx.jpg）',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '岗位发布时间',

    status VARCHAR(20) DEFAULT 'pending' COMMENT '审核状态（pending/approved/rejected）',

    company_id VARCHAR(50) NOT NULL COMMENT '发布公司id'

    fail_info VARCHAR(255) DEFAULT NULL COMMENT '上次审核失败信息'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='岗位表';

CREATE TABLE IF NOT EXISTS admin_users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) DEFAULT NULL,
    avatar_url VARCHAR(255) DEFAULT NULL
);