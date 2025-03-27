USE mysql;
create database user;
create database chat;
create database activity;
create database memoir;
CREATE TABLE user.user (
                           user_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '用户唯一标识',
                           username VARCHAR(50) NOT NULL COMMENT '用户名',
                           email VARCHAR(100) UNIQUE NOT NULL COMMENT '邮箱',
                           password VARCHAR(255) NOT NULL COMMENT '密码哈希值',
                           avatar_url VARCHAR(255) COMMENT '头像URL',
                           bio  VARCHAR(255) COMMENT '用户简介',
                           register_date DATETIME NOT NULL COMMENT '注册时间',
                           theme VARCHAR(20) DEFAULT 'light' COMMENT '主题偏好'
) COMMENT '用户基本信息表';


-- 对话容器表 (存储会话元数据)
CREATE TABLE chat.conversation (
                                   conversation_id  INT AUTO_INCREMENT PRIMARY KEY COMMENT '对话容器唯一ID',
                                   user_id INT NOT NULL COMMENT '关联用户ID',
                                   session_title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '会话标题/话题',
                                   mode VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'happy',
                                   updated_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '会话创建时间',
                                   created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '会话创建时间'
) COMMENT '对话容器表';
CREATE INDEX idx_conversations_user_id_create_time ON chat.conversation(user_id, created_at DESC);

-- 消息内容表 (存储具体聊天记录)
CREATE TABLE chat.message (
                              message_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '消息ID',
                              user_id INT NOT NULL COMMENT '关联用户ID',
                              conversation_id INT NOT NULL COMMENT '所属对话容器ID（关联字段）',
                              content TEXT NOT NULL COMMENT '消息正文内容',
                              sender_type VARCHAR(20) NOT NULL COMMENT '发送方类型，user，ai',
                              send_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间（主排序索引）'
) COMMENT '消息内容表';
CREATE INDEX idx_messages_conversation_id_send_time ON chat.message(conversation_id, send_time DESC);

CREATE TABLE memoir.memoir (
                               memoir_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '回忆录唯一标识',
                               user_id INT NOT NULL COMMENT '关联用户ID',
                               title VARCHAR(255) NOT NULL COMMENT '回忆录标题',
                               type VARCHAR(20) NOT NULL DEFAULT '' COMMENT '回忆录类型，有日记，周记，月记',
                               style VARCHAR(20) NOT NULL DEFAULT '' COMMENT '回忆录风格，有 happy, sad',
                               start_date DATETIME NOT NULL COMMENT '开始日期',
                               end_date DATETIME NOT NULL COMMENT '结束日期',
                               content TEXT NOT NULL COMMENT '回忆录内容文本',
                               created_at TIMESTAMP comment '创建日期'
) COMMENT '回忆录存储表';
CREATE INDEX idx_memoirs_user_id_start_date ON memoir.memoir(user_id, created_at DESC);
