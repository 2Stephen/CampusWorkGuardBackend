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
    UNIQUE KEY uk_email (email)
);

