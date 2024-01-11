CREATE TABLE `union_settled_apply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id', --  
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '申请人uid', -- 
  `country` varchar(10) NOT NULL DEFAULT '' COMMENT '国家', -- 
  `region` varchar(10) NOT NULL DEFAULT '' COMMENT '地区', -- 
  `union_name` varchar(100) NOT NULL DEFAULT '' COMMENT '公会名称', --
  `union_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '公会类型', -- 
  `contact_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '联系方式', -- 
  `contact_info` varchar(100) NOT NULL DEFAULT '' COMMENT '联系信息', -- 
  `contact_type_other` tinyint(4) NOT NULL DEFAULT 0 COMMENT '备用联系方式', -- 
  `contact_info_other` varchar(100) NOT NULL DEFAULT '' COMMENT '备用联系信息', -- 
  `identity_photo` varchar(2000) NOT NULL DEFAULT '' COMMENT '身份照片', -- 
  `union_prove` varchar(2000) NOT NULL DEFAULT '' COMMENT '公会证明', -- 
  `business_license` varchar(2000) NOT NULL DEFAULT '' COMMENT '公司营业执照', -- 
  `union_people_num` bigint(20) NOT NULL DEFAULT 0 COMMENT '预计公会人数', -- 
  `recommend_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '推荐人ID或者公会ID', -- 
  `source_app` varchar(255) NOT NULL DEFAULT '' COMMENT '来源其他app', -- 
  `union_experience` varchar(255) NOT NULL DEFAULT '' COMMENT '公会经验', -- 
  `audit_status` tinyint(2) DEFAULT '0' COMMENT '审批状态', -- 
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `audit_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '审核时间',
  `operator` varchar(100) NOT NULL DEFAULT '' COMMENT '审核人',
  `remark` varchar(150) NOT NULL DEFAULT '' COMMENT '备注-50字以内',
  `reject_reason` varchar(150) NOT NULL DEFAULT '' COMMENT '拒绝原因-50字以内',
  `extend` json DEFAULT NULL COMMENT 'extend',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid` (`uid`) USING BTREE COMMENT '公会长唯一索引',
  KEY `idx_create_time` (`create_time`) USING BTREE COMMENT '创建时间索引',
  KEY `idx_region` (`region`) USING BTREE COMMENT '地区索引',
  KEY `idx_audit_status` (`audit_status`) USING BTREE COMMENT '审批状态'
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '公会入驻申请表';