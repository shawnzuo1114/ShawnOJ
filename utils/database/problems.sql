-- ----------------------------
-- 注意：执行顺序为先建 problems 表（被 test_cases 外键关联）
-- ----------------------------

-- 1. 题目主表（problems）：核心题面信息
DROP TABLE IF EXISTS `problems`;
CREATE TABLE `problems` (
                            `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '题目唯一ID（自增主键）',
                            `title`         VARCHAR(100)    NOT NULL COMMENT '题目标题（如：两数之和）',
                            `slug`          VARCHAR(128)    NOT NULL COMMENT 'URL友好标识（如two-sum，全局唯一，软删除后仍不可重复）',
                            `description`   LONGTEXT        NOT NULL COMMENT '题目描述（Markdown格式，含题干、样例说明）',
                            `input_desc`    TEXT NULL COMMENT '输入格式说明（抽离便于前端单独展示）',
                            `output_desc`   TEXT NULL COMMENT '输出格式说明（抽离便于前端单独展示）',
                            `time_limit`    SMALLINT UNSIGNED NOT NULL DEFAULT 1000 COMMENT '默认时间限制（ms）：100~10000ms',
                            `memory_limit`  SMALLINT UNSIGNED NOT NULL DEFAULT 128 COMMENT '默认内存限制（MB）：16~1024MB',
                            `difficulty`    TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '难度等级：1-简单 2-中等 3-困难',
                            `creator_id`    BIGINT UNSIGNED NOT NULL COMMENT '创建者用户ID（关联users表id）',
                            `is_public`     TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '公开状态：1-公开 0-私有',
                            `submit_count`  INT UNSIGNED    NOT NULL DEFAULT 0 COMMENT '总提交数（避免实时统计）',
                            `accept_count`  INT UNSIGNED    NOT NULL DEFAULT 0 COMMENT '通过提交数（避免实时统计）',
                            `pass_rate`     DECIMAL(5,2)    NOT NULL DEFAULT 0.00 COMMENT '通过率：accept_count/submit_count（百分比）',
                            `created_at`    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at`    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at`    DATETIME NULL COMMENT '软删除时间（NULL=未删除，非NULL=删除时间）',
                            PRIMARY KEY (`id`),
    -- 唯一索引：兼容软删除（deleted_at为NULL时slug唯一）
                            UNIQUE KEY `uk_problem_slug_deleted` (`slug`, `deleted_at`),
    -- 普通索引：高频查询场景
                            KEY `idx_problem_creator` (`creator_id`),                  -- 查用户发布的题目
                            KEY `idx_problem_public_diff` (`is_public`, `difficulty`), -- 筛选公开题目+按难度分类
                            KEY `idx_problem_deleted` (`deleted_at`)                   -- 软删除数据快速过滤
) ENGINE=InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT='题目主表（1:N 关联测试用例表）';

-- 2. 测试用例表（test_cases）：单条用例存储，精细化管理
DROP TABLE IF EXISTS `test_cases`;
CREATE TABLE `test_cases` (
                              `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '测试用例唯一ID',
                              `problem_id`      BIGINT UNSIGNED NOT NULL COMMENT '所属题目ID（关联problems表id）',
                              `input`           TEXT NOT NULL COMMENT '测试用例输入数据（纯文本）',
                              `expected_output` TEXT NOT NULL COMMENT '测试用例期望输出（纯文本）',
                              `is_sample`       TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否样例：1-样例（前端展示） 0-普通（仅评测用）',
                              `score`           TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '单条用例分值（总分100，0=不计分）',
                              `time_limit`      SMALLINT UNSIGNED NULL COMMENT '自定义时间限制（ms），NULL=继承题目默认值',
                              `memory_limit`    SMALLINT UNSIGNED NULL COMMENT '自定义内存限制（MB），NULL=继承题目默认值',
                              `created_at`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at`      DATETIME NULL COMMENT '软删除时间（NULL=未删除，非NULL=删除时间）',
                              PRIMARY KEY (`id`),
    -- 普通索引：高频查询场景
                              KEY `idx_testcase_problem` (`problem_id`),                      -- 查某题的所有用例
                              KEY `idx_testcase_problem_sample` (`problem_id`, `is_sample`),  -- 查某题的样例/普通用例
                              KEY `idx_testcase_deleted` (`deleted_at`),                      -- 软删除数据快速过滤
    -- 外键约束：删除题目时级联删除用例，避免数据孤岛
                              CONSTRAINT `fk_testcase_problem`
                                  FOREIGN KEY (`problem_id`) REFERENCES `problems`(`id`)
                                      ON DELETE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT='题目测试用例表（N:1 关联problems表，单条用例存储）';
